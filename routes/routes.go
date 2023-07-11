package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"www.github.com/biskitsx/go-api/webapp-sample/controller"
)

func Init(app *fiber.App) {

	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))

	authorRoutes(app)
	categoryRoutes(app)
	bookRoutes(app)
	authRoutes(app)
	userRoutes(app)
}

func authorRoutes(app *fiber.App) {
	author := controller.NewAuthorController()
	app.Get("/api/author", author.GetAuthors)
	app.Post("/api/author", author.CreateAuthor)
}

func categoryRoutes(app *fiber.App) {
	category := controller.NewCategoryController()
	app.Get("/api/category", category.GetCategory)
	app.Post("/api/category", category.CreateCategory)
}

func bookRoutes(app *fiber.App) {
	book := controller.NewBookController()
	app.Get("/api/book", book.GetBooks)
	app.Post("/api/book", book.CreateBook)
}

func authRoutes(app *fiber.App) {
	auth := controller.NewAuthController()
	app.Post("/api/auth/signup", auth.Signup)
	app.Post("/api/auth/signin", auth.Signin)

}

func userRoutes(app *fiber.App) {
	user := controller.NewUserController()
	app.Get("/api/user", user.GetUser)
	// app.Post("/api/book", book.CreateBook)
}
