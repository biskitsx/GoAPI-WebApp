package dto

type CategoryDto struct {
	Name string `json:"name"`
}

func NewCategoryDto() *CategoryDto {
	return &CategoryDto{}
}
