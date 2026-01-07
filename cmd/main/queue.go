package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
	_ "time/tzdata"

	"github.com/captainkie/websync-api/config"
	model "github.com/captainkie/websync-api/internal/app/models"
	"github.com/captainkie/websync-api/internal/app/repository"
	"github.com/captainkie/websync-api/internal/app/service"
	"github.com/captainkie/websync-api/pkg/helpers"
	magentoServiceAttribute "github.com/captainkie/websync-api/pkg/magento2/attributes"
	magentoServiceInventory "github.com/captainkie/websync-api/pkg/magento2/inventory"
	magentoServiceMedia "github.com/captainkie/websync-api/pkg/magento2/media"
	magentoServiceProduct "github.com/captainkie/websync-api/pkg/magento2/products"
	"github.com/captainkie/websync-api/types/payload"
	"github.com/captainkie/websync-api/types/request"
	"github.com/go-playground/validator/v10"
	"github.com/gosimple/slug"
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

type StatusResponse struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func producer() {
	// Clear && Load env
	os.Clearenv()
	err := godotenv.Load()
	helpers.ErrorPanic(err)

	os.Setenv("TZ", "Asia/Bangkok")

	// connect RabbitMQ
	rabbitMQURL := os.Getenv("RABBITMQ_URL")
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Printf("Failed to connect to RabbitMQ: %v", err)
	}

	if conn != nil {
		defer conn.Close()

		fmt.Println("🚀 Successfully connected to the RabbitMQ")

		// Create a RabbitMQ channel
		ch, err := conn.Channel()
		if err != nil {
			log.Printf("Failed to open a channel: %v", err)
		}
		defer ch.Close()

		// Declare queues for different task types
		channels := []string{"connection_queue", "product_queue", "postflag_queue", "image_queue", "dailysale_queue"}

		for _, channel := range channels {
			_, err := ch.QueueDeclare(
				channel, // Name of the queue
				true,    // Durable (queue survives server restarts)
				false,   // Auto-delete when unused
				false,   // Exclusive (queue only accessed by this connection)
				false,   // No-wait
				nil,     // Arguments
			)
			if err != nil {
				log.Printf("Failed to declare a queue: %v", err)
			}
		}

		var wg sync.WaitGroup

		// Start workers for each queue concurrently
		for _, channel := range channels {
			wg.Add(1)
			go consume(&wg, ch, channel)
		}

		// Wait for all worker goroutines to finish
		wg.Wait()
	}
}

