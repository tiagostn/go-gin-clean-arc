package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tiagostn/go-gin-clean-arc/controller"
)

func main() {
	fmt.Println("Golang Gin Server 👌")
	router := gin.Default()

	router.GET("/", controller.RootController)

	router.Run()
}
