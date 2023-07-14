package model

type Book struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Title      string    `json:"title"`
	CategoryId uint      `json:"categoryId"`
	Category   *Category `json:"category"`
	AuthorID   uint      `json:"authorId"`
	Author     *Author   `json:"author"`
	Price      uint      `json:"price"`
}

func NewBook(title string, categoryId uint, authorId uint, price uint) *Book {
	return &Book{
		Title:      title,
		CategoryId: categoryId,
		AuthorID:   authorId,
		Price:      price,
	}
}
