package model

type Category struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

func NewCategory(name string) *Category {
	return &Category{Name: name}
}
