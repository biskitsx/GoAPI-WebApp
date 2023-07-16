package controller

import (
	"github.com/gofiber/fiber/v2"
	"www.github.com/biskitsx/go-api/webapp-sample/model/dto"
	"www.github.com/biskitsx/go-api/webapp-sample/service"
	"www.github.com/biskitsx/go-api/webapp-sample/utils"
)

type AuthController interface {
	Signup(c *fiber.Ctx) error
	Signin(c *fiber.Ctx) error
}

type authController struct {
	service     service.AuthService
	userService service.UserService
}

func NewAuthController() AuthController {
	return &authController{
		service:     service.NewAuthService(),
		userService: service.NewUserService(),
	}
}

func (controller *authController) Signup(c *fiber.Ctx) error {
	dto := dto.NewUserDto()
	if err := c.BodyParser(dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// validate input
	if err := dto.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// find if the username already exists
	_, err := controller.userService.FindByUsername(dto.Username)
	if err == nil { // user was found
		return fiber.NewError(fiber.StatusBadRequest, "this username already registerd")
	}

	// hashing password
	user := controller.service.Register(dto.Username, dto.Password)
	return c.JSON(user)
}

func (controller *authController) Signin(c *fiber.Ctx) error {
	dto := dto.NewUserDto()
	if err := c.BodyParser(dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// validate input
	if err := dto.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// find if the username already exists
	user, err := controller.userService.FindByUsername(dto.Username)
	if err != nil { // user wasn't found
		return fiber.NewError(fiber.StatusBadRequest, "this username already registerd")
	}

	if err := controller.service.Authenticate(user.Password, dto.Password); err != nil {
		return fiber.NewError(401, err.Error())
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
