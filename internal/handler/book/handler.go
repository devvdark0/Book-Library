package book

import (
	"github.com/devvdark0/book-library/internal/handler"
	"net/http"
)

type bookHandler struct {
	service BookService
}

func NewBookHandler(serv BookService) handler.Handler {
	return bookHandler{service: serv}
}

func (b bookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (b bookHandler) ListBooks(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (b bookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (b bookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (b bookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
