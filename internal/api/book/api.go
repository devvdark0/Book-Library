package book

import (
	"encoding/json"
	"github.com/devvdark0/book-library/internal/model"
	"github.com/devvdark0/book-library/internal/service"
	"github.com/gorilla/mux"
	"net/http"
)

type bookAPI struct {
	bookService service.BookService
}

func NewAPI(service service.BookService, router *mux.Router) bookAPI {
	handler := bookAPI{bookService: service}

	router.HandleFunc("/book", handler.List).Methods(http.MethodGet)
	router.HandleFunc("/book", handler.Create).Methods(http.MethodPost)
	router.HandleFunc("/book/{id}", handler.Get).Methods(http.MethodGet)
	router.HandleFunc("/book/{id}", handler.Update).Methods(http.MethodPatch)
	router.HandleFunc("/book/{id}", handler.Remove).Methods(http.MethodDelete)

	return handler
}

func (b bookAPI) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateBookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	book, err := b.bookService.CreateBook(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(book); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (b bookAPI) List(w http.ResponseWriter, _ *http.Request) {
	books, err := b.bookService.ListBooks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(&books); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (b bookAPI) Get(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	book, err := b.bookService.GetBook(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(book); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (b bookAPI) Update(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var payload model.UpdateBookRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	book, err := b.bookService.UpdateBook(id, payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(&book); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (b bookAPI) Remove(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := b.bookService.DeleteBook(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
