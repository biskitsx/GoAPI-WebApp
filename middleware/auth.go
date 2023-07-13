package middleware

import (
	"github.com/gofiber/fiber/v2"
	"www.github.com/biskitsx/go-api/webapp-sample/utils"
)

func VerifyUser(c *fiber.Ctx) error {
	token := c.Cookies("access_token")
	if token == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "no access_token")
	}
	tokenManager := utils.NewTokenManager()
	payload, err := tokenManager.VerifyToken(token)

	if err != nil {
		return fiber.NewError(fiber.StatusNonAuthoritativeInformation, "user not authenticated")
	}

	id, ok := payload["id"].(float64)
	if !ok {
		return fiber.NewError(fiber.StatusConflict, "float64 error")
	}
	c.Locals("userId", id)
	return c.Next()
}
