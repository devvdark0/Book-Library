package repository

import "github.com/devvdark0/book-library/internal/model"

type BookRepository interface {
	Create(book model.Book) (model.Book, error)
	List() ([]model.Book, error)
	Get(id int64) (model.Book, error)
	Update(id int64, book model.Book) (model.Book, error)
	Delete(id int64) error
}