func consume(wg *sync.WaitGroup, ch *amqp.Channel, queueName string) {
	defer wg.Done()

	msgs, err := ch.Consume(
		queueName, // Queue name
		"",        // Consumer tag
		false,     // Auto-acknowledge (set to false for manual acknowledgment)
		false,     // Exclusive
		false,     // No-local
		false,     // No-wait
		nil,       // Args
	)
	if err != nil {
		log.Printf("Failed to register a consumer for %s: %v", queueName, err)
	}

	db := config.ConnectDatabase()

	validate := validator.New()

	queueRepository := repository.NewQueueRepositoryImpl(db)
	loggerRepository := repository.NewLoggerRepositoryImpl(db)
	productRepository := repository.NewProductRepositoryImpl(db)
	imageRepository := repository.NewImageRepositoryImpl(db)

	queueService := service.NewQueueServiceImpl(queueRepository, validate)
	loggerService := service.NewLoggerServiceImpl(loggerRepository, validate)
	productService := service.NewProductServiceImpl(productRepository, validate)
	imageService := service.NewImageServiceImpl(imageRepository, validate)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			// process the task of connect queue
			if queueName == "connection_queue" {
				var body model.ConnectionQueues
				errModel := json.Unmarshal(d.Body, &body)
				if errModel != nil {
					log.Println(errModel)
				}

				log.Printf("Received task from %s => %s : %s\n", queueName, "[CONNECTION]", body.TransactionID)

				// processing task, update model status pending to processing
				queueService.UpdateConnectionQueue(body.ID, "processing")
				// fetch data from erp
				code, status, msg, data := get_data_from_erp(queueService)
				log.Printf("SyncStatus :: CONNECTION => [TranID=%s],[CODE=%d],[MSG=%s]\n", body.TransactionID, code, msg)
				// create logs
				add_log(code, body.TransactionID, status, "nil", data, msg, "CONNECTION", "", "", loggerService)
				// complete task, update model status processing to completed
				queueService.UpdateConnectionQueue(body.ID, "completed")
			}

			// process the task of product queue
			if queueName == "product_queue" {
				// add products
				if strings.Contains(string(d.Body), `"Type":"add"`) {
					var body model.AddQueues
					errModel := json.Unmarshal(d.Body, &body)
					if errModel != nil {
						log.Println(errModel)
					}

					var product request.AddUpdateProductRequest
					errJson := json.Unmarshal([]byte(body.JsonData), &product)
					if errJson != nil {
						log.Println(errJson)
					}

					log.Printf("Received task from %s => %s : %s\n", queueName, "[ADD]", product.TRANS_ID)

					// processing task, update model status pending to processing
					queueService.UpdateProductsQueue(body.ID, "ADD", "processing", []byte(body.JsonData))
					// add product to magento
					code, status, msg, jsonData, data := add_product_to_m2(product, productService, imageService)
					log.Printf("SyncStatus :: ADD => [TranID=%s],[CODE=%d],[MSG=%s]\n", product.TRANS_ID, code, msg)
					// create logs
					add_log(code, product.TRANS_ID, status, jsonData, data, msg, "ADD", "", "", loggerService)
					// send to postflag
					send_post_flag_queue(code, product.TRANS_ID, msg, queueService)
					// complete task, update model status processing to completed
					queueService.UpdateProductsQueue(body.ID, "ADD", "completed", []byte(body.JsonData))
				}

				// update products
				if strings.Contains(string(d.Body), `"Type":"update"`) {
					var body model.UpdateQueues
					errModel := json.Unmarshal(d.Body, &body)
					if errModel != nil {
						log.Println(errModel)
					}

					var product request.AddUpdateProductRequest
					errJson := json.Unmarshal([]byte(body.JsonData), &product)
					if errJson != nil {
						log.Println(errJson)
					}

					log.Printf("Received task from %s => %s : %s\n", queueName, "[UPDATE]", product.TRANS_ID)

					// processing task, update model status pending to processing
					queueService.UpdateProductsQueue(body.ID, "UPDATE", "processing", []byte(body.JsonData))
					// update product to magento
					code, status, msg, jsonData, data := update_product_to_m2(product, productService, imageService)
					log.Printf("SyncStatus :: UPDATE => [TranID=%s],[CODE=%d],[MSG=%s]\n", product.TRANS_ID, code, msg)
					// create logs
					add_log(code, product.TRANS_ID, status, jsonData, data, msg, "UPDATE", "", "", loggerService)
					// send to postflag
					send_post_flag_queue(code, product.TRANS_ID, msg, queueService)
					// complete task, update model status processing to completed
					queueService.UpdateProductsQueue(body.ID, "UPDATE", "completed", []byte(body.JsonData))
				}

				// stock products
				if strings.Contains(string(d.Body), `"Type":"stock"`) {
					var body model.StockQueues
					errModel := json.Unmarshal(d.Body, &body)
					if errModel != nil {
						log.Println(errModel)
					}

					var stocks request.UpdateStockRequest
					errJson := json.Unmarshal([]byte(body.JsonData), &stocks)
					if errJson != nil {
						log.Println(errJson)
					}

					transctionID := fmt.Sprintf("%d", stocks.TRANS_ID)

					log.Printf("Received task from %s => %s : %s\n", queueName, "[STOCK]", transctionID)

					// processing task, update model status pending to processing
					queueService.UpdateStockQueue(body.ID, "processing", []byte(body.JsonData))
					// update stock to magento
					code, status, msg, jsonStr, data := update_stock_to_m2(stocks)
					log.Printf("SyncStatus :: STOCK => [TranID=%s],[CODE=%d],[MSG=%s]\n", transctionID, code, msg)
					// create logs
					add_log(code, transctionID, status, jsonStr, data, msg, "STOCK", "", "", loggerService)
					// send to postflag
					send_post_flag_queue(code, transctionID, msg, queueService)
					// complete task, update model status processing to completed
					queueService.UpdateStockQueue(body.ID, "completed", []byte(body.JsonData))
				}

				// store products
				if strings.Contains(string(d.Body), `"Type":"store"`) {
					var body model.StoreQueues
					errModel := json.Unmarshal(d.Body, &body)
					if errModel != nil {
						log.Println(errModel)
					}

					var stores request.UpdateStoreRequest
					errJson := json.Unmarshal([]byte(body.JsonData), &stores)
					if errJson != nil {
						log.Println(errJson)
					}

					transctionID := fmt.Sprintf("%d", stores.TRANS_ID)

					log.Printf("Received task from %s => %s : %s\n", queueName, "[STORE]", transctionID)
					// processing task, update model status pending to processing
					queueService.UpdateStoreQueue(body.ID, "processing", []byte(body.JsonData))
					// update stock to magento
					code, status, msg, jsonStr, data := update_store_to_m2(stores)
					log.Printf("SyncStatus :: STORE => [TranID=%s],[CODE=%d],[MSG=%s]\n", transctionID, code, msg)
					// create logs
					add_log(code, transctionID, status, jsonStr, data, msg, "STORE", "", "", loggerService)
					// send to postflag
					send_post_flag_queue(code, transctionID, msg, queueService)
					// complete task, update model status processing to completed
					queueService.UpdateStoreQueue(body.ID, "completed", []byte(body.JsonData))
				}
			}

			// process the task of postflag queue
			if queueName == "postflag_queue" {
				var body model.PostflagQueues
				errModel := json.Unmarshal(d.Body, &body)
				if errModel != nil {
					log.Println(errModel)
				}

				var postflag request.CreatePostflagRequest
				errJson := json.Unmarshal([]byte(body.JsonData), &postflag)
				if errJson != nil {
					log.Println(errJson)
				}

				transctionID := fmt.Sprintf("%d", postflag.TransactionId)

				log.Printf("Received task from %s => %s : %s\n", queueName, "[POSTFLAG]", transctionID)
				// processing task, update model status pending to processing
				queueService.UpdatePostflagQueue(body.ID, "processing")
				// call postflag api
				code, status, msg, jsonStr, data := send_post_flag_to_erp(postflag)
				log.Printf("SyncStatus :: POSTFLAG => [TranID=%s],[CODE=%d],[MSG=%s]\n", transctionID, code, msg)
				// create logs
				add_log(code, transctionID, status, jsonStr, data, msg, "POSTFLAG", "", "", loggerService)
				// complete task, update model status processing to completed
				queueService.UpdatePostflagQueue(body.ID, "completed")
			}

			// process the task of image queue
			if queueName == "image_queue" {
				var body model.ImageQueues
				errModel := json.Unmarshal(d.Body, &body)
				if errModel != nil {
					log.Println(errModel)
				}

				image := request.CreateImageQueueRequest{
					TransactionID: body.TransactionID,
					Image:         body.Image,
					DirectoryPath: body.DirectoryPath,
					SyncDate:      body.SyncDate,
				}

				log.Printf("Received task from %s => %s : %s\n", queueName, "[IMAGE]", body.TransactionID)
				// processing task, update model status pending to processing
				queueService.UpdateImageQueue(body.ID, "processing")
				// process the task of image queue
				code, status, msg, jsonData, data := update_image_process(image, productService, imageService)
				log.Printf("SyncStatus :: IMAGE => [TranID=%s],[CODE=%d],[MSG=%s]\n", body.TransactionID, code, msg)
				// create logs
				add_log(code, body.TransactionID, status, jsonData, data, msg, "IMAGE", "", body.DirectoryPath+"/"+body.Image, loggerService)
				// complete task, update model status processing to completed
				queueService.UpdateImageQueue(body.ID, "completed")
				// delete task
				queueService.DeleteImageQueue(body.ID)
				// sleep for 1 second
				time.Sleep(1 * time.Second)
			}

			// process the task of dailysale queue
			if queueName == "dailysale_queue" {
				var body model.DailysaleQueues
				errModel := json.Unmarshal(d.Body, &body)
				if errModel != nil {
					log.Println(errModel)
				}

				var dailysale request.CreateDailySalesRequest
				errJson := json.Unmarshal([]byte(body.JsonData), &dailysale)
				if errJson != nil {
					log.Println(errJson)
				}

				log.Printf("Received task from %s => %s : %s\n", queueName, "[DAILYSALE]", body.TransactionID)
				// processing task, update model status pending to processing
				queueService.UpdateDailySalesQueue(body.ID, "processing")
				// process the task of image queue
				code, status, msg, jsonData, data := send_dailysale_to_erp(dailysale)
				log.Printf("SyncStatus :: DAILYSALE => [TranID=%s],[CODE=%d],[MSG=%s]\n", body.TransactionID, code, msg)
				// create logs
				add_log(code, body.TransactionID, status, jsonData, data, msg, "DAILYSALE", dailysale.DocNo, "", loggerService)
				// complete task, update model status processing to completed
				queueService.UpdateDailySalesQueue(body.ID, "completed")
			}

			// Acknowledge the message once it's processed
			d.Ack(false)
			log.Printf("Completed task from %s\n", queueName)
		}
	}()
	<-forever
}

