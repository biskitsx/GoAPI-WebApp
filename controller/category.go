package controller

import (
	"github.com/gofiber/fiber/v2"
	"www.github.com/biskitsx/go-api/webapp-sample/model/dto"
	"www.github.com/biskitsx/go-api/webapp-sample/service"
)

type CategoryController interface {
	CreateCategory(c *fiber.Ctx) error
	GetCategory(c *fiber.Ctx) error
	GetCategoryById(c *fiber.Ctx) error
}

type categoryController struct {
	service service.CategoryService
}

func NewCategoryController() CategoryController {
	return &categoryController{service: service.NewCategoryService()}
}

func (controller *categoryController) CreateCategory(c *fiber.Ctx) error {
	dto := dto.NewCategoryDto()
	if err := c.BodyParser(dto); err != nil {
		return fiber.NewError(400, err.Error())
	}
	category, err := controller.service.Create(dto)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}
	return c.JSON(category)
}

func (controller *categoryController) GetCategory(c *fiber.Ctx) error {
	categories := controller.service.FindAll()
	return c.JSON(categories)
}

func (controller *categoryController) GetCategoryById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(400, "Bad parameter")
	}
	categories, err := controller.service.FindById(id)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}
	return c.JSON(categories)
}
