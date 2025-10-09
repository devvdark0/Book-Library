package service

import (
	"github.com/devvdark0/book-library/internal/model"
	"github.com/google/uuid"
)

type Service interface {
	CreateBook(req model.CreateBookRequest) (model.Book, error)
	ListBooks() ([]model.Book, error)
	GetBook(id uuid.UUID) (model.Book, error)
	UpdateBook(id uuid.UUID, req model.UpdateBookRequest) error
	DeleteBook(id uuid.UUID) error
}
