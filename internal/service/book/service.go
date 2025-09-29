package book

import (
	"github.com/devvdark0/book-library/internal/model"
	"github.com/devvdark0/book-library/internal/repository"
	"time"
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
	book := model.Book{
		Name:        req.Name,
		Description: req.Description,
		Author:      req.Author,
		Year:        req.Year,
		CreatedAt:   time.Now(),
	}
	book, err := b.repo.Create(book)
	if err != nil {
		return model.Book{}, err
	}
	return book, nil
}

func (b bookService) GetBook() (model.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b bookService) ListBooks() ([]model.Book, error) {
	books, err := b.repo.List()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (b bookService) UpdateBook() (model.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b bookService) DeleteBook() error {
	//TODO implement me
	panic("implement me")
}
