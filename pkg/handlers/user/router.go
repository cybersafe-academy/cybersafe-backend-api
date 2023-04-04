package user

import (
	"cybersafe-backend-api/pkg/components"
	"cybersafe-backend-api/pkg/server/middlewares"
	"net/http"

	"github.com/go-chi/chi"
)

func SetupRoutes(c *components.Components) http.Handler {

	subRouter := chi.NewMux()

	subRouter.Use(middlewares.Authenticate)
	subRouter.Use(middlewares.PaginationParams)

	subRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		ListUsersHandler(components.HttpComponents(w, r, c))
	})

	return subRouter

}
