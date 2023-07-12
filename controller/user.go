package controller

import (
	"github.com/gofiber/fiber/v2"
	"www.github.com/biskitsx/go-api/webapp-sample/db"
	"www.github.com/biskitsx/go-api/webapp-sample/model"
)

type UserController interface {
	GetUser(c *fiber.Ctx) error
}
type userController struct {
}

func NewUserController() UserController {
	return &userController{}
}

func (controller *userController) GetUser(c *fiber.Ctx) error {
	users := &[]model.User{}
	db.Db.Preload("Books").Find(&users)
	return c.JSON(users)
}
