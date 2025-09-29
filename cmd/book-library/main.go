package main

import (
	"database/sql"
	bookAPI "github.com/devvdark0/book-library/internal/api/book"
	bookRepo "github.com/devvdark0/book-library/internal/repository/book"
	bookService "github.com/devvdark0/book-library/internal/service/book"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	db, err := configureDatabase()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	store := bookRepo.NewPostgresRepository(db)
	service := bookService.NewBookService(store)
	bookAPI.NewAPI(service, r)

	if err := http.ListenAndServe(":80", r); err != nil {
		log.Fatal(err)
	}
}

func configureDatabase() (*sql.DB, error) {
	db, err := sql.Open(
		"postgres",
		"host=localhost user=vladislav password=mypassword dbname=book_db sslmode=disable",
	)
	defer db.Close()
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
