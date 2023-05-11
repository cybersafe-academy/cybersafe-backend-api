package server

import (
	"cybersafe-backend-api/internal/api/components"
	"cybersafe-backend-api/internal/api/handlers"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Routes(c *components.Components) {

	c.Router.Route("/api", func(r chi.Router) {
		r.Mount("/", handlers.SetupRoutes(c))
	})

	c.Router.Get("/*", httpSwagger.Handler())

}
