package handlers

import (
	_ "cybersafe-backend-api/docs"
	"cybersafe-backend-api/internal/api/components"
	"cybersafe-backend-api/internal/api/handlers/authentication"
	"cybersafe-backend-api/internal/api/handlers/course"
	"cybersafe-backend-api/internal/api/handlers/user"
	"net/http"

	"github.com/go-chi/chi"
)

func SetupRoutes(c *components.Components) http.Handler {
	subRouter := chi.NewMux()

	subRouter.Mount("/auth", authentication.SetupRoutes(c))
	subRouter.Mount("/users", user.SetupRoutes(c))
	subRouter.Mount("/courses", course.SetupRoutes(c))

	return subRouter
}
