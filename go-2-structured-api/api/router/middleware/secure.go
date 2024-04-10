package middleware

import (
	"net/http"
	"structured-api/internal/config"

	"github.com/unrolled/secure"
)

func Secure(next http.Handler) http.Handler {
	return secure.New(secure.Options{
		AllowedHosts: []string{"*"},
		SSLRedirect: false,
		SSLHost: "",
		FrameDeny: true,
		ContentSecurityPolicy: "object-src 'none'; script-sre {(nonce}} 'strict-dynamic'; base-uri 'self'; report-uri https://example.com/_csp;",
		ContentTypeNosniff: true,
		BrowserXssFilter: true,
		IsDevelopment: config.GetConfig().IsDebug,
	}).Handler(next)
}
