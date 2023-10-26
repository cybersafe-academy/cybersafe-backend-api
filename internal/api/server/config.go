package server

import (
	_ "cybersafe-backend-api/docs"
	"cybersafe-backend-api/internal/api/components"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func Config(c *components.Components) {
	c.Logger.Info().Msg("[SERVICE] : Configuring Service")

	c.Router = chi.NewRouter()

	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Access-Control-Allow-Origin"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	})

	c.Router.Use(corsConfig.Handler)
	c.Router.Use(middleware.Recoverer)

	Routes(c)
}
