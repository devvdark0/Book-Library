package book

import (
	"database/sql"
	"errors"
	"fmt"
	bookErr "github.com/devvdark0/book-library/internal/errors/book"
	"github.com/devvdark0/book-library/internal/logger"
	"github.com/devvdark0/book-library/internal/model"
	"github.com/devvdark0/book-library/internal/repository"
	"github.com/google/uuid"
	"strings"
)

type bookPostgresRepository struct {
	db  *sql.DB
	log logger.Logger
}

func NewPostgresBookRepository(db *sql.DB, log logger.Logger) repository.Repository {
	return bookPostgresRepository{db: db, log: log}
}

func (b bookPostgresRepository) Create(book model.Book) error {
	b.log.Debug("creating book with title:", book.Title)
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
		b.log.Warn("book was not created")
		return err
	}
	return nil
}

func (b bookPostgresRepository) List() ([]model.Book, error) {
	var books []model.Book
	b.log.Debug("fetching all books from db...")
	rows, err := b.db.Query(`SELECT * FROM book`)
	if err != nil {
		b.log.Warn("cannot complete query:", err)
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
			b.log.Error("error getting data:", err)
			return nil, err
		}
		books = append(books, book)
	}
	b.log.Debug("get all books")
	return books, nil
}

func (b bookPostgresRepository) Get(id uuid.UUID) (model.Book, error) {
	var book model.Book
	b.log.Debug("fetching book from db,", "bookId:", id)
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
			b.log.Warn("no books with id:", id)
			return model.Book{}, bookErr.ErrNotFound
		}
		b.log.Error("failed to complete query:", err)
		return model.Book{}, err
	}
	b.log.Debug("get book with id:", id)
	return book, nil
}

func (b bookPostgresRepository) Update(id uuid.UUID, book model.Book) error {
	b.log.Warn("start updating book with id:", id)
	var clauses []string
	var args []interface{}
	counter := 1
	if book.Title != "" {
		clauses = append(clauses, fmt.Sprintf("title=$%d", counter))
		args = append(args, book.Title)
		counter++
	}
	if book.Description != "" {
		clauses = append(clauses, fmt.Sprintf("description=$%d", counter))
		args = append(args, book.Description)
		counter++
	}
	if book.AuthorName != "" {
		clauses = append(clauses, fmt.Sprintf("author_name=$%d", counter))
		args = append(args, book.AuthorName)
		counter++
	}
	if book.Year != 0 {
		clauses = append(clauses, fmt.Sprintf("year=$%d", counter))
		args = append(args, book.Year)
		counter++
	}
	query := fmt.Sprintf("UPDATE book SET %s WHERE id=%q", strings.Join(clauses, ", "), id)
	_, err := b.db.Exec(query, args)
	if err != nil {
		if errors.Is(err, bookErr.ErrNotFound) {
			b.log.Warn("no books with id:", id)
			return bookErr.ErrNotFound
		}
		b.log.Error("updating book with id", id, "was not finished")
		return err
	}
	b.log.Debug("book with id", id, "was updated")
	return nil
}

func (b bookPostgresRepository) Delete(id uuid.UUID) error {
	b.log.Debug("start deleting book with id:", id)
	_, err := b.db.Exec(`DELETE FROM book WHERE id=$1`, id)
	if err != nil {
		if errors.Is(err, bookErr.ErrNotFound) {
			b.log.Warn("not books with id:", id)
			return bookErr.ErrNotFound
		}
		b.log.Error("deleting book was not finished:", err)
		return err
	}
	b.log.Debug("book with id", id, "was deleted")
	return nil
}
