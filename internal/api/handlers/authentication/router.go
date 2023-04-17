package authentication

import (
	"cybersafe-backend-api/internal/api/components"
	"cybersafe-backend-api/internal/api/server/middlewares"
	"net/http"

	"github.com/go-chi/chi"
)

func SetupRoutes(c *components.Components) http.Handler {

	subRouter := chi.NewMux()

	subRouter.Route("/", func(r chi.Router) {
		subRouter.Use(middlewares.Authenticator)

		subRouter.Post("/logoff", func(w http.ResponseWriter, r *http.Request) {
			LogOffHandler(components.HttpComponents(w, r, c))
		})

		subRouter.Get("/refresh", func(w http.ResponseWriter, r *http.Request) {
			RefreshTokenHandler(components.HttpComponents(w, r, c))
		})
	})

	subRouter.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		LoginHandler(components.HttpComponents(w, r, c))
	})

	return subRouter

}
