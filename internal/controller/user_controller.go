package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tiagostn/go-gin-clean-arc/internal/model"
	"github.com/tiagostn/go-gin-clean-arc/internal/service"
)

type UserController interface {
	FindAll(ctx *gin.Context) []model.User
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (controller *userController) FindAll(ctx *gin.Context) []model.User {
	return nil
}
