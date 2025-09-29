package book

import (
	"github.com/devvdark0/book-library/internal/model"
	"github.com/devvdark0/book-library/internal/repository"
)

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) bookService {
	return bookService{
		repo: repo,
	}
}

func (b bookService) CreateBook(req model.CreateBookRequest) (model.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b bookService) GetBook() (model.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b bookService) ListBooks() ([]model.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b bookService) UpdateBook() (model.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b bookService) DeleteBook() error {
	//TODO implement me
	panic("implement me")
}
