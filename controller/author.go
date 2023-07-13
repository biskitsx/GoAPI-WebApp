package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"www.github.com/biskitsx/go-api/webapp-sample/db"
	"www.github.com/biskitsx/go-api/webapp-sample/model"
	"www.github.com/biskitsx/go-api/webapp-sample/model/dto"
)

type AuthorController interface {
	CreateAuthor(c *fiber.Ctx) error
	GetAuthors(c *fiber.Ctx) error
}

type authorController struct{}

func NewAuthorController() AuthorController {
	return &authorController{}
}

func (controller *authorController) CreateAuthor(c *fiber.Ctx) error {
	dto := dto.NewAuthorDto()
	if err := c.BodyParser(dto); err != nil {
		return fiber.NewError(400, err.Error())
	}
	author := model.NewAuthor(dto.Name)
	db.Db.Create(&author)
	return c.JSON(dto)
}

func (controller *authorController) GetAuthors(c *fiber.Ctx) error {
	authors := new([]model.Author)
	db.Db.Find(authors)
	fmt.Println(c.Locals("userId"))
	return c.JSON(authors)
}
