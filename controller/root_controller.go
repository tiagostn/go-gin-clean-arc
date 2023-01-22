package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootController(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "hello world 👌 from controller!"})
}
