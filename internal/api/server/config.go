package server

import (
	_ "cybersafe-backend-api/docs"
	"cybersafe-backend-api/internal/api/components"
	"cybersafe-backend-api/pkg/environment"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func Config(c *components.Components) {
	c.Logger.Info().Msg("[SERVICE] : Configuring Service")

	c.Router = chi.NewRouter()

	allowedOrigins := []string{}

	if c.Environment == environment.Local {
		allowedOrigins = append(allowedOrigins, "http://localhost:*")
	}

	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
	})

	c.Router.Use(corsConfig.Handler)
	c.Router.Use(middleware.Recoverer)

	Routes(c)
}
