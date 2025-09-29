package main

import (
	bookAPI "github.com/devvdark0/book-library/internal/api/book"
	bookRepo "github.com/devvdark0/book-library/internal/repository/book"
	bookService "github.com/devvdark0/book-library/internal/service/book"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	store := bookRepo.NewPostgresRepository(nil)
	service := bookService.NewBookService(store)
	bookAPI.NewAPI(service, r)

	if err := http.ListenAndServe(":80", r); err != nil {
		log.Fatal(err)
	}
}
