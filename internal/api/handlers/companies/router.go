package companies

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
		r.Use(middlewares.Authorizer(c, models.MasterUserRole))

		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			CreateCompanyHandler(components.HttpComponents(w, r, c))
		})
	})

	return subRouter
}
