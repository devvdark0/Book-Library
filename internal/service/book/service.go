package book

import (
	"github.com/devvdark0/book-library/internal/model"
	"github.com/devvdark0/book-library/internal/repository"
	"strconv"
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

func (b bookService) GetBook(id string) (model.Book, error) {
	i, err := strconv.Atoi(id)
	if err != nil {
		return model.Book{}, err
	}
	book, err := b.repo.Get(int64(i))
	if err != nil {
		return model.Book{}, err
	}
	return book, nil
}

func (b bookService) ListBooks() ([]model.Book, error) {
	books, err := b.repo.List()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (b bookService) UpdateBook(id string, dto model.UpdateBookRequest) (model.Book, error) {
	normalizeId, err := strconv.Atoi(id)
	if err != nil {
		return model.Book{}, err
	}
	var book model.Book
	if dto.Name != nil {
		book.Name = *dto.Name
	}
	if dto.Description != nil {
		book.Description = *dto.Description
	}
	if dto.Author != nil {
		book.Author = *dto.Author
	}
	if dto.Year != nil {
		book.Year = *dto.Year
	}

	updatedBook, err := b.repo.Update(int64(normalizeId), book)
	if err != nil {
		return model.Book{}, err
	}
	return updatedBook, nil
}

func (b bookService) DeleteBook(id string) error {
	normalizedId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	if err = b.repo.Delete(int64(normalizedId)); err != nil {
		return err
	}
	return nil
}
