package controller

import (
	"github.com/gofiber/fiber/v2"
	"www.github.com/biskitsx/go-api/webapp-sample/db"
	"www.github.com/biskitsx/go-api/webapp-sample/model"
	"www.github.com/biskitsx/go-api/webapp-sample/model/dto"
	response "www.github.com/biskitsx/go-api/webapp-sample/utils"
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
		res := response.CreateError(400, err)
		return c.JSON(res)
	}
	author := model.NewAuthor(dto.Name)
	db.Db.Create(&author)
	return c.JSON(dto)
}

func (controller *authorController) GetAuthors(c *fiber.Ctx) error {
	authors := new([]model.Author)
	db.Db.Find(authors)
	return c.JSON(authors)
}
