package server

import (
	"cybersafe-backend-api/internal/api/components"
	"cybersafe-backend-api/internal/api/domains"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Routes(c *components.Components) {

	c.Router.Route("/api", func(r chi.Router) {
		r.Mount("/", domains.SetupRoutes(c))
	})

	c.Router.Get("/*", httpSwagger.Handler())

}
