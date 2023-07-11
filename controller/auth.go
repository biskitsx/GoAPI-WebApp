package controller

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"www.github.com/biskitsx/go-api/webapp-sample/db"
	"www.github.com/biskitsx/go-api/webapp-sample/model"
	"www.github.com/biskitsx/go-api/webapp-sample/model/dto"
	response "www.github.com/biskitsx/go-api/webapp-sample/utils"
)

type AuthController interface {
	Signup(c *fiber.Ctx) error
	Signin(c *fiber.Ctx) error
}

type authController struct {
}

func NewAuthController() AuthController {
	return &authController{}
}

func (controller *authController) Signup(c *fiber.Ctx) error {
	dto := dto.NewUserDto()
	if err := c.BodyParser(dto); err != nil {
		res := response.CreateError(400, err)
		return c.JSON(res)
	}

	// validate input
	if dto.Username == "" || dto.Password == "" {
		res := response.CreateError(400, "all field is required")
		return c.JSON(res)
	}

	// find if the username already exists
	user := model.User{}
	db.Db.First(&user, "username = ?", dto.Username)
	if user.Username != "" {
		res := response.CreateError(400, "this username already signed up")
		return c.JSON(res)
	}

	// hashing password
	hash, _ := bcrypt.GenerateFromPassword([]byte(dto.Password), 10)
	user.Username = dto.Username
	user.Password = string(hash)
	db.Db.Create(&user)

	return c.JSON(user)
}

func (controller *authController) Signin(c *fiber.Ctx) error {
	dto := dto.NewUserDto()
	if err := c.BodyParser(dto); err != nil {
		res := response.CreateError(400, err)
		return c.JSON(res)
	}

	// validate input
	if dto.Username == "" || dto.Password == "" {
		res := response.CreateError(400, "all field is required")
		return c.JSON(res)
	}

	// find if the username already exists
	user := model.User{}
	db.Db.First(&user, "username = ?", dto.Username)
	if user.Username == "" {
		res := response.CreateError(400, "wrong email")
		return c.JSON(res)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password)); err != nil {
		res := response.CreateError(400, "wrong password")
		return c.JSON(res)
	}

	return c.JSON(fiber.Map{
		"login": "successfully",
		"user":  user,
	})
}
