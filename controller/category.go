package controller

import (
	"github.com/gofiber/fiber/v2"
	"www.github.com/biskitsx/go-api/webapp-sample/db"
	"www.github.com/biskitsx/go-api/webapp-sample/model"
	"www.github.com/biskitsx/go-api/webapp-sample/model/dto"
	"www.github.com/biskitsx/go-api/webapp-sample/utils"
)

type CategoryController interface {
	CreateCategory(c *fiber.Ctx) error
	GetCategory(c *fiber.Ctx) error
}

type categoryController struct {
}

func NewCategoryController() CategoryController {
	return &categoryController{}
}

func (controller *categoryController) CreateCategory(c *fiber.Ctx) error {
	dto := dto.NewCategoryDto()
	if err := c.BodyParser(dto); err != nil {
		res := utils.CreateError(400, err)
		return c.JSON(res)
	}
	category := model.NewCategory(dto.Name)
	db.Db.Create(&category)
	return c.JSON(category)
}

func (controller *categoryController) GetCategory(c *fiber.Ctx) error {
	categories := new([]model.Category)
	db.Db.Find(&categories)
	return c.JSON(categories)
}
