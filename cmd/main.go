package main

import (
	"github.com/devvdark0/book-library/internal/config"
	"github.com/devvdark0/book-library/internal/handler"
	bookHandler "github.com/devvdark0/book-library/internal/handler/book"
	bookRepository "github.com/devvdark0/book-library/internal/repository/book"
	bookService "github.com/devvdark0/book-library/internal/service/book"
	"github.com/devvdark0/book-library/pkg/database"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	cfg := config.Load()

	db, err := database.ConfigureDb(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := bookRepository.NewPostgresBookRepository(db)
	serv := bookService.NewBookService(repo)
	h := bookHandler.NewBookHandler(serv)

	r := configureRouter(h)

	if err := http.ListenAndServe(":80", r); err != nil {
		log.Fatal(err)
	}

}

func configureRouter(h handler.Handler) *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/books").Subrouter()
	s.HandleFunc("", h.ListBooks).Methods(http.MethodGet)
	s.HandleFunc("", h.CreateBook).Methods(http.MethodPost)
	s.HandleFunc("/{id}", h.GetBook).Methods(http.MethodGet)
	s.HandleFunc("/{id}", h.UpdateBook).Methods(http.MethodPut)
	s.HandleFunc("/{id}", h.DeleteBook).Methods(http.MethodDelete)
	return r
}
