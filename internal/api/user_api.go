package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tiagostn/go-gin-clean-arc/internal/controller"
)

type UserApi struct {
	path       string
	controller controller.UserController
}

func GetUserApi(userController controller.UserController) *UserApi {
	return &UserApi{
		path:       "/users",
		controller: userController,
	}
}

func (api *UserApi) Register(server *gin.Engine) {
	users := server.Group(api.path)
	{
		users.GET("", api.getAll)
		// users.POST("", api.Create)
		// users.PUT(":id", api.Update)
		// users.DELETE(":id", api.Delete)
	}
}

// getAll godoc
// @Summary List existing users
// @Description Get all the existing users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} model.User
// @Router /users [get]
func (api *UserApi) getAll(ctx *gin.Context) {
	ctx.JSON(200, api.controller.FindAll(ctx))
}
