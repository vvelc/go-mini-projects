package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const PORT = 3000 // This goest to internal or config

// TODO: Here goes server setup function

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome!"))
	})

	host := fmt.Sprintf(":%d", PORT)

	// Serve HTTP Server
	err := http.ListenAndServe(host, r)

	if err != nil {
		fmt.Println(err.Error())
	}
}

