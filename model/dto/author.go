package dto

type AuthorDto struct {
	Name string `json:"name"`
}

func NewAuthorDto() *AuthorDto {
	return &AuthorDto{}
}