func add_log(code int, id, status, json, data, msg, condition, order_id, image string, service service.LoggerService) {
	log := request.CreateLogRequest{
		TransactionID: id,
		Status:        status,
		StatusCode:    code,
		Message:       msg,
		SyncJson:      json,
		SyncData:      data,
		SyncDate:      time.Now(),
	}

	if condition == "CONNECTION" {
		service.CreateConnectionLog(log)
	} else if condition == "ADD" {
		service.CreateAddLog(log)
	} else if condition == "UPDATE" {
		service.CreateUpdateLog(log)
	} else if condition == "STOCK" {
		service.CreateStockLog(log)
	} else if condition == "STORE" {
		service.CreateStoreLog(log)
	} else if condition == "POSTFLAG" {
		service.CreatePostflagLog(log)
	} else if condition == "IMAGE" {
		imageLog := request.CreateLogImageRequest{
			TransactionID: id,
			Status:        status,
			StatusCode:    code,
			Message:       msg,
			SyncImage:     image,
			// SyncJson:      json,
			SyncJson: "nil",
			SyncData: data,
			SyncDate: time.Now(),
		}

		service.CreateImageLog(imageLog)
	} else if condition == "DAILYSALE" {
		dailyLog := request.CreateLogDailyRequest{
			TransactionID: id,
			OrderID:       order_id,
			Status:        status,
			StatusCode:    code,
			Message:       msg,
			SyncJson:      json,
			SyncData:      data,
			SyncDate:      time.Now(),
		}
		service.CreateDailysaleLog(dailyLog)
	}
}

