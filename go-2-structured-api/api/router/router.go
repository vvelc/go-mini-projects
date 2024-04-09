package router

import (
	"structured-api/pkg/book"
	"structured-api/pkg/health"
	"structured-api/pkg/router/middleware"

	"github.com/go-chi/chi/v5"
	mdw "github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
)

func New(l *zerolog.Logger) *chi.Mux {
	r := chi.NewRouter()

	// Global Middlewares
	r.Use(mdw.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.ContentTypeJSON)

	r.Get("/livez", health.Read)

	r.Route("/v1", func(r chi.Router) {
		r.Route("/books", func(r chi.Router) {
			r.Post("/", book.Create)
			r.Get("/", book.GetAll)
			r.Get("/{id}", book.GetOne)
			r.Patch("/{id}", book.Update)
			r.Delete("/{id}", book.Delete)
		})
	})

	return r
}
