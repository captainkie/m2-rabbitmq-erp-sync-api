package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	_ "time/tzdata"

	"github.com/captainkie/websync-api/config"
	_ "github.com/captainkie/websync-api/docs"
	controller "github.com/captainkie/websync-api/internal/app/controllers"
	model "github.com/captainkie/websync-api/internal/app/models"
	"github.com/captainkie/websync-api/internal/app/repository"
	"github.com/captainkie/websync-api/internal/app/routes"
	"github.com/captainkie/websync-api/internal/app/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
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
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("❌ Failed to load environment variables: %v", err)
	}

	os.Setenv("TZ", "Asia/Bangkok")

	// Initialize critical connections
	if err := initializeConnections(); err != nil {
		log.Fatalf("❌ Failed to initialize critical connections: %v", err)
	}

	go producer()
	go imageScheduler()
	go syncProductScheduler()

	initialize()
}

// initializeConnections checks and establishes critical connections
func initializeConnections() error {
	// Check database connection
	db := config.ConnectDatabase()
	if db == nil {
		return fmt.Errorf("❌ Failed to connect to database")
	}

	// Check RabbitMQ connection
	rabbitMQURL := os.Getenv("RABBITMQ_URL")
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		return fmt.Errorf("❌ Failed to connect to RabbitMQ: %v", err)
	}
	conn.Close()

	return nil
}

func initialize() {
	// Connect database
	db := config.ConnectDatabase()
	if db == nil {
		log.Fatal("❌ Failed to connect to database")
	}

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

	db.Table("connection_logs").AutoMigrate(&model.ConnectionLogs{})
	db.Table("add_logs").AutoMigrate(&model.AddLogs{})
	db.Table("update_logs").AutoMigrate(&model.UpdateLogs{})
	db.Table("stock_logs").AutoMigrate(&model.StockLogs{})
	db.Table("store_logs").AutoMigrate(&model.StoreLogs{})
	db.Table("postflag_logs").AutoMigrate(&model.PostflagLogs{})
	db.Table("image_logs").AutoMigrate(&model.ImageLogs{})
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

	// Create server
	server := &http.Server{
		Addr:           ":" + os.Getenv("APP_PORT"),
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Start server in a goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutting down server...")

	// Create shutdown context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Close database connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Error getting database connection: %v", err)
	} else {
		if err := sqlDB.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	fmt.Println("Server exiting")
}
