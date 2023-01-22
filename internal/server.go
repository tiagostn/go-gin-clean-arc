package internal

import (
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tiagostn/go-gin-clean-arc/docs"
	"github.com/tiagostn/go-gin-clean-arc/internal/api"
	"github.com/tiagostn/go-gin-clean-arc/internal/controller"
	"github.com/tiagostn/go-gin-clean-arc/internal/middlewares"
	"github.com/tiagostn/go-gin-clean-arc/internal/service"
)

var (
	userService service.UserService = service.NewUserService()

	userController controller.UserController = controller.NewUserController(userService)
)

func Server() *gin.Engine {
	server := gin.Default()

	server.Use(helmet.Default())

	server.Use(middlewares.Logger())
	server.Use(middlewares.Analytics())

	docs.SwaggerInfo.Description = "Example Golang REST API - overrided"

	api.GetUserApi(userController).Register(server)

	server.GET("/", controller.RootController)
	server.GET("/user", controller.DefaultUserController)
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return server
}
