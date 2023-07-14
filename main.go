package main

import (
	"github.com/gofiber/fiber/v2"
	"www.github.com/biskitsx/go-api/webapp-sample/config"
	"www.github.com/biskitsx/go-api/webapp-sample/db"
	_ "www.github.com/biskitsx/go-api/webapp-sample/docs"
	_ "www.github.com/biskitsx/go-api/webapp-sample/model"
	_ "www.github.com/biskitsx/go-api/webapp-sample/model/dto"
	"www.github.com/biskitsx/go-api/webapp-sample/routes"
)

// @title Swagger Example API
// @version 1.0
// @description A API for fiber framework
// @host localhost:3000/api

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: config.CustomErrorHandler,
	})

	db.ConnectDb()

	routes.Init(app)
	app.Listen(":3000")
}
