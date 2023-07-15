package dto

import (
	"github.com/go-playground/validator/v10"
)

type BookDto struct {
	Title      string `json:"title" validate:"required"`
	CategoryID uint   `json:"categoryId" validate:"required,number"`
	AuthorID   uint   `json:"authorId" validate:"required,number"`
	Price      uint   `json:"price" validate:"required,number"`
}

func NewBookDto() *BookDto {
	return &BookDto{}
}
func (dto *BookDto) Validate() error {
	err := validator.New().Struct(dto)
	if err != nil {
		return err
	}
	return nil
}
