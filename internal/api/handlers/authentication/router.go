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
		r.Use(middlewares.Authorizer(c))

		r.Post("/logoff", func(w http.ResponseWriter, r *http.Request) {
			LogOffHandler(components.HttpComponents(w, r, c))
		})

		r.Post("/refresh", func(w http.ResponseWriter, r *http.Request) {
			RefreshTokenHandler(components.HttpComponents(w, r, c))
		})
	})

	subRouter.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		LoginHandler(components.HttpComponents(w, r, c))
	})

	subRouter.Post("/forgot-password", func(w http.ResponseWriter, r *http.Request) {
		ForgotPasswordHandler(components.HttpComponents(w, r, c))
	})

	subRouter.Post("/update-password", func(w http.ResponseWriter, r *http.Request) {
		UpdatePasswordHandler(components.HttpComponents(w, r, c))
	})

	subRouter.Post("/first-access", func(w http.ResponseWriter, r *http.Request) {
		FirstAccessHandler(components.HttpComponents(w, r, c))
	})

	subRouter.Post("/finish-signup", func(w http.ResponseWriter, r *http.Request) {
		FinishSignupHandler(components.HttpComponents(w, r, c))
	})

	return subRouter
}
