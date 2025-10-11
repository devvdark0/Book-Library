package book

import (
	"github.com/devvdark0/book-library/internal/model"
	"github.com/devvdark0/book-library/internal/repository"
	"github.com/devvdark0/book-library/internal/service"
	"github.com/google/uuid"
)

type bookService struct {
	repo repository.Repository
}

func NewBookService(repo repository.Repository) service.Service {
	return bookService{repo: repo}
}

func (b bookService) CreateBook(req model.CreateBookRequest) error {
	//TODO implement me
	panic("implement me")
}

func (b bookService) ListBooks() ([]model.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b bookService) GetBook(id uuid.UUID) (model.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b bookService) UpdateBook(id uuid.UUID, req model.UpdateBookRequest) error {
	//TODO implement me
	panic("implement me")
}

func (b bookService) DeleteBook(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
