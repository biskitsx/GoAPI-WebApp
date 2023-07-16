package service

import (
	"errors"

	"www.github.com/biskitsx/go-api/webapp-sample/db"
	"www.github.com/biskitsx/go-api/webapp-sample/model"
	"www.github.com/biskitsx/go-api/webapp-sample/model/dto"
)

type BookService interface {
	FindAll() *[]model.Book
	FindById(id int) (*model.Book, error)
	Create(dto *dto.BookDto) (*model.Book, error)
	AddBook(user *model.User, book *model.Book)
	RemoveBook(user *model.User, book *model.Book)
}

type bookService struct {
}

func NewBookService() BookService {
	return &bookService{}
}

func (service *bookService) FindAll() *[]model.Book {
	books := new([]model.Book)
	db.Db.Preload("Category").Preload("Author").Find(&books)
	return books
}

func (service *bookService) FindById(id int) (*model.Book, error) {
	book := new(model.Book)
	db.Db.First(&book, "id = ?", id)
	if book.Title == "" {
		return book, errors.New("book not founded")
	}
	return book, nil
}

func (service *bookService) Create(dto *dto.BookDto) (*model.Book, error) {
	err := dto.Validate()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	book := model.NewBook(dto.Title, dto.CategoryID, dto.AuthorID, dto.Price)
	db.Db.Create(book)
	return book, nil
}

func (service *bookService) AddBook(user *model.User, book *model.Book) {
	db.Db.First(user).Association("Books").Append(book)
}

func (service *bookService) RemoveBook(user *model.User, book *model.Book) {
	db.Db.Model(user).Association("Books").Delete(book)
}
