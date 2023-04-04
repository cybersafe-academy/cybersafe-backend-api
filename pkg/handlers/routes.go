package handlers

import (
	_ "cybersafe-backend-api/docs"
	"cybersafe-backend-api/pkg/components"
	"cybersafe-backend-api/pkg/handlers/user"
	"net/http"

	"github.com/go-chi/chi"
)

func SetupRoutes(c *components.Components) http.Handler {
	subRouter := chi.NewRouter()

	subRouter.Mount("/users", user.SetupRoutes(c))
	// subRouter.Mount("/users", user.SetupRoutes(c))
	// subRouter.Mount("/users", user.SetupRoutes(c))
	// subRouter.Mount("/users", user.SetupRoutes(c))
	// subRouter.Mount("/users", user.SetupRoutes(c))
	// subRouter.Mount("/users", user.SetupRoutes(c))

	return subRouter
}
