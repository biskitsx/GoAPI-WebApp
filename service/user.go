package service

import (
	"errors"

	"www.github.com/biskitsx/go-api/webapp-sample/db"
	"www.github.com/biskitsx/go-api/webapp-sample/model"
)

type UserService interface {
	FindAll() *[]model.User
	FindById(id int) (*model.User, error)
	DeleteById(id int) (*model.User, error)
	FindByUsername(username string) (*model.User, error)
	// Create(dto *dto.UserDto) (*model.User, error)
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (service *userService) FindAll() *[]model.User {
	users := new([]model.User)
	db.Db.Find(users)
	return users
}

func (service *userService) FindById(id int) (*model.User, error) {
	user := new(model.User)
	db.Db.First(&user, "id = ?", id)
	if user.Username == "" {
		return user, errors.New("user not found")
	}
	return user, nil
}

func (service *userService) FindByUsername(username string) (*model.User, error) {
	user := new(model.User)
	db.Db.First(&user, "username = ?", username)
	if user.Username == "" {
		return user, errors.New("user not found")
	}
	return user, nil
}

func (service *userService) DeleteById(id int) (*model.User, error) {
	user := new(model.User)
	db.Db.Delete(&user, "id = ?", id)
	if user.Username == "" {
		return user, errors.New("user not found")
	}
	return user, nil
}

// func (service *userService) Create(dto *dto.UserDto) (*model.User, error) {
// 	err := dto.Validate()
// 	if err != nil {
// 		return nil, errors.New(err.Error())
// 	}
// 	user := model.NewUser(dto.Username, dto.Password)
// 	db.Db.Create(user)
// 	return user, nil
// }
