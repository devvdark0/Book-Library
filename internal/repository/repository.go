package repository

import "github.com/devvdark0/book-library/internal/model"

type BookRepository interface {
	Create(book model.Book) (model.Book, error)
	List() ([]model.Book, error)
	Get()
	Update()
	Delete()
}
