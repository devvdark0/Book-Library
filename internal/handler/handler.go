package handler

import "net/http"

type Handler interface {
	CreateBook(w http.ResponseWriter, r *http.Request)
	ListBooks(w http.ResponseWriter, r *http.Request)
	GetBook(w http.ResponseWriter, r *http.Request)
	UpdateBook(w http.ResponseWriter, r *http.Request)
	DeleteBook(w http.ResponseWriter, r *http.Request)
}
