package service

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/captainkie/websync-api/config"
	model "github.com/captainkie/websync-api/internal/app/models"
	"github.com/captainkie/websync-api/internal/app/repository"
	"github.com/captainkie/websync-api/internal/app/utils"
	"github.com/captainkie/websync-api/pkg/helpers"
	"github.com/captainkie/websync-api/types/request"
	"github.com/captainkie/websync-api/types/response"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type QueueServiceImpl struct {
	QueueRepository repository.QueueRepository
	Validate        *validator.Validate
}

func NewQueueServiceImpl(queueRepository repository.QueueRepository, validate *validator.Validate) QueueService {
	return &QueueServiceImpl{
		QueueRepository: queueRepository,
		Validate:        validate,
	}
}

// Create implements QueueService interface
func (q *QueueServiceImpl) Connection(connection []byte) response.CreateConnectionResponse {
	// Pretty-print the JSON response
	// helpers.PrintPrettyJson(connection)

	// Unmarshal the JSON data into the ResponseData struct
	var body request.CreateConnectionRequest
	if config.EnableTestMode() == "development" {
		var response request.StrapiConnectionRequest
		err := json.Unmarshal(connection, &response)
		helpers.ErrorPanic(err)

		body = response.Data.Attributes.Body
	} else {
		var response request.CreateConnectionRequest
		err := json.Unmarshal(connection, &response)
		helpers.ErrorPanic(err)

		body = response
	}

	// save connection to database
	newConnection := model.Connections{
		MessageCode:       body.MessageCode,
		MessageDesc:       body.MessageDesc,
		TotalRecordAdd:    body.TotalRecordAdd,
		TotalRecordUpdate: body.TotalRecordUpdate,
		TotalRecordStock:  body.TotalRecordStock,
		TotalRecordStore:  body.TotalRecordStore,
		SyncDate:          time.Now(),
	}

	connectionData, err := q.QueueRepository.Connection(newConnection)
	helpers.ErrorPanic(err)

	connectionResponse := response.CreateConnectionResponse{
		ID:                connectionData.ID,
		MessageCode:       connectionData.MessageCode,
		MessageDesc:       connectionData.MessageDesc,
		TotalRecordAdd:    connectionData.TotalRecordAdd,
		TotalRecordUpdate: connectionData.TotalRecordUpdate,
		TotalRecordStock:  connectionData.TotalRecordStock,
		TotalRecordStore:  connectionData.TotalRecordStore,
		Created:           connectionData.Created,
		Updated:           connectionData.Updated,
	}

	return connectionResponse
}

func (q *QueueServiceImpl) CreateConnectionQueue(id string) {
	var newConnection []model.ConnectionQueues
	newConnection = append(newConnection, model.ConnectionQueues{
		TransactionID: id,
	})

	connectionData, err := q.QueueRepository.CreateConnection(newConnection)
	helpers.ErrorPanic(err)

	// Simulate adding tasks to the queues
	utils.ConnectionTask(connectionData)
}

// Update implements QueueService interface
func (q *QueueServiceImpl) UpdateConnectionQueue(id int, status string) {
	// save connection to database
	updateConnection := model.ConnectionQueues{
		ID:     id,
		Status: status,
	}

	q.QueueRepository.UpdateConnection(updateConnection)
}

// CreateProductsQueue implements QueueService interface
func (q *QueueServiceImpl) CreateProductsQueue(qtype string, products []byte) {
	// Unmarshal the JSON data into the ResponseData struct
	var body []request.AddUpdateProductRequest
	if config.EnableTestMode() == "development" {
		var response request.StrapiAddUpdateProductRequest
		err := json.Unmarshal(products, &response)
		helpers.ErrorPanic(err)

		for _, value := range response.Data.Attributes.Body {
			body = append(body, value)
		}
	} else {
		var response []request.AddUpdateProductRequest
		err := json.Unmarshal(products, &response)
		helpers.ErrorPanic(err)

		body = response
	}

	if qtype == "ADD" {
		// save products to database
		var newProducts []model.AddQueues
		for _, value := range body {
			jsonData, err := json.Marshal(value)
			helpers.ErrorPanic(err)
			TransactionID, err := strconv.Atoi(value.TRANS_ID)
			helpers.ErrorPanic(err)

			newProducts = append(newProducts, model.AddQueues{
				TransactionID: TransactionID,
				JsonData:      string(jsonData),
			})
		}

		productData, err := q.QueueRepository.CreateTypeAdd(newProducts)
		helpers.ErrorPanic(err)

		// Simulate adding tasks to the queues
		utils.AddProductsTask(productData)
	} else if qtype == "UPDATE" {
		// save products to database
		var newProducts []model.UpdateQueues
		for _, value := range body {
			jsonData, err := json.Marshal(value)
			helpers.ErrorPanic(err)
			TransactionID, err := strconv.Atoi(value.TRANS_ID)
			helpers.ErrorPanic(err)

			newProducts = append(newProducts, model.UpdateQueues{
				TransactionID: TransactionID,
				JsonData:      string(jsonData),
			})
		}

		productData, err := q.QueueRepository.CreateTypeUpdate(newProducts)
		helpers.ErrorPanic(err)

		// Simulate adding tasks to the queues
		utils.UpdateProductsTask(productData)
	}
}

