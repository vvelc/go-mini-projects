package router

import (
	"encoding/json"
	"io"
	"net/http"
	"structured-api/api/modules/book"
	"structured-api/api/modules/health"
	"structured-api/api/router/middleware"

	"github.com/go-chi/chi/v5"
	mdw "github.com/go-chi/chi/v5/middleware"
	chizero "github.com/ironstar-io/chizerolog"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func New(l *zerolog.Logger) *chi.Mux {
	r := chi.NewRouter()

	// Global Middlewares
	r.Use(chizero.LoggerMiddleware(l))
	r.Use(mdw.Recoverer)
	r.Use(middleware.Secure)
	r.Use(middleware.CleanXSS)
	r.Use(middleware.Cors)
	r.Use(mdw.RedirectSlashes)
	r.Use(middleware.RequestID)
	r.Use(middleware.ContentTypeJSON)

	// Health route
	r.Get("/healthz", health.Read)

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

	r.Post("/testxss/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		if name == "" {
			log.Warn().Msg("empty name param")
		}

		body, err := io.ReadAll(r.Body)
		log.Info().Msg("body to return: " + string(body))
		defer r.Body.Close()
		if err != nil {
			log.Warn().Msg("error reading body")
		}

		query := r.URL.Query().Get("query")
		if query == "" {
			log.Warn().Msg("empty query param")
		}

		var output struct {
			Name  string `json:"name"`
			Body  string    `json:"body"`
			Query string `json:"query"`
		}
		output.Name = name
		output.Body = string(body)
		output.Query = query

		json.NewEncoder(w).Encode(output)
	})

	return r
}
