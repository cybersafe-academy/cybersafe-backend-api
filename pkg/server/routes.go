package server

import (
	"cybersafe-backend-api/pkg/components"
	"cybersafe-backend-api/pkg/handlers"

	"github.com/go-chi/chi"
)

func Routes(c *components.Components) {

	c.Router.Route("/api", func(r chi.Router) {
		r.Mount("/", handlers.SetupRoutes(c))
	})

	// Components.Router.Get("/swagger/*", httpSwagger.Handler())
}
