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
	r.Use(chizero.LoggerMiddleware(l))
	r.Use(mdw.Recoverer)
	r.Use(middleware.Secure)
	r.Use(middleware.Cors)
	r.Use(mdw.RedirectSlashes)
	r.Use(middleware.RequestID)
	r.Use(middleware.ContentTypeJSON)

	// Health route
	r.Get("/livez", health.Read)

	// v1 API routes
	r.Route("/v1", func(r chi.Router) {
		// Books routes
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
