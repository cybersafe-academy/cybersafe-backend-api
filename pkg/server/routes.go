package server

import (
	"cybersafe-backend-api/pkg/components"
	service "cybersafe-backend-api/pkg/services"

	"github.com/go-chi/chi"
)

func Routes(c *components.Components) {

	c.Router.Route("/api", func(r chi.Router) {
		r.Mount("/", service.SetupRoutes(c))
	})

	// Components.Router.Get("/swagger/*", httpSwagger.Handler())
}
