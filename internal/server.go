package internal

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tiagostn/go-gin-clean-arc/docs"
	"github.com/tiagostn/go-gin-clean-arc/internal/controller"
)

func Server() *gin.Engine {
	engine := gin.Default()

	docs.SwaggerInfo.Description = "Example Golang REST API - overrided"

	engine.GET("/", controller.RootController)
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return engine
}
