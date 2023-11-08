package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/captainkie/websync-api/config"
	_ "github.com/captainkie/websync-api/docs"
	controller "github.com/captainkie/websync-api/internal/app/controllers"
	model "github.com/captainkie/websync-api/internal/app/models"
	"github.com/captainkie/websync-api/internal/app/repository"
	"github.com/captainkie/websync-api/internal/app/routes"
	"github.com/captainkie/websync-api/internal/app/service"
	"github.com/captainkie/websync-api/pkg/helpers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

// @title			WebSync API
// @version		1.0.0
// @description	This is a sync service data from erp to magento 2.

// @contact.name   captainkie
// @contact.url    https://github.com/captainkie

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath  /api

// @securityDefinitions.apiKey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	// Clear && Load env
	os.Clearenv()
	err := godotenv.Load()
	helpers.ErrorPanic(err)

	go producer()

	initialize()
}

func initialize() {
	// Connect database
	db := config.ConnectDatabase()
	fmt.Println("🚀 Successfully connected to the database")

	validate := validator.New()

	// Migrate the schema
	db.Table("users").AutoMigrate(&model.Users{})
	db.Table("connections").AutoMigrate(&model.Connections{})
	db.Table("images").AutoMigrate(&model.Images{})
	db.Table("configurable_products").AutoMigrate(&model.ConfigurableProducts{})

	db.Table("connection_queues").AutoMigrate(&model.ConnectionQueues{})
	db.Table("add_queues").AutoMigrate(&model.AddQueues{})
	db.Table("update_queues").AutoMigrate(&model.UpdateQueues{})
	db.Table("stock_queues").AutoMigrate(&model.StockQueues{})
	db.Table("store_queues").AutoMigrate(&model.StoreQueues{})
	db.Table("postflag_queues").AutoMigrate(&model.PostflagQueues{})
	db.Table("image_queues").AutoMigrate(&model.ImageQueues{})
	db.Table("dailysale_queues").AutoMigrate(&model.DailysaleQueues{})

	db.Table("add_logs").AutoMigrate(&model.AddLogs{})
	db.Table("update_logs").AutoMigrate(&model.UpdateLogs{})
	db.Table("stock_logs").AutoMigrate(&model.StockLogs{})
	db.Table("store_logs").AutoMigrate(&model.StoreLogs{})
	db.Table("postflag_logs").AutoMigrate(&model.PostflagLogs{})
	db.Table("image_logs").AutoMigrate(&model.PostflagLogs{})
	db.Table("dailysale_logs").AutoMigrate(&model.DailysaleLogs{})

	// Init Repository
	userRepository := repository.NewUsersRepositoryImpl(db)
	queueRepository := repository.NewQueueRepositoryImpl(db)
	iamgeRepository := repository.NewImageRepositoryImpl(db)

	// Init Service
	authenticationService := service.NewAuthenticationServiceImpl(userRepository, validate)
	userService := service.NewUsersServiceImpl(userRepository, validate)
	queueService := service.NewQueueServiceImpl(queueRepository, validate)
	imageService := service.NewImageServiceImpl(iamgeRepository, validate)

	// Init controller
	authenticationController := controller.NewAuthenticationController(authenticationService)
	userController := controller.NewUserController(userService)
	queueController := controller.NewQueueController(queueService, imageService)

	// Set gin mode
	if os.Getenv("GIN_MODE") == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Init routes
	routes := routes.MainRouter(userRepository, authenticationController, userController, queueController)

	// Run Server
	server := &http.Server{
		Addr:           ":" + os.Getenv("APP_PORT"),
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server_err := server.ListenAndServe()
	helpers.ErrorPanic(server_err)
}
