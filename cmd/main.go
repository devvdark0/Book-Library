package main

import (
	"github.com/devvdark0/book-library/internal/handler/book"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	serv := NewBookService()
	h := book.NewBookHandler(serv)
	r.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello into Book Libary!"))
	})
	s := r.PathPrefix("/books").Subrouter()
	s.HandleFunc("", h.ListBooks)
	s.HandleFunc("", h.CreateBook)
	s.HandleFunc("/{id}", h.GetBook)
	s.HandleFunc("/{id}", h.UpdateBook)
	s.HandleFunc("/{id}", h.DeleteBook)

	if err := http.ListenAndServe(":80", r); err != nil {
		log.Fatal(err)
	}

}
