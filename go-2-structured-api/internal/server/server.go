package server

import (
	"fmt"
	"net/http"
	"structured-api/internal/config"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	HttpServer *http.Server
	router     *chi.Mux
}

func NewServer(r *chi.Mux) *Server {
	host := fmt.Sprintf(":%d",  config.GetConfig().PORT)

	return &Server{
		HttpServer: &http.Server{Addr: host, Handler: r},
		router:     r,
	}
}