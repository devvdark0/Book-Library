package book

import (
	"database/sql"
	"github.com/devvdark0/book-library/internal/model"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) postgresRepository {
	return postgresRepository{db: db}
}

func (p postgresRepository) Create(book model.Book) (model.Book, error) {
	stmt, err := p.db.Prepare(
		`INSERT INTO book(name, description, author, year) VALUES(?, ?, ?, ?) RETURNING id`,
	)
	if err != nil {
		return model.Book{}, err
	}
	res, err := stmt.Exec(book.Name, book.Description, book.Author, book.Year)
	if err != nil {
		return model.Book{}, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return model.Book{}, err
	}
	book.ID = id
	return book, nil
}

func (p postgresRepository) List() ([]model.Book, error) {
	var books []model.Book
	rows, err := p.db.Query(`SELECT * FROM book`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(books); err != nil {
			return nil, err
		}
	}
	return books, err
}

func (p postgresRepository) Get() {
	//TODO implement me
	panic("implement me")
}

func (p postgresRepository) Update() {
	//TODO implement me
	panic("implement me")
}

func (p postgresRepository) Delete() {
	//TODO implement me
	panic("implement me")
}
