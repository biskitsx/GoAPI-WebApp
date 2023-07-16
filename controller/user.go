package controller

import (
	"github.com/gofiber/fiber/v2"
	"www.github.com/biskitsx/go-api/webapp-sample/db"
	"www.github.com/biskitsx/go-api/webapp-sample/model"
	"www.github.com/biskitsx/go-api/webapp-sample/service"
)

type UserController interface {
	GetUser(c *fiber.Ctx) error
	GetUserById(c *fiber.Ctx) error
	DeleteUserById(c *fiber.Ctx) error
	UpdateUserById(c *fiber.Ctx) error
}
type userController struct {
	service service.UserService
}

func NewUserController() UserController {
	return &userController{service: service.NewUserService()}
}

func (controller *userController) GetUser(c *fiber.Ctx) error {
	users := controller.service.FindAll()
	return c.JSON(users)
}

func (controller *userController) GetUserById(c *fiber.Ctx) error {
	userId, _ := c.ParamsInt("id")
	user, err := controller.service.FindById(userId)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}
	return c.JSON(user)
}

func (controller *userController) DeleteUserById(c *fiber.Ctx) error {
	userId, _ := c.ParamsInt("id")
	_, err := controller.service.DeleteById(userId)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}
	return c.JSON(fiber.Map{
		"msg": "delete successfully",
	})
}

func (controller *userController) UpdateUserById(c *fiber.Ctx) error {
	user := &model.User{}
	userId, _ := c.ParamsInt("id")
	db.Db.First(&user, "id = ?", userId)
	return c.JSON(fiber.Map{
		"msg": "update successfully",
	})
}
