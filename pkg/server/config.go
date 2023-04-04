package server

import (
	_ "cybersafe-backend-api/docs"
	"cybersafe-backend-api/pkg/components"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Config(c *components.Components) {
	c.Logger.Info().Msg("[SERVICE] : Configuring Service")

	c.Router = chi.NewRouter()
	c.Router.Use(middleware.Recoverer)

	Routes(c)
}