func add_product_to_m2(product request.AddUpdateProductRequest, service service.ProductService, imgService service.ImageService) (int, string, string, string, string) {
	// add product to magento
	var status, msg, jsonData, data string
	var code int

	if product.PRODUCT_TYPE == "configurable" {
		insertSku, mapping, statusCode, err, jsonDataToM2 := magentoServiceProduct.CreateConfigurableProduct(product, "")
		if err != nil {
			// err to string
			status = "ERROR"
			msg = get_msg_from_json(string(err.Error()))
		} else {
			status = "SUCCESS"
			msg = "Add configurable product success"
		}

		code = statusCode
		data = insertSku
		jsonData = jsonDataToM2

		// add to configurable product mapping table
		if mapping != "" {
			var insertSkuObj payload.SimpleProductPayload
			errParse := json.Unmarshal([]byte(insertSku), &insertSkuObj)
			if errParse == nil {
				configSku := helpers.PadString(insertSkuObj.Sku, 8, '#')
				configSku = strings.ReplaceAll(configSku, "#", "")
				addProductMapping := request.ConfigurableProductRequest{
					Sku:           configSku,
					FirstChildSku: mapping,
				}

				service.CreateConfigurableProduct(addProductMapping)
			}
		}
	} else if product.PRODUCT_TYPE == "simple" {
		insertSku, statusCode, err, jsonDataToM2 := magentoServiceProduct.CreateSimpleProduct(product, "", 4)
		if err != nil {
			// err to string
			status = "ERROR"
			msg = get_msg_from_json(string(err.Error()))
		} else {
			status = "SUCCESS"
			msg = "Add simple product success"
		}

		code = statusCode
		data = insertSku
		jsonData = jsonDataToM2
	} else {
		errType := errors.New("ERROR, PRODUCT_TYPE data not correct")
		status = "ERROR"
		code = 400
		jsonData = "nil"
		data = "nil"
		msg = get_msg_from_json(string(errType.Error()))
	}

	// add image data
	if code == 200 {
		imageFiles := []string{
			product.PIC_FILE,
			product.PIC_FILE2,
			product.PIC_FILE3,
			product.PIC_FILE4,
			product.PIC_FILE5,
		}

		var imageRequests []request.CreateImageRequest
		for _, imageFile := range imageFiles {
			if imageFile != "" {
				imageRequest := request.CreateImageRequest{
					Sku:         product.PROD_CODE,
					ProductType: product.PRODUCT_TYPE,
					SyncDate:    time.Now(),
					Image:       imageFile,
				}

				imageRequests = append(imageRequests, imageRequest)
			}
		}

		if len(imageRequests) > 0 {
			imgService.CreateImage(imageRequests)
		}
	}

	return code, status, msg, jsonData, data
}

func update_product_to_m2(product request.AddUpdateProductRequest, service service.ProductService, imgService service.ImageService) (int, string, string, string, string) {
	// update product to magento
	var status, msg, jsonData, data string
	var code int

	if product.PRODUCT_TYPE == "configurable" {
		insertSku, statusCode, err, jsonDataToM2 := magentoServiceProduct.UpdateSimpleProduct(product, "", "All", 1, "not-update-stock")
		insertSkuTH, statusCodeTH, errTH, jsonDataToM2TH := magentoServiceProduct.UpdateSimpleProduct(product, "", "TH", 1, "not-update-stock")
		if errTH != nil {
			fmt.Println("errTH: ", statusCodeTH, errTH, insertSkuTH)
		}

		if err != nil {
			// err to string
			status = "ERROR"
			msg = get_msg_from_json(string(err.Error()))
		} else {
			status = "SUCCESS"
			msg = "Update configurable product success"
		}

		code = statusCode
		data = insertSku

		// update to master configurable product
		findMaster := service.FindBySku(product.PROD_CODE)
		var configJsonEN, configJsonTH string
		if findMaster.ID != 0 {
			updateMaster, codeMaster, errMaster, jsonConfigDataToM2 := magentoServiceProduct.UpdateConfigurableProduct(product, findMaster.Sku, "", "All", 4)
			updateMasterTH, codeMasterTH, errMasterTH, jsonConfigDataToM2TH := magentoServiceProduct.UpdateConfigurableProduct(product, findMaster.Sku, "", "TH", 4)
			if errMaster != nil && errMasterTH != nil {
				fmt.Println("updateMaster: ", codeMaster, codeMasterTH, updateMaster, updateMasterTH)
			}

			configJsonEN = jsonConfigDataToM2
			configJsonTH = jsonConfigDataToM2TH
		}

		jsonDataArr := fmt.Sprintf(`{"simple-en": %s, "simple-th": %s, "config-en": %s, "config-th": %s}`, jsonDataToM2, jsonDataToM2TH, configJsonEN, configJsonTH)

		jsonData = jsonDataArr

	} else if product.PRODUCT_TYPE == "simple" {
		// update product
		insertSku, statusCode, err, jsonDataToM2 := magentoServiceProduct.UpdateSimpleProduct(product, "", "All", 4, "not-update-stock")
		insertSkuTH, statusCodeTH, errTH, jsonDataToM2TH := magentoServiceProduct.UpdateSimpleProduct(product, "", "TH", 4, "not-update-stock")
		if errTH != nil {
			fmt.Println("errTH: ", statusCodeTH, errTH, insertSkuTH)
		}

		if err != nil {
			// err to string
			status = "ERROR"
			msg = get_msg_from_json(string(err.Error()))
		} else {
			status = "SUCCESS"
			msg = "Update simple product success"
		}

		code = statusCode
		data = insertSku

		jsonDataArr := fmt.Sprintf(`{"simple-en": %s, "simple-th": %s}`, jsonDataToM2, jsonDataToM2TH)

		jsonData = jsonDataArr
	} else {
		errType := errors.New("ERROR, PRODUCT_TYPE data not correct")
		status = "ERROR"
		code = 400
		jsonData = "nil"
		data = "nil"
		msg = get_msg_from_json(string(errType.Error()))
	}

	// add image data
	if code == 200 {
		imageFiles := []string{
			product.PIC_FILE,
			product.PIC_FILE2,
			product.PIC_FILE3,
			product.PIC_FILE4,
			product.PIC_FILE5,
		}

		var imageRequests []request.CreateImageRequest
		for _, imageFile := range imageFiles {
			if imageFile != "" {
				imageRequest := request.CreateImageRequest{
					Sku:         product.PROD_CODE,
					ProductType: product.PRODUCT_TYPE,
					SyncDate:    time.Now(),
					Image:       imageFile,
				}

				imageRequests = append(imageRequests, imageRequest)
			}
		}

		if len(imageRequests) > 0 {
			// delete image by sku
			imgService.DeleteImageBySku(product.PROD_CODE)
			// add new image
			imgService.CreateImage(imageRequests)
		}
	}

	return code, status, msg, jsonData, data
}

