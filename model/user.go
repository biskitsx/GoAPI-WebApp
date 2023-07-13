package model

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	// Books    []Book `gorm:"foreignKey:BookID"json:"books"`
	Books []Book `gorm:"many2many:user_books" json:"books"`
	// BooksID  []uint  `json:"booksId"`
}

func NewUser(username string, password string) *User {
	return &User{Username: username, Password: password}
}
