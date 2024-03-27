package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const PORT = 3000

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome!"))
	})

	// Books routes
	r.Get("/books", func(w http.ResponseWriter, r *http.Request) {
		booksList, err := ListBooks()

		if err != nil {
			http.Error(w, http.StatusText(500), 500)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&booksList)
	})

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("."))
	})

	host := fmt.Sprintf(":%d", PORT)

	// Serve HTTP Server
	err := http.ListenAndServe(host, r)

	if err != nil {
		fmt.Println(err.Error())
	}
}
