package book

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/devvdark0/book-library/internal/model"
	"log"
	"strings"
)

var ErrNotFound = errors.New("book not found")

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) postgresRepository {
	return postgresRepository{db: db}
}

func (p postgresRepository) Create(book model.Book) (model.Book, error) {
	log.Print("start creating a book...")
	var id int64
	err := p.db.QueryRow(
		`INSERT INTO book(name, description, author, year, created_at) VALUES($1, $2, $3, $4, $5) RETURNING id`,
		book.Name,
		book.Description,
		book.Author,
		book.Year,
		book.CreatedAt,
	).Scan(&id)
	if err != nil {
		log.Print(err)
		return model.Book{}, err
	}

	book.ID = id
	return book, nil
}

func (p postgresRepository) List() ([]model.Book, error) {
	var books []model.Book
	rows, err := p.db.Query(`SELECT * FROM book`)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	for rows.Next() {
		var book model.Book
		if err := rows.Scan(
			&book.ID,
			&book.Name,
			&book.Description,
			&book.Author,
			&book.Year,
			&book.CreatedAt,
		); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, err
}

func (p postgresRepository) Get(id int64) (model.Book, error) {
	var book model.Book

	err := p.db.QueryRow(`SELECT * FROM book WHERE id=$1`, id).
		Scan(&book.ID,
			&book.Name,
			&book.Description,
			&book.Author,
			&book.Year,
			&book.CreatedAt,
		)
	if err != nil {
		log.Print(err)
		return model.Book{}, err
	}
	return book, nil
}

func (p postgresRepository) Update(id int64, book model.Book) (model.Book, error) {
	var query strings.Builder
	var updates []interface{}
	if book.Name != "" {
		query.WriteString("name=?")
		updates = append(updates, book.Name)
	}
	if book.Description != "" {
		query.WriteString("description=?")
		updates = append(updates, book.Description)
	}
	if book.Author != "" {
		query.WriteString("author=?")
		updates = append(updates, book.Author)
	}
	if book.Year != 0 {
		query.WriteString("year=?")
		updates = append(updates, book.Year)
	}
	updates = append(updates, id)
	result, err := p.db.Exec(
		fmt.Sprintf("UPDATE book SET %s", query.String()),
		updates...,
	)
	if err != nil {
		log.Print(err)
		return model.Book{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return model.Book{}, err
	}
	if rowsAffected == 0 {
		return model.Book{}, ErrNotFound
	}
	book.ID = id
	return book, nil
}

func (p postgresRepository) Delete(id int64) error {
	res, err := p.db.Exec(`DELETE FROM book WHERE id=$1`, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}
