package courses

import (
	"cybersafe-backend-api/internal/api/components"
	"cybersafe-backend-api/internal/api/server/middlewares"
	"net/http"

	"github.com/go-chi/chi"
)

func SetupRoutes(c *components.Components) http.Handler {

	subRouter := chi.NewMux()

	subRouter.Use(middlewares.Authorizer(c))

	subRouter.Post("/", func(w http.ResponseWriter, r *http.Request) {
		CreateCourseHandler(components.HttpComponents(w, r, c))
	})
	subRouter.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
		DeleteCourseHandler(components.HttpComponents(w, r, c))
	})
	subRouter.Put("/{id}", func(w http.ResponseWriter, r *http.Request) {
		UpdateCourseHandler(components.HttpComponents(w, r, c))
	})

	subRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		ListCoursesHandler(components.HttpComponents(w, r, c))
	})
	subRouter.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		GetCourseByID(components.HttpComponents(w, r, c))
	})

	return subRouter

}
