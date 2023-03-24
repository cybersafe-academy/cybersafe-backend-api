package user

import (
	"cybersafe-backend-api/pkg/components"
	"net/http"

	"github.com/go-chi/chi"
)

func SetupRoutes(c *components.Components) http.Handler {

	subRouter := chi.NewRouter()

	subRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		GetAll(components.HttpComponents(w, r, c))
	})
	// subRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	GetAll(components.HttpComponents(w, r, c))
	// })
	// subRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	GetAll(components.HttpComponents(w, r, c))
	// })

	return subRouter

}
