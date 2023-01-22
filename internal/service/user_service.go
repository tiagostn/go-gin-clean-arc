package service

import (
	"fmt"

	"github.com/tiagostn/go-gin-clean-arc/internal/model"
)

type UserService interface {
	Save(model.User) error
	Update(model.User) error
	Delete(model.User) error
	FindAll() []model.User
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (service *userService) Save(user model.User) error {
	fmt.Print(user)
	return nil
}

func (service *userService) Update(user model.User) error {
	fmt.Print(user)
	return nil
}

func (service *userService) Delete(user model.User) error {
	fmt.Print(user)
	return nil
}

func (service *userService) FindAll() []model.User {
	fmt.Print("findAll")
	return nil
}
