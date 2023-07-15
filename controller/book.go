package controller

import (
	"github.com/gofiber/fiber/v2"
	"www.github.com/biskitsx/go-api/webapp-sample/db"
	"www.github.com/biskitsx/go-api/webapp-sample/model"
	"www.github.com/biskitsx/go-api/webapp-sample/model/dto"
	"www.github.com/biskitsx/go-api/webapp-sample/service"
)

type BookController interface {
	CreateBook(c *fiber.Ctx) error
	GetBooks(c *fiber.Ctx) error
	AddBookToUser(c *fiber.Ctx) error
	RemoveBookFromUser(c *fiber.Ctx) error
}

type bookController struct {
	service service.BookService
}

func NewBookController() BookController {
	return &bookController{
		service: service.NewBookService(),
	}
}

func (controller *bookController) CreateBook(c *fiber.Ctx) error {
	dto := dto.NewBookDto()
	if err := c.BodyParser(dto); err != nil {
		return fiber.NewError(405, "error")
	}
	book, err := controller.service.Create(dto)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}
	return c.JSON(book)
}

func (controller *bookController) GetBooks(c *fiber.Ctx) error {
	books := controller.service.FindAll()
	return c.JSON(books)
}

func (controller *bookController) AddBookToUser(c *fiber.Ctx) error {
	// find book
	bookId, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(400, "invalid param")
	}
	book, err := controller.service.FindById(bookId)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	// find user
	userId := c.Locals("userId")
	user := new(model.User)
	db.Db.First(user, "id = ?", userId)

	// add book
	controller.service.AddBook(user, book)
	return c.JSON(user)
}

func (controller *bookController) RemoveBookFromUser(c *fiber.Ctx) error {
	// find book
	bookId, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(400, "invalid param")
	}
	book, err := controller.service.FindById(bookId)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	// find user
	userId := c.Locals("userId")
	user := new(model.User)
	db.Db.First(user, "id = ?", userId)

	// add book
	controller.service.RemoveBook(user, book)
	return c.JSON(user)
}
