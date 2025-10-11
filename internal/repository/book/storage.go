package book

import (
	"database/sql"
	"errors"
	"fmt"
	bookErr "github.com/devvdark0/book-library/internal/errors/book"
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
	_, err := b.db.Exec(
		`INSERT INTO book(id, title, description, author_name, year, created_at) VALUES($1, $2, $3, $4, $5, $6)`,
		book.ID,
		book.Title,
		book.Description,
		book.AuthorName,
		book.Year,
		book.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (b bookPostgresRepository) List() ([]model.Book, error) {
	var books []model.Book
	rows, err := b.db.Query(`SELECT * FROM book`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var book model.Book
		if err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Description,
			&book.AuthorName,
			&book.Year,
			&book.CreatedAt,
		); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (b bookPostgresRepository) Get(id uuid.UUID) (model.Book, error) {
	var book model.Book
	err := b.db.QueryRow(
		`SELECT * FROM book WHERE id = $1`,
		id,
	).Scan(
		&book.ID,
		&book.Title,
		&book.Description,
		&book.AuthorName,
		&book.Year,
		&book.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Book{}, bookErr.ErrNotFound
		}
		return model.Book{}, err
	}
	return book, nil
}

func (b bookPostgresRepository) Update(id uuid.UUID, book model.Book) error {
	//TODO implement me
	panic("implement me")
}

func (b bookPostgresRepository) Delete(id uuid.UUID) error {
	_, err := b.db.Exec(`DELETE FROM book WHERE id=$1`, id)
	if err != nil {
		return fmt.Errorf("failed to delete book")
	}
	return nil
}
