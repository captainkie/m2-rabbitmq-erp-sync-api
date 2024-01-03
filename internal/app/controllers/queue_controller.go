package controller

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/captainkie/websync-api/internal/app/service"
	"github.com/captainkie/websync-api/types/request"
	"github.com/captainkie/websync-api/types/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type QueueController struct {
	queueService service.QueueService
	imageService service.ImageService
}

func NewQueueController(qService service.QueueService, imgService service.ImageService) *QueueController {
	return &QueueController{
		queueService: qService,
		imageService: imgService,
	}
}

// ProductsSync		godoc
// @Summary		ProductsSync Add,Update,Stock,Store request queue
// @Description	create new queue
// @Produce  application/json
// @tags Queue
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /queue/products [get]
// @Security BearerAuth
func (controller *QueueController) ProductsSync(ctx *gin.Context) {
	// request to erp system connection
	requestID := uuid.New().String()
	controller.queueService.CreateConnectionQueue(requestID)

	// Use the imageService to delete images older than targetDate
	targetDate := time.Now().AddDate(0, 0, -1)
	controller.imageService.DeleteImage(targetDate)

	// Delete old image folder
	targetFolder := time.Now().AddDate(0, 0, -2).Format("20060102")
	baseDirectory := os.Getenv("UPLOAD_PATH")
	directoryPath := fmt.Sprintf("%s/%s", baseDirectory, targetFolder)
	if _, err := os.Stat(directoryPath); !os.IsNotExist(err) {
		os.RemoveAll(directoryPath)
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully create queue!",
		Data:    nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// ImageSync		godoc
// @Summary		ImageSync data to magento request queue
// @Description	create new queue
// @Produce  application/json
// @tags Queue
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /queue/images [get]
// @Security BearerAuth
func (controller *QueueController) ImagesSync(ctx *gin.Context) {
	// Get the current date in the format YYYYMMDD.
	currentDate := time.Now().Format("20060102")
	// Define the base directory path.
	baseDirectory := os.Getenv("UPLOAD_PATH")
	// Create full directory path
	directoryPath := fmt.Sprintf("%s/%s", baseDirectory, currentDate)
	// Open the directory.
	dir, err := os.ReadDir(directoryPath)
	if err != nil {
		webResponse := response.Response{
			Code:    400,
			Status:  "Bad Request",
			Message: fmt.Sprintf("%s", err),
			Data:    nil,
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	// Check file exist in folder
	var fileNames []string
	for _, entry := range dir {
		if entry.IsDir() {
			// Handle directories, Don't do anything
		} else {
			// Handle files.
			fileNames = append(fileNames, entry.Name())
		}
	}

	if len(fileNames) > 0 {
		// send to queue
		// controller.queueService.CreateImageQueue(fileNames, directoryPath, currentDate)
		controller.queueService.CreateImageQueue()

		webResponse := response.Response{
			Code:    200,
			Status:  "Ok",
			Message: "Successfully create image sync queue!",
			Data:    nil,
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusOK, webResponse)

	} else {
		webResponse := response.Response{
			Code:    400,
			Status:  "Failed",
			Message: "No files found in the directory.",
			Data:    nil,
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusOK, webResponse)
	}
}

// DailySales		godoc
// @Summary		Daily Sales request queue
// @Description	create daily sales queue
// @Param    Request body request.CreateDailySalesRequest true "CreateDailySales"
// @Produce  application/json
// @tags Queue
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /queue/daily-sales [post]
// @Security BearerAuth
func (controller *QueueController) DailySales(ctx *gin.Context) {
	createDailySaleRequest := request.CreateDailySalesRequest{}
	err := ctx.ShouldBindJSON(&createDailySaleRequest)
	if err != nil {
		webResponse := response.Response{
			Code:    400,
			Status:  "Bad Request",
			Message: fmt.Sprintf("%s", err),
			Data:    nil,
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	// Validate the createDailySaleRequest
	validate := validator.New()
	if err := validate.Struct(createDailySaleRequest); err != nil {
		webResponse := response.Response{
			Code:    400,
			Status:  "Bad Request",
			Message: fmt.Sprintf("%s", err),
			Data:    nil,
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	controller.queueService.CreateDailySalesQueue(createDailySaleRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully create daily sale queue!",
		Data:    nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
