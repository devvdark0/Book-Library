package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello into Book Libary!"))
	})

	if err := http.ListenAndServe(":80", r); err != nil {
		log.Fatal(err)
	}

}
