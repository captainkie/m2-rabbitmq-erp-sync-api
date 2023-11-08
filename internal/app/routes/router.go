package routes

import (
	"net/http"

	controller "github.com/captainkie/websync-api/internal/app/controllers"
	"github.com/captainkie/websync-api/internal/app/middleware"
	"github.com/captainkie/websync-api/internal/app/repository"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func MainRouter(
	userRepository repository.UsersRepository, authenticationController *controller.AuthenticationController, usersController *controller.UserController, queueController *controller.QueueController) *gin.Engine {
	router := gin.Default()
	// swagger docs
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// home route
	router.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "api running ...")
	})

	// 404 route
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	// group routes
	baseRouter := router.Group("/api")

	// authentication routes
	authenticationRouter := baseRouter.Group("/authentication")
	authenticationRouter.POST("/login", authenticationController.Login)
	authenticationRouter.POST("/register", authenticationController.Register)

	// users routes
	usersRouter := baseRouter.Group("/users")
	usersRouter.GET("", middleware.AuthMiddleware(userRepository), usersController.FindAll)
	usersRouter.GET("/:id", middleware.AuthMiddleware(userRepository), usersController.FindById)
	usersRouter.POST("", middleware.AuthMiddleware(userRepository), usersController.Create)
	usersRouter.PATCH("/:id", middleware.AuthMiddleware(userRepository), usersController.Update)
	usersRouter.DELETE("/:id", middleware.AuthMiddleware(userRepository), usersController.Delete)

	// queue routes
	queueRouter := baseRouter.Group("/queue")
	// queueRouter.GET("/products", middleware.AuthMiddleware(userRepository), queueController.ProductsSync)
	queueRouter.GET("/products", queueController.ProductsSync)
	queueRouter.GET("/images", queueController.ImagesSync)
	usersRouter.POST("/daily-sales", middleware.AuthMiddleware(userRepository), queueController.DailySales)

	return router
}
