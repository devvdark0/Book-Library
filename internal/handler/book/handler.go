package book

import (
	"encoding/json"
	"errors"
	bookErr "github.com/devvdark0/book-library/internal/errors/book"
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
	books, err := b.service.ListBooks()
	if err != nil {
		if errors.Is(err, bookErr.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&books); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (b bookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	uuidId, err := uuid.Parse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	book, err := b.service.GetBook(uuidId)
	if err != nil {
		if errors.Is(err, bookErr.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&book); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (b bookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	uuidId, err := uuid.Parse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var updatedBookReq model.UpdateBookRequest
	if err := json.NewDecoder(r.Body).Decode(&updatedBookReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := b.service.UpdateBook(uuidId, updatedBookReq); err != nil {
		if errors.Is(err, bookErr.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
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