// UpdateProductsQueue implements QueueService interface
func (q *QueueServiceImpl) UpdateProductsQueue(id int, qtype, status string, products []byte) {
	// Unmarshal the JSON data into the ResponseData struct
	var body request.AddUpdateProductRequest
	if config.EnableTestMode() == "development" {
		var response request.StrapiAddUpdateProductRequest
		err := json.Unmarshal(products, &response)
		helpers.ErrorPanic(err)

		for _, value := range response.Data.Attributes.Body {
			body = value
		}
	} else {
		var response request.AddUpdateProductRequest
		err := json.Unmarshal(products, &response)
		helpers.ErrorPanic(err)

		body = response
	}

	if qtype == "ADD" {
		// save products to database
		var updateProduct model.AddQueues
		jsonData, err := json.Marshal(body)
		helpers.ErrorPanic(err)
		updateProduct = model.AddQueues{
			ID:       id,
			Status:   status,
			JsonData: string(jsonData),
		}

		q.QueueRepository.UpdateTypeAdd(updateProduct)
	} else if qtype == "UPDATE" {
		// save products to database
		var updateProduct model.UpdateQueues
		jsonData, err := json.Marshal(body)
		helpers.ErrorPanic(err)
		updateProduct = model.UpdateQueues{
			ID:       id,
			Status:   status,
			JsonData: string(jsonData),
		}

		q.QueueRepository.UpdateTypeUpdate(updateProduct)
	}
}

// CreateStockQueue implements QueueService interface
func (q *QueueServiceImpl) CreateStockQueue(stocks []byte) {
	// Unmarshal the JSON data into the ResponseData struct
	var body []request.UpdateStockRequest
	if config.EnableTestMode() == "development" {
		var response request.StrapiUpdateStockRequest
		err := json.Unmarshal(stocks, &response)
		helpers.ErrorPanic(err)

		for _, value := range response.Data.Attributes.Body {
			body = append(body, value)
		}
	} else {
		var response []request.UpdateStockRequest
		err := json.Unmarshal(stocks, &response)
		helpers.ErrorPanic(err)

		body = response
	}

	// save stocks to database
	var newStock []model.StockQueues
	for _, value := range body {
		jsonData, err := json.Marshal(value)
		helpers.ErrorPanic(err)
		newStock = append(newStock, model.StockQueues{
			TransactionID: value.TRANS_ID,
			JsonData:      string(jsonData),
		})
	}

	stockData, err := q.QueueRepository.CreateStock(newStock)
	helpers.ErrorPanic(err)

	// Simulate adding tasks to the queues
	utils.StockTask(stockData)
}

// UpdateStockQueue implements QueueService interface
func (q *QueueServiceImpl) UpdateStockQueue(id int, status string, stocks []byte) {
	// Unmarshal the JSON data into the ResponseData struct
	var body request.UpdateStockRequest
	if config.EnableTestMode() == "development" {
		var response request.StrapiUpdateStockRequest
		err := json.Unmarshal(stocks, &response)
		helpers.ErrorPanic(err)

		for _, value := range response.Data.Attributes.Body {
			body = value
		}
	} else {
		var response request.UpdateStockRequest
		err := json.Unmarshal(stocks, &response)
		helpers.ErrorPanic(err)

		body = response
	}

	// save stocks to database
	var updateStock model.StockQueues
	jsonData, err := json.Marshal(body)
	helpers.ErrorPanic(err)
	updateStock = model.StockQueues{
		ID:       id,
		Status:   status,
		JsonData: string(jsonData),
	}

	q.QueueRepository.UpdateStock(updateStock)
}

// CreateStoreQueue implements QueueService interface
func (q *QueueServiceImpl) CreateStoreQueue(stores []byte) {
	// Unmarshal the JSON data into the ResponseData struct
	var body []request.UpdateStoreRequest
	if config.EnableTestMode() == "development" {
		var response request.StrapiUpdateStoreRequest
		err := json.Unmarshal(stores, &response)
		helpers.ErrorPanic(err)

		for _, value := range response.Data.Attributes.Body {
			body = append(body, value)
		}
	} else {
		var response []request.UpdateStoreRequest
		err := json.Unmarshal(stores, &response)
		helpers.ErrorPanic(err)

		body = response
	}

	// save stores to database
	var createStore []model.StoreQueues
	for _, value := range body {
		jsonData, err := json.Marshal(value)
		helpers.ErrorPanic(err)
		createStore = append(createStore, model.StoreQueues{
			TransactionID: value.TRANS_ID,
			JsonData:      string(jsonData),
		})
	}

	storeData, err := q.QueueRepository.CreateStore(createStore)
	helpers.ErrorPanic(err)

	// Simulate adding tasks to the queues
	utils.StoreTask(storeData)
}

