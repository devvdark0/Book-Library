package repository

import (
	"github.com/devvdark0/book-library/internal/model"
	"github.com/google/uuid"
)

type Repository interface {
	Create(book model.Book) error
	List() ([]model.Book, error)
	Get(id uuid.UUID) (model.Book, error)
	Update(id uuid.UUID, book model.Book) error
	Delete(id uuid.UUID) error
}
