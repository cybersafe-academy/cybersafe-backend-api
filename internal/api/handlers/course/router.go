package course

import (
	"cybersafe-backend-api/internal/api/components"
	"cybersafe-backend-api/internal/api/server/middlewares"
	"net/http"

	"github.com/go-chi/chi"
)

func SetupRoutes(c *components.Components) http.Handler {

	subRouter := chi.NewMux()

	subRouter.Use(middlewares.Authenticator)

	subRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		ListCoursesHandler(components.HttpComponents(w, r, c))
	})

	return subRouter

}