func update_stock_to_m2(stock request.UpdateStockRequest) (int, string, string, string, string) {
	// update stock to magento
	var status, msg, jsonData, data string
	var code int

	// get current stock qty
	stockQty, codeQty, errQty := get_product_by_sku(stock.PROD_CODE)
	if errQty != nil {
		status = "ERROR"
		code = codeQty
		msg = get_msg_from_json(string(errQty.Error()))
		jsonData = "nil"
		data = "nil"
		return code, status, msg, jsonData, data
	}

	var getProductPayload payload.SimpleProductPayload
	errParse := json.Unmarshal([]byte(stockQty), &getProductPayload)
	if errParse != nil {
		return 500, "ERROR", "Error parse simple payload", "nil", "nil"
	}

	// init new qty
	// oldStockQty, err := strconv.Atoi(stockQty)
	// if err != nil {
	// 	oldStockQty = 0
	// }

	oldStockQty := getProductPayload.ExtensionAttributes.StockItem.Qty
	newStockQty, err := strconv.Atoi(stock.StockQty)
	if err != nil {
		newStockQty = 0
	}

	updateStockQty := newStockQty + oldStockQty
	var isInStock bool
	if updateStockQty > 0 {
		isInStock = true
	} else {
		isInStock = false
	}

	requestData := request.StockItemsRequest{
		StockItem: request.StockItem{
			ItemId:                         getProductPayload.ExtensionAttributes.StockItem.ItemId,
			ProductId:                      getProductPayload.ExtensionAttributes.StockItem.ProductId,
			StockId:                        getProductPayload.ExtensionAttributes.StockItem.StockId,
			Qty:                            updateStockQty,
			IsInStock:                      isInStock,
			IsQtyDecimal:                   getProductPayload.ExtensionAttributes.StockItem.IsQtyDecimal,
			ShowDefaultNotificationMessage: getProductPayload.ExtensionAttributes.StockItem.ShowDefaultNotificationMessage,
			UseConfigMinQty:                getProductPayload.ExtensionAttributes.StockItem.UseConfigMinQty,
			MinQty:                         getProductPayload.ExtensionAttributes.StockItem.MinQty,
			UseConfigMinSaleQty:            getProductPayload.ExtensionAttributes.StockItem.UseConfigMinSaleQty,
			MinSaleQty:                     getProductPayload.ExtensionAttributes.StockItem.MinSaleQty,
			UseConfigMaxSaleQty:            getProductPayload.ExtensionAttributes.StockItem.UseConfigMaxSaleQty,
			MaxSaleQty:                     getProductPayload.ExtensionAttributes.StockItem.MaxSaleQty,
			UseConfigBackorders:            getProductPayload.ExtensionAttributes.StockItem.UseConfigBackorders,
			Backorders:                     getProductPayload.ExtensionAttributes.StockItem.Backorders,
			UseConfigNotifyStockQty:        getProductPayload.ExtensionAttributes.StockItem.UseConfigNotifyStockQty,
			NotifyStockQty:                 getProductPayload.ExtensionAttributes.StockItem.NotifyStockQty,
			UseConfigQtyIncrements:         getProductPayload.ExtensionAttributes.StockItem.UseConfigQtyIncrements,
			QtyIncrements:                  getProductPayload.ExtensionAttributes.StockItem.QtyIncrements,
			UseConfigEnableQtyInc:          getProductPayload.ExtensionAttributes.StockItem.UseConfigEnableQtyInc,
			EnableQtyIncrements:            getProductPayload.ExtensionAttributes.StockItem.EnableQtyIncrements,
			UseConfigManageStock:           getProductPayload.ExtensionAttributes.StockItem.UseConfigManageStock,
			ManageStock:                    getProductPayload.ExtensionAttributes.StockItem.ManageStock,
			LowStockDate:                   getProductPayload.ExtensionAttributes.StockItem.LowStockDate,
			IsDecimalDivided:               getProductPayload.ExtensionAttributes.StockItem.IsDecimalDivided,
			StockStatusChangedAuto:         getProductPayload.ExtensionAttributes.StockItem.StockStatusChangedAuto,
		},
	}

	// update stock
	updateStockData, statusCode, err := magentoServiceInventory.UpdateStockItems(requestData, getProductPayload.Sku, "")

	if err != nil {
		// err to string
		status = "ERROR"
		msg = get_msg_from_json(string(err.Error()))
	} else {
		status = "SUCCESS"
		msg = "Update stock data success"
	}

	code = statusCode
	data = updateStockData

	jsonToMagento, _ := json.Marshal(requestData)
	jsonData = string(jsonToMagento)

	return code, status, msg, jsonData, data
}

