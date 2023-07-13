package controller

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"www.github.com/biskitsx/go-api/webapp-sample/db"
	"www.github.com/biskitsx/go-api/webapp-sample/model"
	"www.github.com/biskitsx/go-api/webapp-sample/model/dto"
	"www.github.com/biskitsx/go-api/webapp-sample/utils"
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
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// validate input
	if dto.Username == "" || dto.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "All field is required")
	}

	// find if the username already exists
	user := model.User{}
	db.Db.First(&user, "username = ?", dto.Username)
	if user.Username != "" {
		return fiber.NewError(fiber.StatusBadRequest, "this username already signed up")
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
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// validate input
	if dto.Username == "" || dto.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "All field is required")
	}

	// find if the username already exists
	user := model.User{}
	db.Db.First(&user, "username = ?", dto.Username)
	if user.Username == "" {

		return fiber.NewError(fiber.StatusBadRequest, "Wrong username")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password)); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Wrong password")
	}

	// token & cookie
	tokenManager := utils.NewTokenManager()
	token, _ := tokenManager.CreateToken(user.ID)
	cookieToken := &fiber.Cookie{
		Name:  "access_token",
		Value: token,
	}
	c.Cookie(cookieToken)

	return c.JSON(fiber.Map{
		"login": "successfully",
		"user":  user,
	})
}
