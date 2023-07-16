package dto

import "github.com/go-playground/validator/v10"

type UserDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func NewUserDto() *UserDto {
	return &UserDto{}
}

func (dto *UserDto) Validate() error {
	err := validator.New().Struct(dto)
	if err != nil {
		return err
	}
	return nil
}
