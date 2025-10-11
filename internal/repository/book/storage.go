package book

import (
	"database/sql"
	"github.com/devvdark0/book-library/internal/model"
	"github.com/devvdark0/book-library/internal/repository"
	"github.com/google/uuid"
)

type bookPostgresRepository struct {
	db *sql.DB
}

func NewPostgresBookRepository(db *sql.DB) repository.Repository {
	return bookPostgresRepository{db: db}
}

func (b bookPostgresRepository) Create(book model.Book) error {
	//TODO implement me
	panic("implement me")
}

func (b bookPostgresRepository) List() ([]model.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b bookPostgresRepository) Get(id uuid.UUID) (model.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b bookPostgresRepository) Update(id uuid.UUID, book model.Book) error {
	//TODO implement me
	panic("implement me")
}

func (b bookPostgresRepository) Delete(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
