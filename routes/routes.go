package routes

import (
	"github.com/gofiber/fiber/v2"
	"www.github.com/biskitsx/go-api/webapp-sample/controller"
)

func Init(app *fiber.App) {
	authorRoutes(app)
}

func authorRoutes(app *fiber.App) {
	author := controller.NewAuthorController()
	app.Get("/api/author", author.GetAuthors)
	app.Post("/api/author", author.CreateAuthor)
}
