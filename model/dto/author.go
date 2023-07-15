package dto

type AuthorDto struct {
	Name string `json:"name" validate:"required"`
}

func NewAuthorDto() *AuthorDto {
	return &AuthorDto{}
}