func update_store_to_m2(store request.UpdateStoreRequest) (int, string, string, string, string) {
	// update stock to magento
	var status, msg, jsonData, data string
	var code int
	var jsonStr string

	if store.StockQty == "" {
		// update empty product_store to magento
		jsonStr = `{"product": {
      "custom_attributes": [
        {
          "attribute_code": "product_store",
          "value": ""
        }
      ]
    }}`
	} else {
		// get options product_store id
		attrOpt, codeAttrOpt, errAttrOpt := magentoServiceAttribute.GetAttributeOptionByCode("", "all", "product_store")
		if errAttrOpt != nil {
			return codeAttrOpt, "ERROR", get_msg_from_json(string(errAttrOpt.Error())), "nil", "nil"
		}

		var attrOptPayload payload.ProductAttributeOptionPayload
		errParseAttr := json.Unmarshal([]byte(attrOpt), &attrOptPayload)
		if errParseAttr != nil {
			return 400, "ERROR", get_msg_from_json(string(errParseAttr.Error())), "nil", "nil"
		}

		splitArray := strings.Split(store.StockQty, ",")
		var storeArray []string
		for _, v := range splitArray {
			for _, v2 := range attrOptPayload {
				if v == v2.Label {
					storeArray = append(storeArray, v2.Value)
				}
			}
		}

		jsonStr = `{"product": {
		  "custom_attributes": [
		    {
		      "attribute_code": "product_store",
		      "value": "` + strings.Join(storeArray, ",") + `"
		    }
		  ]
		}}`
	}

	// update store
	updateStoreData, statusCode, err := magentoServiceInventory.UpdateProductStore(store, "", jsonStr)

	if err != nil {
		// err to string
		status = "ERROR"
		msg = get_msg_from_json(string(err.Error()))
	} else {
		status = "SUCCESS"
		msg = "Update store data success"
	}

	code = statusCode
	data = updateStoreData
	jsonData = jsonStr

	return code, status, msg, jsonData, data
}

func update_image_process(image request.CreateImageQueueRequest, productService service.ProductService, imgService service.ImageService) (int, string, string, string, string) {
	// check image exist in folder
	bucket := os.Getenv("AWS_S3_BUCKET")
	// baseDirectory := os.Getenv("UPLOAD_PATH")
	imageKey := fmt.Sprintf("%s/%s", image.DirectoryPath, image.Image)
	// Create an S3 client
	s3Client := helpers.SetupS3Client()
	// Get the file from S3
	fileContent, mimeType, err := helpers.GetFileAsBase64(s3Client, bucket, imageKey)
	if err != nil {
		return 400, "ERROR", "Image not found", "nil", "nil"
	}

	// Encode the content as base64
	base64Content := base64.StdEncoding.EncodeToString(fileContent)

	if base64Content == "" {
		return 400, "ERROR", "Error encoding image to base64", "nil", "nil"
	}

	// get image sku
	product, err := imgService.FindImageSkuByName(image.Image)

	if err != nil {
		fmt.Println("Error get image sku:", err)
		// deleteMedia, codeMedia, errMedia := magentoServiceMedia.DeleteMedia("", mediaApiPath)
		// if codeMedia == 200 {
		// 	fmt.Println("Success removing media:", mediaApiPath, deleteMedia)
		// } else {
		// 	fmt.Println("Error removing media:", mediaApiPath, errMedia)
		// }

		// Delete the file in s3
		err = helpers.DeleteFile(s3Client, bucket, imageKey)
		fmt.Println("Error deleting file:", err)
		// return error
		return 400, "ERROR", "Product SKU not found", "nil", "nil"
	}

	// update image to simple product
	code, status, msg, jsonStr, data := update_image_to_m2(image, base64Content, mimeType, product.Sku)

	// update image to master configurable product
	findMaster := productService.FindBySku(product.Sku)
	if findMaster.ID != 0 {
		code, status, msg, jsonStr, data = update_image_to_m2(image, base64Content, mimeType, findMaster.Sku)
	}

	// remove image in folder
	if code == 200 {
		// deleteMedia, codeMedia, errMedia := magentoServiceMedia.DeleteMedia("", mediaApiPath)
		// if codeMedia == 200 {
		// 	fmt.Println("Success removing media:", mediaApiPath, deleteMedia)
		// } else {
		// 	fmt.Println("Error removing media:", mediaApiPath, errMedia)
		// }

		// Delete the file in s3
		err = helpers.DeleteFile(s3Client, bucket, imageKey)
		fmt.Println("Error deleting file:", err)
	}

	return code, status, msg, jsonStr, data
}

