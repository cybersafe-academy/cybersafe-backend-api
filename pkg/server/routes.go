package server

import (
	"cybersafe-backend-api/pkg/components"
	service "cybersafe-backend-api/pkg/services"

	"github.com/go-chi/chi"
)

func Routes(Components *components.Components) {

	// Components.Logger.Info().Msg("[SERVICE] : Setup routes")
	Components.Router.Route("/api", func(r chi.Router) {
		r.Mount("/", service.SetupRoutes(Components))
	})

	// Components.Router.Get("/swagger/*", httpSwagger.Handler())
}
