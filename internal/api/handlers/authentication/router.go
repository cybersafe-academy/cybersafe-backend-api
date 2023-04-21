package authentication

import (
	"cybersafe-backend-api/internal/api/components"
	"cybersafe-backend-api/internal/api/server/middlewares"
	"net/http"

	"github.com/go-chi/chi"
)

func SetupRoutes(c *components.Components) http.Handler {

	subRouter := chi.NewMux()

	subRouter.Group(func(r chi.Router) {
		r.Use(middlewares.Authorizer())

		r.Post("/logoff", func(w http.ResponseWriter, r *http.Request) {
			LogOffHandler(components.HttpComponents(w, r, c))
		})

		r.Get("/refresh", func(w http.ResponseWriter, r *http.Request) {
			RefreshTokenHandler(components.HttpComponents(w, r, c))
		})
	})

	subRouter.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		LoginHandler(components.HttpComponents(w, r, c))
	})

	return subRouter
}