func update_image_to_m2(image request.CreateImageQueueRequest, base64File, mimeTypeFile, sku string) (int, string, string, string, string) {
	var status, msg, data string
	var code int

	getMedia, codeMedia, errMedia := magentoServiceMedia.GetMediaBySKU("", sku)

	var jsonStr string

	if codeMedia == 200 {
		var mediaPayload payload.MediaPayload
		errParse := json.Unmarshal([]byte(getMedia), &mediaPayload)
		if errParse != nil {
			return 400, "ERROR", "Error parse media payload", "nil", "nil"
		}

		// check if image name exist in magento
		var imageExist int = 0
		var updatePayload request.UpdateMediaRequest
		for _, media := range mediaPayload {
			if media.Label == image.Image {
				imageExist = media.ID

				updatePayload = request.UpdateMediaRequest{
					Entry: request.UpdateMediaEntryRequest{
						ID:        media.ID,
						MediaType: media.MediaType,
						Label:     image.Image,
						Position:  media.Position,
						Disabled:  media.Disabled,
						Types:     media.Types,
						Content: request.CreateMediaContentRequest{
							Base64EncodedData: base64File,
							Type:              mimeTypeFile,
							Name:              slug.Make(image.Image),
						},
					},
				}
			}
		}

		if imageExist == 0 {
			// add image to magento
			requestData := request.CreateMediaRequest{
				Entry: request.CreateMediaEntryRequest{
					MediaType: "image",
					Label:     image.Image,
					Position:  len(mediaPayload) + 1,
					Disabled:  false,
					Content: request.CreateMediaContentRequest{
						Base64EncodedData: base64File,
						Type:              mimeTypeFile,
						Name:              slug.Make(image.Image),
					},
				},
			}

			addMedia, codeAddMedia, errAddMedia := magentoServiceMedia.CreateMedia("", sku, requestData)
			if errAddMedia != nil {
				return codeAddMedia, "ERROR", get_msg_from_json(string(errAddMedia.Error())), "nil", "nil"
			}

			code = codeAddMedia
			status = "SUCCESS"
			data = addMedia
			msg = "Add image success"

			// requestData to json
			jsonObjStr, _ := json.Marshal(requestData)
			jsonStr = string(jsonObjStr)
		} else {
			// update image to magento
			updateMedia, codeUpdateMedia, errUpdateMedia := magentoServiceMedia.UpdateMedia("", sku, imageExist, updatePayload)
			if errUpdateMedia != nil {
				return codeUpdateMedia, "ERROR", get_msg_from_json(string(errUpdateMedia.Error())), "nil", "nil"
			}

			code = codeUpdateMedia
			status = "SUCCESS"
			data = updateMedia
			msg = "Update image success"

			jsonObjStr, _ := json.Marshal(updatePayload)
			jsonStr = string(jsonObjStr)
		}
	} else {
		return codeMedia, "ERROR", get_msg_from_json(string(errMedia.Error())), "nil", "nil"
	}

	return code, status, msg, jsonStr, data
}

func get_msg_from_json(jsonStr string) string {
	var postMsg string
	if helpers.IsJSONString(jsonStr) {
		var response StatusResponse
		errParseMsg := json.Unmarshal([]byte(jsonStr), &response)
		if errParseMsg != nil {
			postMsg = "Error, something went wrong"
		}

		postMsg = response.Message
	} else {
		postMsg = jsonStr
	}

	return postMsg
}

func get_product_by_sku(sku string) (string, int, error) {
	fmt.Println("GetCurrentStockQty :: SKU => ", sku)
	getProduct, codeProduct, errProduct := magentoServiceProduct.GetProductBySKU("", sku)
	if errProduct != nil {
		fmt.Printf("GetCurrentStockQty :: ERROR => %s , %d, %s\n", sku, codeProduct, errProduct)
	} else {
		fmt.Printf("GetCurrentStockQty :: SUCCESS => %s\n", sku)
	}

	return getProduct, codeProduct, errProduct
}

func get_data_from_erp(service service.QueueService) (int, string, string, string) {
	var serviceURL, serviceMethod string
	if config.EnableTestMode() == "development" {
		serviceURL = os.Getenv("STRAPI_URL") + "/main-job-proc-data-api"
		serviceMethod = "GET"
	} else {
		serviceURL = os.Getenv("SERVICE_URL") + "/MainJobProcDataApi"
		serviceMethod = "POST"
	}

	req, err := http.NewRequest(serviceMethod, serviceURL, nil)
	if err != nil {
		return 500, "ERROR", get_msg_from_json(string(err.Error())), "nil"
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return 500, "ERROR", get_msg_from_json(string(err.Error())), "nil"
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 500, "ERROR", get_msg_from_json(string(err.Error())), "nil"
	}

	// add main job proc data
	connection := service.Connection(responseBody)

	// add queue
	addTotal, err := strconv.Atoi(connection.TotalRecordAdd)
	if err == nil && addTotal > 0 {
		// request to erp system [GetNewProdInfoApi]
		var serviceURL, serviceMethod string

		if config.EnableTestMode() == "development" {
			serviceURL = os.Getenv("STRAPI_URL") + "/get-new-prod-info-api"
			serviceMethod = "GET"
		} else {
			serviceURL = os.Getenv("SERVICE_URL") + "/GetNewProdInfoApi"
			serviceMethod = "POST"
		}

		jsonStr := []byte(`{"TotalRecord":"` + connection.TotalRecordAdd + `"}`)
		req, err := http.NewRequest(serviceMethod, serviceURL, bytes.NewBuffer(jsonStr))
		helpers.ErrorPanic(err)

		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			return 500, "ERROR", get_msg_from_json(string(err.Error())), "nil"
		}

		defer resp.Body.Close()

		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return 500, "ERROR", get_msg_from_json(string(err.Error())), "nil"
		}

		service.CreateProductsQueue("ADD", responseBody)
	}

	// update queue
	updateTotal, err := strconv.Atoi(connection.TotalRecordUpdate)
	if err == nil && updateTotal > 0 {
		// request to erp system [GetUpdateProdInfoApi]
		var serviceURL, serviceMethod string

		if config.EnableTestMode() == "development" {
			serviceURL = os.Getenv("STRAPI_URL") + "/get-update-prod-info-api"
			serviceMethod = "GET"
		} else {
			serviceURL = os.Getenv("SERVICE_URL") + "/GetUpdateProdInfoApi"
			serviceMethod = "POST"
		}

		jsonStr := []byte(`{"TotalRecord":"` + connection.TotalRecordUpdate + `"}`)
		req, err := http.NewRequest(serviceMethod, serviceURL, bytes.NewBuffer(jsonStr))
		helpers.ErrorPanic(err)

		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			return 500, "ERROR", get_msg_from_json(string(err.Error())), "nil"
		}

		defer resp.Body.Close()

		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return 500, "ERROR", get_msg_from_json(string(err.Error())), "nil"
		}

		service.CreateProductsQueue("UPDATE", responseBody)
	}

	// stock queue
	stockTotal, err := strconv.Atoi(connection.TotalRecordStock)
	if err == nil && stockTotal > 0 {
		// request to erp system [GetUpdateStockApi]
		var serviceURL, serviceMethod string

		if config.EnableTestMode() == "development" {
			serviceURL = os.Getenv("STRAPI_URL") + "/get-update-stock-api"
			serviceMethod = "GET"
		} else {
			serviceURL = os.Getenv("SERVICE_URL") + "/GetUpdateStockApi"
			serviceMethod = "POST"
		}

		jsonStr := []byte(`{"TotalRecord":"` + connection.TotalRecordStock + `"}`)
		req, err := http.NewRequest(serviceMethod, serviceURL, bytes.NewBuffer(jsonStr))
		helpers.ErrorPanic(err)

		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			return 500, "ERROR", get_msg_from_json(string(err.Error())), "nil"
		}

		defer resp.Body.Close()

		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return 500, "ERROR", get_msg_from_json(string(err.Error())), "nil"
		}

		service.CreateStockQueue(responseBody)
	}

	// store queue
	storeTotal, err := strconv.Atoi(connection.TotalRecordStore)
	if err == nil && storeTotal > 0 {
		// request to erp system [GetUpdateStoreApi]
		var serviceURL, serviceMethod string

		if config.EnableTestMode() == "development" {
			serviceURL = os.Getenv("STRAPI_URL") + "/get-update-store-api"
			serviceMethod = "GET"
		} else {
			serviceURL = os.Getenv("SERVICE_URL") + "/GetUpdateStoreApi"
			serviceMethod = "POST"
		}

		jsonStr := []byte(`{"TotalRecord":"` + connection.TotalRecordStore + `"}`)
		req, err := http.NewRequest(serviceMethod, serviceURL, bytes.NewBuffer(jsonStr))
		helpers.ErrorPanic(err)

		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			return 500, "ERROR", get_msg_from_json(string(err.Error())), "nil"
		}

		defer resp.Body.Close()

		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return 500, "ERROR", get_msg_from_json(string(err.Error())), "nil"
		}

		service.CreateStoreQueue(responseBody)
	}

	return 200, "Ok", "", "nil"
}

