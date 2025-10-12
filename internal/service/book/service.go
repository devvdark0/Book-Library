package book

import (
	"errors"
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
	books, err := b.repo.List()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (b bookService) GetBook(id uuid.UUID) (model.Book, error) {
	book, err := b.repo.Get(id)
	if err != nil {
		return model.Book{}, err
	}
	return book, nil
}

func (b bookService) UpdateBook(id uuid.UUID, req model.UpdateBookRequest) error {
	if id == uuid.Nil {
		return errors.New("not valid uuid")
	}
	if *req.Year < 0 {
		return errors.New("invalid year of book")
	}
	newBook := model.Book{
		ID:          id,
		Title:       *req.Title,
		Description: *req.Description,
		AuthorName:  *req.Author,
		Year:        *req.Year,
	}
	if err := b.repo.Update(id, newBook); err != nil {
		return err
	}
	return nil
}

func (b bookService) DeleteBook(id uuid.UUID) error {
	if id == uuid.Nil {
		id = uuid.New()
	}

	if err := b.repo.Delete(id); err != nil {
		return err
	}
	return nil
}
