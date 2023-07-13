package main

import (
	"github.com/gofiber/fiber/v2"
	"www.github.com/biskitsx/go-api/webapp-sample/config"
	"www.github.com/biskitsx/go-api/webapp-sample/db"
	"www.github.com/biskitsx/go-api/webapp-sample/routes"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: config.CustomErrorHandler,
	})

	db.ConnectDb()

	routes.Init(app)
	app.Listen(":3000")
}
