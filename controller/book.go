package controller

import (
	"github.com/gofiber/fiber/v2"
	"www.github.com/biskitsx/go-api/webapp-sample/db"
	"www.github.com/biskitsx/go-api/webapp-sample/model"
	"www.github.com/biskitsx/go-api/webapp-sample/model/dto"
	response "www.github.com/biskitsx/go-api/webapp-sample/utils"
)

type BookController interface {
	CreateBook(c *fiber.Ctx) error
	GetBooks(c *fiber.Ctx) error
	AddBookToUser(c *fiber.Ctx) error
}

type bookController struct{}

func NewBookController() BookController {
	return &bookController{}
}

func (controller *bookController) CreateBook(c *fiber.Ctx) error {
	dto := dto.NewBookDto()
	if err := c.BodyParser(dto); err != nil {
		res := response.CreateError(400, err)
		return c.JSON(res)
	}
	book := model.NewBook(dto.Title, dto.CategoryID, dto.AuthorID, dto.Price)
	db.Db.Create(&book)
	return c.JSON(book)
}

func (controller *bookController) GetBooks(c *fiber.Ctx) error {
	books := new([]model.Book)
	// mode
	db.Db.Preload("Category").Preload("Author").Find(books)
	return c.JSON(books)
}

func (controller *bookController) AddBookToUser(c *fiber.Ctx) error {
	// find book
	bookId, err := c.ParamsInt("id")
	book := model.Book{}
	db.Db.First(&book, "id = ?", bookId)
	if err != nil {
		return fiber.NewError(400, "invalid param")
	}

	// find user
	userId := c.Locals("userId")
	user := model.User{}
	db.Db.First(&user, "id = ?", userId)

	// add book
	user.Books = append(user.Books, book)
	db.Db.Preload("Books").Save(&user)

	return c.JSON(user)
}
