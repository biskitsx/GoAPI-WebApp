package dto

import (
	"github.com/go-playground/validator/v10"
)

type CategoryDto struct {
	Name string `json:"name" validate:"required"`
}

func NewCategoryDto() *CategoryDto {
	return &CategoryDto{}
}

func (c *CategoryDto) Validate() error {
	err := validator.New().Struct(c)
	if err != nil {
		return err
	}
	return nil
	// if err == nil {
	// 	return nil
	// }
	// errors := err.(validator.ValidationErrors)
	// if len(errors) == 0 {
	// 	return nil
	// }

	// return err
}
