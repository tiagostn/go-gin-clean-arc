package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RootController godoc
// @Summary      Root
// @Description  get hello world response
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Router       / [get]
func RootController(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "hello world 👌 from docker!"})
}
