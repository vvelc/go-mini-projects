package main

import (
	"log"
	"net/http"
	"structured-api/internal/config"
	"structured-api/internal/server"
	"structured-api/api/router"
	"structured-api/util/logger"
)

func main() {
	l := logger.New(config.GetConfig().IS_DEBUG)
	r := router.New(l)

	s := server.NewServer(r)
	log.Printf("Listening on: http://localhost:%d 🚀", config.GetConfig().PORT)
	err := s.HttpServer.ListenAndServe()

	if err != http.ErrServerClosed {
		log.Fatalf("Listen: %s\n", err)
	}

	log.Println("service stopped")
}
