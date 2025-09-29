package service

import "github.com/devvdark0/book-library/internal/model"

type BookService interface {
	CreateBook(req model.CreateBookRequest) (model.Book, error)
	GetBook(id string) (model.Book, error)
	ListBooks() ([]model.Book, error)
	UpdateBook(id string, request model.UpdateBookRequest) (model.Book, error)
	DeleteBook() error
}
