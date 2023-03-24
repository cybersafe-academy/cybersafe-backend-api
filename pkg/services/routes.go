package service

import (
	"cybersafe-backend-api/pkg/components"
	"cybersafe-backend-api/pkg/services/user"
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
