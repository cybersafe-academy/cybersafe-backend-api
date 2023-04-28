package user

import (
	"cybersafe-backend-api/internal/api/components"
	"cybersafe-backend-api/internal/api/server/middlewares"
	"cybersafe-backend-api/internal/models"
	"net/http"

	"github.com/go-chi/chi"
)

func SetupRoutes(c *components.Components) http.Handler {

	subRouter := chi.NewMux()

	subRouter.Group(func(r chi.Router) {

		r.Use(middlewares.Authorizer(c, models.AdminUserRole, models.MasterUserRole))

		r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
			DeleteUserHandler(components.HttpComponents(w, r, c))
		})
		r.Put("/{id}", func(w http.ResponseWriter, r *http.Request) {
			UpdateUserHandler(components.HttpComponents(w, r, c))
		})
	})

	subRouter.Group(func(r chi.Router) {
		r.Use(middlewares.Authorizer(c))

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			ListUsersHandler(components.HttpComponents(w, r, c))
		})
		r.Get("/me", func(w http.ResponseWriter, r *http.Request) {
			GetAuthenticatedUserHandler(components.HttpComponents(w, r, c))
		})
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			GetUserByIDHandler(components.HttpComponents(w, r, c))
		})
	})

	subRouter.Post("/", func(w http.ResponseWriter, r *http.Request) {
		CreateUserHandler(components.HttpComponents(w, r, c))
	})

	return subRouter

}
