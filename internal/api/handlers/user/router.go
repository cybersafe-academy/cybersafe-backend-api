package user

import (
	"cybersafe-backend-api/internal/api/components"
	"net/http"

	"github.com/go-chi/chi"
)

func SetupRoutes(c *components.Components) http.Handler {

	subRouter := chi.NewMux()

	// subRouter.Use(middlewares.Authenticator)

	subRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		ListUsersHandler(components.HttpComponents(w, r, c))
	})
	subRouter.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		GetUserByIDHandler(components.HttpComponents(w, r, c))
	})
	subRouter.Post("/", func(w http.ResponseWriter, r *http.Request) {
		CreateUserHandler(components.HttpComponents(w, r, c))
	})
	subRouter.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
		DeleteUserHandler(components.HttpComponents(w, r, c))
	})
	subRouter.Put("/{id}", func(w http.ResponseWriter, r *http.Request) {
		UpdateUserHandler(components.HttpComponents(w, r, c))
	})

	return subRouter

}
