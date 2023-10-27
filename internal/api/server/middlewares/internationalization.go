package middlewares

import (
	"cybersafe-backend-api/internal/api/components"
	"net/http"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func LocalizerMiddleware(c *components.Components) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			lang := r.Header.Get("Accept-Language")
			localizer := i18n.NewLocalizer(c.Bundle, lang)

			c.Localizer = localizer

			next.ServeHTTP(w, r)
		})
	}
}
