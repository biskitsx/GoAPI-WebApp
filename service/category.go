package service

import (
	"errors"

	"www.github.com/biskitsx/go-api/webapp-sample/db"
	"www.github.com/biskitsx/go-api/webapp-sample/model"
	"www.github.com/biskitsx/go-api/webapp-sample/model/dto"
)

type CategoryService interface {
	FindAll() *[]model.Category
	FindById(id int) (*model.Category, error)
	Create(dto *dto.CategoryDto) (*model.Category, error)
}
type categoryService struct {
}

func NewCategoryService() CategoryService {
	return &categoryService{}
}

func (service *categoryService) FindAll() *[]model.Category {
	category := new([]model.Category)
	db.Db.Find(&category)
	return category
}

func (service *categoryService) FindById(id int) (*model.Category, error) {
	category := new(model.Category)
	db.Db.First(&category, "id = ?", id)
	if category.Name == "" {
		return category, errors.New("category not founded")
	}
	return category, nil
}

func (service *categoryService) Create(dto *dto.CategoryDto) (*model.Category, error) {
	err := dto.Validate()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	category := model.NewCategory(dto.Name)
	db.Db.Create(&category)
	return category, nil
}
