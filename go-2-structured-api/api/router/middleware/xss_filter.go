package middleware

import (
	"bytes"
	"io"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
)

func sanitizeString(data string) string {
	data = strings.Replace(data, "<", "&lt;", -1)
	data = strings.Replace(data, ">", "&gt;", -1)
	data = strings.Replace(data, "`", "&#96;", -1)
	data = strings.Replace(data, "=", "&#61;", -1)
	return data
}

func CleanXSS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Sanitize request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Error().Msg("error reading body")
		}
		sanitizedBody := sanitizeString(string(body))
		defer r.Body.Close()

		r.Body = io.NopCloser(bytes.NewBufferString(sanitizedBody))

		// Sanitize query params (assuming map[string][]string)
		queryParams := r.URL.Query()

		for key, values := range queryParams {
			sanitizedValues := make([]string, len(values))
			for i, val := range values {
				sanitizedValues[i] = sanitizeString(val)
			}
			queryParams[key] = sanitizedValues
		}
		r.URL.RawQuery = queryParams.Encode()

		//Sanitize path params (assuming map[string]string)
		//   pathParams := chi.RouteParams(r)
		//   for key, value := range pathParams {
		// 	sanitizedValue := sanitizeString(value)
		// 	pathParams[key] = sanitizedValue
		//   }
		//   chi.RouteParams(r.Context, pathParams)

		// Call the next handler with sanitized request
		next.ServeHTTP(w, r)
	})
}
