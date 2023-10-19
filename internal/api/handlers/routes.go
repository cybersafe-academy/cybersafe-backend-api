package handlers

import (
	_ "cybersafe-backend-api/docs"
	"cybersafe-backend-api/internal/api/components"
	"cybersafe-backend-api/internal/api/handlers/authentication"
	"cybersafe-backend-api/internal/api/handlers/companies"
	"cybersafe-backend-api/internal/api/handlers/courses"
	"cybersafe-backend-api/internal/api/handlers/users"
	"cybersafe-backend-api/internal/api/server/middlewares"
	"net/http"

	"github.com/go-chi/chi"
)

func SetupRoutes(c *components.Components) http.Handler {
	subRouter := chi.NewMux()

	subRouter.Use(middlewares.LocalizerMiddleware(c))

	subRouter.Mount("/auth", authentication.SetupRoutes(c))
	subRouter.Mount("/users", users.SetupRoutes(c))
	subRouter.Mount("/courses", courses.SetupRoutes(c))
	subRouter.Mount("/companies", companies.SetupRoutes(c))

	return subRouter
}
