package controller

import (
	"github.com/gofiber/fiber/v2"
	"www.github.com/biskitsx/go-api/webapp-sample/db"
	"www.github.com/biskitsx/go-api/webapp-sample/model"
)

type UserController interface {
	GetUser(c *fiber.Ctx) error
	GetUserById(c *fiber.Ctx) error
	DeleteUserById(c *fiber.Ctx) error
	UpdateUserById(c *fiber.Ctx) error
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

func (controller *userController) GetUserById(c *fiber.Ctx) error {
	user := &model.User{}
	userId, _ := c.ParamsInt("id")
	db.Db.Preload("Books").First(&user, "id = ?", userId)
	return c.JSON(user)
}

func (controller *userController) DeleteUserById(c *fiber.Ctx) error {
	user := &model.User{}
	userId, _ := c.ParamsInt("id")
	db.Db.Delete(&user, "id = ?", userId)
	return c.JSON(fiber.Map{
		"msg": "delete successfully",
	})
}

func (controller *userController) UpdateUserById(c *fiber.Ctx) error {
	user := &model.User{}
	userId, _ := c.ParamsInt("id")
	db.Db.First(&user, "id = ?", userId)
	return c.JSON(fiber.Map{
		"msg": "delete successfully",
	})
}