// UpdateStoreQueue implements QueueService interface
func (q *QueueServiceImpl) UpdateStoreQueue(id int, status string, stores []byte) {
	// Unmarshal the JSON data into the ResponseData struct
	var body request.UpdateStoreRequest
	if config.EnableTestMode() == "development" {
		var response request.StrapiUpdateStoreRequest
		err := json.Unmarshal(stores, &response)
		helpers.ErrorPanic(err)

		for _, value := range response.Data.Attributes.Body {
			body = value
		}
	} else {
		var response request.UpdateStoreRequest
		err := json.Unmarshal(stores, &response)
		helpers.ErrorPanic(err)

		body = response
	}

	// save stores to database
	var updateStore model.StoreQueues
	jsonData, err := json.Marshal(body)
	helpers.ErrorPanic(err)
	updateStore = model.StoreQueues{
		ID:       id,
		Status:   status,
		JsonData: string(jsonData),
	}

	q.QueueRepository.UpdateStore(updateStore)
}

// CreatePostflagQueue implements QueueService interface
func (q *QueueServiceImpl) CreatePostflagQueue(id, status, message string) {
	// Unmarshal the JSON data into the ResponseData struct
	transID, _ := strconv.Atoi(id)
	request := request.CreatePostflagRequest{
		TransactionId: transID,
		FlagStatus:    status,
		ErrMsg:        helpers.ReplaceAllQuot(message),
	}
	// Encode the struct to a JSON string
	jsonData, _ := json.Marshal(request)

	// Convert the JSON byte slice to a string
	jsonString := string(jsonData)
	// save stocks to database
	var newPostflag []model.PostflagQueues
	newPostflag = append(newPostflag, model.PostflagQueues{
		TransactionID: transID,
		JsonData:      jsonString,
	})

	postflagData, _ := q.QueueRepository.CreatePostflag(newPostflag)

	// Simulate adding tasks to the queues
	utils.PostflagTask(postflagData)
}

// UpdatePostflagQueue implements QueueService interface
func (q *QueueServiceImpl) UpdatePostflagQueue(id int, status string) {
	//update postflag to database
	updatePostflag := model.PostflagQueues{
		ID:     id,
		Status: status,
	}

	q.QueueRepository.UpdatePostflag(updatePostflag)
}

// CreateImageQueue implements QueueService interface
func (q *QueueServiceImpl) CreateImageQueue() {
	// Init variables
	// Get the current date in the format YYYYMMDD.
	currentDate := time.Now().Format("20060102")
	// bucket name
	bucket := os.Getenv("AWS_S3_BUCKET")
	// Create an S3 client
	s3Client := helpers.SetupS3Client()
	// List all objects with the specified prefix
	targetFolder := "sync/" + currentDate
	resp, err := helpers.ListObjects(s3Client, bucket, targetFolder)
	if err != nil {
		fmt.Println("❌ Failed to list AWS S3 items", err)
	}

	// uploadfile to s3
	// localfile := "public/uploads/sync/20240206/JP0ICE11BAA005.jpg"
	// prefix := "20240206"
	// err = helpers.UploadFile(s3Client, bucket, prefix, localfile)

	// loop images
	var newImage []model.ImageQueues
	if len(resp) > 0 {
		for _, value := range resp {
			requestID := uuid.New().String()
			newImage = append(newImage, model.ImageQueues{
				TransactionID: requestID,
				Image:         value,
				DirectoryPath: targetFolder, // s3 prefix
				SyncDate:      currentDate,
			})
		}
	}

	if len(newImage) > 0 {
		imageData, _ := q.QueueRepository.CreateImage(newImage)
		// Simulate adding tasks to the queues
		utils.ImageTask(imageData)
	}
}

// UpdateImageQueue implements QueueService interface
func (q *QueueServiceImpl) UpdateImageQueue(id int, status string) {
	// update image to database
	updateImage := model.ImageQueues{
		ID:     id,
		Status: status,
	}

	q.QueueRepository.UpdateImage(updateImage)
}

// DeleteImageQueue implements QueueService interface
func (q *QueueServiceImpl) DeleteImageQueue(id int) {
	q.QueueRepository.DeleteImage(id)
}

// CreateDailySalesQueue implements QueueService interface
func (q *QueueServiceImpl) CreateDailySalesQueue(order request.CreateDailySalesRequest) {
	var newDailySale []model.DailysaleQueues
	requestID := uuid.New().String()
	jsonData, _ := json.Marshal(order)

	newDailySale = append(newDailySale, model.DailysaleQueues{
		TransactionID: requestID,
		OrderID:       order.DocNo,
		JsonData:      string(jsonData),
	})

	dailysaleData, _ := q.QueueRepository.CreateDailySales(newDailySale)

	// Simulate adding tasks to the queues
	utils.DailysaleTask(dailysaleData)
}

// UpdateDailySaleQueue implements QueueService interface
func (q *QueueServiceImpl) UpdateDailySalesQueue(id int, status string) {
	// update image to database
	UpdateDailySale := model.DailysaleQueues{
		ID:     id,
		Status: status,
	}

	q.QueueRepository.UpdateDailySales(UpdateDailySale)
}
