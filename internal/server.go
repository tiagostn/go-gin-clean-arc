package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/tiagostn/go-gin-clean-arc/internal/controller"
)

func Server() *gin.Engine {
	engine := gin.Default()
	engine.GET("/", controller.RootController)
	return engine
}
