package model

type Book struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Title      string    `json:"title"`
	CategoryID uint      `json:"categoryId"`
	Category   *Category `json:"category"`
	AuthorID   uint      `json:"authorId"`
	Author     *Author   `json:"author"`
}
