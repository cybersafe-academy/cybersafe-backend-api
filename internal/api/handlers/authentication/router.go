package authentication

import (
	"cybersafe-backend-api/internal/api/components"
	"net/http"

	"github.com/go-chi/chi"
)

func SetupRoutes(c *components.Components) http.Handler {

	subRouter := chi.NewMux()

	subRouter.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		LoginHandler(components.HttpComponents(w, r, c))
	})

	return subRouter

}
