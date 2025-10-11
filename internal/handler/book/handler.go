package book

import (
	"encoding/json"
	"github.com/devvdark0/book-library/internal/handler"
	"github.com/devvdark0/book-library/internal/model"
	"github.com/devvdark0/book-library/internal/service"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type bookHandler struct {
	service service.Service
}

func NewBookHandler(serv service.Service) handler.Handler {
	return bookHandler{service: serv}
}

func (b bookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var createBookReq model.CreateBookRequest
	if err := json.NewDecoder(r.Body).Decode(&createBookReq); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := b.service.CreateBook(createBookReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
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
	id := mux.Vars(r)["id"]
	uuidId, err := uuid.Parse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	if err := b.service.DeleteBook(uuidId); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