func send_post_flag_queue(code int, id, msg string, service service.QueueService) {
	var postMsg string
	if helpers.IsJSONString(msg) {
		var response StatusResponse
		errParseMsg := json.Unmarshal([]byte(msg), &response)
		if errParseMsg != nil {
			postMsg = "Error, something went wrong"
		}

		postMsg = response.Message
	} else {
		postMsg = msg
	}

	if code == 200 {
		service.CreatePostflagQueue(id, "UPDATE-COMPLETED", postMsg)
	} else {
		service.CreatePostflagQueue(id, "UPDATE-UN-COMPLETED", postMsg)
	}
}

func send_post_flag_to_erp(postflag request.CreatePostflagRequest) (int, string, string, string, string) {
	var status, msg, data string
	var code int

	// request to erp system [PostUpdateSendFlagApi]
	serviceURL := os.Getenv("SERVICE_URL") + "/PostUpdateSendFlagApi"
	requestData := request.CreatePostflagRequest{
		TransactionId: postflag.TransactionId,
		FlagStatus:    postflag.FlagStatus,
		ErrMsg:        postflag.ErrMsg,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return 500, "ERROR", get_msg_from_json(string(err.Error())), "nil", "nil"
	}

	req, err := http.NewRequest("POST", serviceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return 500, "ERROR", get_msg_from_json(string(err.Error())), "nil", "nil"
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return 500, "ERROR", get_msg_from_json(string(err.Error())), "nil", "nil"
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 500, "ERROR", get_msg_from_json(string(err.Error())), "nil", "nil"
	}

	code = resp.StatusCode
	status = "SUCCESS"
	msg = "Send Postflag Success."
	data = string(responseBody)

	return code, status, msg, string(jsonData), data
}

func send_dailysale_to_erp(dailysale request.CreateDailySalesRequest) (int, string, string, string, string) {
	var status, msg, data string
	var code int

	// request to erp system [PostUpdateSendFlagApi]
	serviceURL := os.Getenv("SERVICE_URL") + "/CreateDailySaleApi"
	jsonData, err := json.Marshal(dailysale)
	if err != nil {
		return 500, "ERROR", get_msg_from_json(string(err.Error())), "nil", "nil"
	}

	req, err := http.NewRequest("POST", serviceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return 500, "ERROR", get_msg_from_json(string(err.Error())), "nil", "nil"
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return 500, "ERROR", get_msg_from_json(string(err.Error())), "nil", "nil"
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 500, "ERROR", get_msg_from_json(string(err.Error())), "nil", "nil"
	}

	code = resp.StatusCode
	status = "SUCCESS"
	msg = "Send Dailysale Success."
	data = string(responseBody)

	// dailysale to json string
	dailysaleJson, err := json.Marshal(dailysale)

	return code, status, msg, string(dailysaleJson), data
}
