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

func NewServer(c *config.Config, r *chi.Mux) *Server {
	host := fmt.Sprintf(":%d", c.Port)

	return &Server{
		HttpServer: &http.Server{Addr: host, Handler: r},
		router:     r,
	}
}
