package server

import (
	"cybersafe-backend-api/pkg/components"

	"github.com/go-chi/chi"
)

func Config(c *components.Components) {
	c.Logger.Info().Msg("[SERVICE] : Configuring Service")

	c.Router = chi.NewRouter()

	Routes(c)
}
