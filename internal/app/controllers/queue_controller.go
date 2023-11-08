package controller

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/captainkie/websync-api/internal/app/service"
	"github.com/captainkie/websync-api/pkg/helpers"
	"github.com/captainkie/websync-api/types/request"
	"github.com/captainkie/websync-api/types/response"
	"github.com/gin-gonic/gin"
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
func (controller *QueueController) ProductsSync(ctx *gin.Context) {
	// request to erp system connection
	requestID := uuid.New().String()
	controller.queueService.CreateConnectionQueue(requestID)

	// Use the imageService to delete images older than targetDate
	targetDate := time.Now().AddDate(0, 0, -1)
	controller.imageService.DeleteImage(targetDate)

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
func (controller *QueueController) ImagesSync(ctx *gin.Context) {
	// Get the current date in the format YYYYMMDD.
	currentDate := time.Now().Format("20060102")
	// Define the base directory path.
	baseDirectory := os.Getenv("UPLOAD_PATH")
	// Create full directory path
	directoryPath := fmt.Sprintf("%s/%s", baseDirectory, currentDate)
	// Open the directory.
	var webResponse response.Response
	dir, err := os.ReadDir(directoryPath)
	if err != nil {
		webResponse = response.Response{
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

	var msg string
	if len(fileNames) > 0 {
		// send to queue

		msg = "Successfully create image sync queue!"
	} else {
		msg = "No files found in the directory."
	}

	webResponse = response.Response{
		Code:    200,
		Status:  "Ok",
		Message: msg,
		Data:    nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
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
	helpers.ErrorPanic(err)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully create daily sale queue!",
		Data:    nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
