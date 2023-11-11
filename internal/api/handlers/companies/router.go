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

	subRouter.Group(
		func(r chi.Router) {
			r.Use(middlewares.Authorizer(c, models.MasterUserRole))

			r.Post("/", func(w http.ResponseWriter, r *http.Request) {
				CreateCompanyHandler(components.HttpComponents(w, r, c))
			})

			r.Put("/{id}", func(w http.ResponseWriter, r *http.Request) {
				UpdateCompanyHandler(components.HttpComponents(w, r, c))
			})

			r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
				DeleteCompanyHandler(components.HttpComponents(w, r, c))
			})
		},
	)

	subRouter.Group(
		func(r chi.Router) {
			r.Use(middlewares.Authorizer(c, models.AdminUserRole, models.MasterUserRole))

			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				ListCompaniesHandler(components.HttpComponents(w, r, c))
			})

			r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
				GetCompanyByIdHandler(components.HttpComponents(w, r, c))
			})

			r.Get("/{id}/content-recommendations/{mbti}", func(w http.ResponseWriter, r *http.Request) {
				GetCompanyContentRecommendationsHandler(components.HttpComponents(w, r, c))
			})

			r.Put("/{id}/content-recommendations", func(w http.ResponseWriter, r *http.Request) {
				UpdateCompanyContentRecommendationsHandler(components.HttpComponents(w, r, c))
			})
		},
	)

	return subRouter
}
