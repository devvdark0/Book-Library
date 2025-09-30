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
	var updatedBook model.Book
	var query []string
	var updates []interface{}
	counter := 1

	if book.Name != "" {
		q := fmt.Sprintf("name=$%d", counter)
		query = append(query, q)
		updates = append(updates, book.Name)
		counter++
	}
	if book.Description != "" {
		q := fmt.Sprintf("description=$%d", counter)
		query = append(query, q)
		updates = append(updates, book.Description)
		counter++
	}
	if book.Author != "" {
		q := fmt.Sprintf("author=$%d", counter)
		query = append(query, q)
		updates = append(updates, book.Author)
		counter++
	}
	if book.Year != 0 {
		q := fmt.Sprintf("year=$%d", counter)
		query = append(query, q)
		updates = append(updates, book.Year)
		counter++
	}

	updates = append(updates, id)
	err := p.db.QueryRow(
		fmt.Sprintf(`UPDATE book SET %s WHERE id=$%d RETURNING name,description,author,year,created_at`,
			strings.Join(query, ", "), counter), updates...,
	).Scan(&updatedBook.Name,
		&updatedBook.Description,
		&updatedBook.Author,
		&updatedBook.Year,
		&updatedBook.CreatedAt,
	)
	if err != nil {
		log.Print(err)
		return model.Book{}, err
	}
	updatedBook.ID = id
	return updatedBook, nil
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
