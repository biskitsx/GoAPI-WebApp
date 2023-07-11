package dto

type BookDto struct {
	Title      string `json:"title"`
	CategoryID uint   `json:"categoryId"`
	AuthorID   uint   `json:"authorId"`
	Price      uint   `json:"price"`
}

func NewBookDto() *BookDto {
	return &BookDto{}
}
