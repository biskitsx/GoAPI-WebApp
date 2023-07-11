package model

type Author struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

func NewAuthor(name string) *Author {
	return &Author{Name: name}
}
