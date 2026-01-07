package routes

import (
	"net/http"

	controller "github.com/captainkie/websync-api/internal/app/controllers"
	"github.com/captainkie/websync-api/internal/app/middleware"
	"github.com/captainkie/websync-api/internal/app/repository"
	"github.com/captainkie/websync-api/pkg/errors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func MainRouter(
	userRepository repository.UsersRepository, authenticationController *controller.AuthenticationController, usersController *controller.UserController, queueController *controller.QueueController) *gin.Engine {
	router := gin.Default()

	// Apply global middlewares
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.IPRateLimiter())
	router.Use(errors.ErrorHandler())

	// swagger docs
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health check endpoints
	healthController := controller.NewHealthController()
	router.GET("/health", healthController.HealthCheck)
	router.GET("/health/live", healthController.HealthCheckLiveness)
	router.GET("/health/ready", healthController.HealthCheckReadiness)

	// home route
	router.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "api running ...")
	})

	// 404 route
	router.NoRoute(func(c *gin.Context) {
		err := errors.NewNotFoundError("Page not found")
		c.Error(err)
		c.Abort()
	})

	// group routes
	baseRouter := router.Group("/api")

	// authentication routes
	authenticationRouter := baseRouter.Group("/authentication")
	authenticationRouter.POST("/login", authenticationController.Login)
	authenticationRouter.POST("/register", authenticationController.Register)

	// users routes
	usersRouter := baseRouter.Group("/users")
	usersRouter.Use(middleware.AuthMiddleware(userRepository))
	usersRouter.GET("", usersController.FindAll)
	usersRouter.GET("/:id", usersController.FindById)
	usersRouter.POST("", usersController.Create)
	usersRouter.PATCH("/:id", usersController.Update)
	usersRouter.DELETE("/:id", usersController.Delete)

	// queue routes
	queueRouter := baseRouter.Group("/queue")
	queueRouter.Use(middleware.AuthMiddleware(userRepository))
	queueRouter.GET("/products", queueController.ProductsSync)
	queueRouter.GET("/images", queueController.ImagesSync)
	queueRouter.POST("/daily-sales", queueController.DailySales)

	return router
}
