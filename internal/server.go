package internal

import (
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tiagostn/go-gin-clean-arc/docs"
	"github.com/tiagostn/go-gin-clean-arc/internal/controller"
	"github.com/tiagostn/go-gin-clean-arc/internal/middlewares"
)

func Server() *gin.Engine {
	server := gin.Default()

	server.Use(helmet.Default())

	server.Use(middlewares.Logger())
	server.Use(middlewares.Analytics())

	docs.SwaggerInfo.Description = "Example Golang REST API - overrided"

	server.GET("/", controller.RootController)
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return server
}
