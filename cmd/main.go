package main

import (
	"github.com/devvdark0/book-library/internal/config"
	"github.com/devvdark0/book-library/internal/handler"
	bookHandler "github.com/devvdark0/book-library/internal/handler/book"
	"github.com/devvdark0/book-library/internal/logger"
	bookRepository "github.com/devvdark0/book-library/internal/repository/book"
	bookService "github.com/devvdark0/book-library/internal/service/book"
	"github.com/devvdark0/book-library/pkg/database"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	cfg := config.Load()

	log := logger.InitLogger(cfg)

	log.Info("connecting to the database...")
	db, err := database.ConfigureDb(cfg)
	if err != nil {
		log.Error("failed to connect to db:", err)
		return
	}
	log.Info("Successfully connected to db!")
	defer db.Close()

	repo := bookRepository.NewPostgresBookRepository(db, log)
	serv := bookService.NewBookService(repo, log)
	h := bookHandler.NewBookHandler(serv)

	r := configureRouter(h)
	log.Info("starting server on port:", cfg.Addr)
	if err := http.ListenAndServe(cfg.Addr, r); err != nil {
		log.Error("failed to run the server:", err)
		return
	}
	log.Info("server started!!!")
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
