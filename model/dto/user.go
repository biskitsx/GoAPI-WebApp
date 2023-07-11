package dto

type UserDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUserDto() *UserDto {
	return &UserDto{}
}
