package router

import (
	"structured-api/api/modules/book"
	"structured-api/api/modules/health"
	"structured-api/api/router/middleware"

	"github.com/go-chi/chi/v5"
	mdw "github.com/go-chi/chi/v5/middleware"
	chizero "github.com/ironstar-io/chizerolog"
	"github.com/rs/zerolog"
)

func New(l *zerolog.Logger) *chi.Mux {
	r := chi.NewRouter()

	// Global Middlewares
	r.Use(mdw.RedirectSlashes)
	r.Use(middleware.RequestID)
	r.Use(middleware.ContentTypeJSON)
	r.Use(chizero.LoggerMiddleware(l))

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
