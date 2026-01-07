package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/captainkie/websync-api/config"
	"github.com/captainkie/websync-api/types/response"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

type HealthController struct{}

func NewHealthController() *HealthController {
	return &HealthController{}
}

// HealthCheck godoc
// @Summary Health check endpoint
// @Description Check if the API is running
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func (h *HealthController) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Service is healthy",
	})
}

// ReadinessCheck godoc
// @Summary Readiness check endpoint
// @Description Check if the API is ready to handle requests
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 503 {object} map[string]string
// @Router /health/ready [get]
func (h *HealthController) ReadinessCheck(c *gin.Context) {
	// TODO: Add checks for database, RabbitMQ, and other critical services
	// For now, we'll just return 200 OK
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Service is ready",
	})
}

// LivenessCheck godoc
// @Summary Liveness check endpoint
// @Description Check if the API is alive
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 503 {object} map[string]string
// @Router /health/live [get]
func (h *HealthController) LivenessCheck(c *gin.Context) {
	// TODO: Add checks for memory usage, CPU usage, and other system metrics
	// For now, we'll just return 200 OK
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Service is alive",
	})
}

// HealthCheckLiveness godoc
// @Summary Check application liveness
// @Description Get the liveness status of the application
// @Produce json
// @Tags Health
// @Success 200 {object} response.Response{}
// @Router /health/live [get]
func (controller *HealthController) HealthCheckLiveness(ctx *gin.Context) {
	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "healthy",
		Message: "Application is alive",
		Data: map[string]interface{}{
			"timestamp": time.Now().Format(time.RFC3339),
		},
	}

	ctx.JSON(http.StatusOK, webResponse)
}

// HealthCheckReadiness godoc
// @Summary Check application readiness
// @Description Get the readiness status of the application
// @Produce json
// @Tags Health
// @Success 200 {object} response.Response{}
// @Failure 503 {object} response.Response{}
// @Router /health/ready [get]
func (controller *HealthController) HealthCheckReadiness(ctx *gin.Context) {
	// Check database connection
	db := config.ConnectDatabase()
	if db == nil {
		webResponse := response.Response{
			Code:    http.StatusServiceUnavailable,
			Status:  "unhealthy",
			Message: "Database connection failed",
			Data: map[string]interface{}{
				"timestamp": time.Now().Format(time.RFC3339),
			},
		}
		ctx.JSON(http.StatusServiceUnavailable, webResponse)
		return
	}

	// Check RabbitMQ connection
	rabbitMQURL := os.Getenv("RABBITMQ_URL")
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusServiceUnavailable,
			Status:  "unhealthy",
			Message: "RabbitMQ connection failed",
			Data: map[string]interface{}{
				"timestamp": time.Now().Format(time.RFC3339),
			},
		}
		ctx.JSON(http.StatusServiceUnavailable, webResponse)
		return
	}
	conn.Close()

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "healthy",
		Message: "Application is ready",
		Data: map[string]interface{}{
			"timestamp": time.Now().Format(time.RFC3339),
		},
	}

	ctx.JSON(http.StatusOK, webResponse)
}
