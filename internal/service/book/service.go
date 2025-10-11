package book

import (
	"fmt"
	"github.com/devvdark0/book-library/internal/model"
	"github.com/devvdark0/book-library/internal/repository"
	"github.com/devvdark0/book-library/internal/service"
	"github.com/google/uuid"
	"time"
)

type bookService struct {
	repo repository.Repository
}

func NewBookService(repo repository.Repository) service.Service {
	return bookService{repo: repo}
}

func (b bookService) CreateBook(req model.CreateBookRequest) error {
	if req.Year < 0 {
		return fmt.Errorf("invalid year of book")
	}

	id := uuid.New()
	book := model.Book{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		AuthorName:  req.Author,
		Year:        req.Year,
		CreatedAt:   time.Now(),
	}
	if err := b.repo.Create(book); err != nil {
		return err
	}
	return nil
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
