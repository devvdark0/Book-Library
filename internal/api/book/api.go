package book

import (
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
	router.HandleFunc("book/{id}", handler.Get).Methods(http.MethodGet)
	router.HandleFunc("/book/{id}", handler.Update).Methods(http.MethodPatch)
	router.HandleFunc("/book/{id}", handler.Remove).Methods(http.MethodDelete)

	return handler
}

func (b bookAPI) Create(w http.ResponseWriter, r *http.Request) {

}

func (b bookAPI) List(w http.ResponseWriter, r *http.Request) {

}

func (b bookAPI) Get(w http.ResponseWriter, r *http.Request) {

}

func (b bookAPI) Update(w http.ResponseWriter, r *http.Request) {

}

func (b bookAPI) Remove(w http.ResponseWriter, r *http.Request) {

}
