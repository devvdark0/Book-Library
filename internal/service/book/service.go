package book

import (
	"errors"
	"fmt"
	"github.com/devvdark0/book-library/internal/logger"
	"github.com/devvdark0/book-library/internal/model"
	"github.com/devvdark0/book-library/internal/repository"
	"github.com/devvdark0/book-library/internal/service"
	"github.com/google/uuid"
	"time"
)

type bookService struct {
	repo repository.Repository
	log  logger.Logger
}

func NewBookService(repo repository.Repository, log logger.Logger) service.Service {
	return bookService{repo: repo, log: log}
}

func (b bookService) CreateBook(req model.CreateBookRequest) error {
	b.log.Info("creating book...")
	if req.Year < 0 {
		b.log.Error("invalid year of book", req.Year)
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
		b.log.Error("book was not created...")
		return err
	}
	b.log.Info("book was successfully created!")
	return nil
}

func (b bookService) ListBooks() ([]model.Book, error) {
	b.log.Info("listing books...")
	books, err := b.repo.List()
	if err != nil {
		b.log.Error("failed to get all the books")
		return nil, err
	}
	b.log.Info("successfully listing of books")
	return books, nil
}

func (b bookService) GetBook(id uuid.UUID) (model.Book, error) {
	b.log.Info("getting book with id", id)
	book, err := b.repo.Get(id)
	if err != nil {
		b.log.Error("failed to get the book with id", id)
		return model.Book{}, err
	}
	b.log.Info("book with id", id, "successfully gotten")
	return book, nil
}

func (b bookService) UpdateBook(id uuid.UUID, req model.UpdateBookRequest) error {
	b.log.Info("updating book with id", id)
	if id == uuid.Nil {
		b.log.Error("invalid id", id)
		return errors.New("not valid uuid")
	}
	if *req.Year < 0 {
		b.log.Error("invalid year of book", *req.Year)
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
		b.log.Error("failed to update book with id", id)
		return err
	}
	b.log.Info("book with id", id, "successfully updated!")
	return nil
}

func (b bookService) DeleteBook(id uuid.UUID) error {
	b.log.Info("deleting book with id", id)
	if id == uuid.Nil {
		id = uuid.New()
	}

	if err := b.repo.Delete(id); err != nil {
		b.log.Error("failed to delete the book with id", id)
		return err
	}
	b.log.Info("book with id", id, "was successfully deleted")
	return nil
}
