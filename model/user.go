package model

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Books    []Book `gorm:"many2many:user_books" json:"books"`
}

func NewUser(username string, password string) *User {
	return &User{Username: username, Password: password}
}
