package main

import (
	"github.com/gofiber/fiber/v2"
	"www.github.com/biskitsx/go-api/webapp-sample/db"
	"www.github.com/biskitsx/go-api/webapp-sample/routes"
)

func main() {
	app := fiber.New()
	db.ConnectDb()

	routes.Init(app)
	app.Listen(":3000")
}
