package book

import (
	"database/sql"
	"github.com/devvdark0/book-library/internal/model"
	"log"
)

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

func (p postgresRepository) Update() {
	//TODO implement me
	panic("implement me")
}

func (p postgresRepository) Delete() {
	//TODO implement me
	panic("implement me")
}
