package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"www.github.com/biskitsx/go-api/webapp-sample/db"
	"www.github.com/biskitsx/go-api/webapp-sample/model"
)

type AuthService interface {
	Register(username string, password string) *model.User
	Authenticate(hash string, password string) error
}

type authService struct {
}

func NewAuthService() AuthService {
	return &authService{}
}

func (service *authService) Register(username string, password string) *model.User {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	user := model.NewUser(username, string(hash))
	db.Db.Create(user)
	return user
}
func (service *authService) Authenticate(hash string, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return errors.New("wrong password")
	}
	return nil
}
