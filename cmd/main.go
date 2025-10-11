package main

import (
	handler "github.com/devvdark0/book-library/internal/handler/book"
	repository "github.com/devvdark0/book-library/internal/repository/book"
	service "github.com/devvdark0/book-library/internal/service/book"
	"github.com/devvdark0/book-library/pkg/database"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	db, err := database.ConfigureDb()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewPostgresBookRepository(db)
	serv := service.NewBookService(repo)
	h := handler.NewBookHandler(serv)

	r := mux.NewRouter()
	s := r.PathPrefix("/books").Subrouter()
	s.HandleFunc("", h.ListBooks).Methods(http.MethodGet)
	s.HandleFunc("", h.CreateBook).Methods(http.MethodPost)
	s.HandleFunc("/{id}", h.GetBook).Methods(http.MethodGet)
	s.HandleFunc("/{id}", h.UpdateBook).Methods(http.MethodPut)
	s.HandleFunc("/{id}", h.DeleteBook).Methods(http.MethodDelete)

	if err := http.ListenAndServe(":80", r); err != nil {
		log.Fatal(err)
	}

}
