package main

import (
	"log"
	"net/http"
	"structured-api/internal/config"
	"structured-api/internal/server"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	s := server.NewServer(r)
	log.Printf("Listening on: http://localhost%d 🚀", config.GetConfig().PORT)
	err := s.HttpServer.ListenAndServe()

	if err != http.ErrServerClosed {
	 log.Fatalf("Listen: %s\n", err)
	}

	log.Println("service stopped")
   
   }

