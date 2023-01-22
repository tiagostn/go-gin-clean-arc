package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiagostn/go-gin-clean-arc/internal/model"
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

// DefaultUserController godoc
// @Summary      Default User
// @Description  get default user
// @Produce      json
// @Success      200  {object} model.User
// @Router       /user [get]
func DefaultUserController(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"user": model.User{Name: "Tiago", Email: "tiagostn@gmail.com"}})
}
