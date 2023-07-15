package dto

type UserDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func NewUserDto() *UserDto {
	return &UserDto{}
}
