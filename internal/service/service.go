package service

import "github.com/devvdark0/book-library/internal/model"

type BookService interface {
	CreateBook(req model.CreateBookRequest) (model.Book, error)
	GetBook() (model.Book, error)
	ListBooks() ([]model.Book, error)
	UpdateBook() (model.Book, error)
	DeleteBook() error
}
