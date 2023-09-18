package courses

import (
	"cybersafe-backend-api/internal/api/components"
	"cybersafe-backend-api/internal/api/server/middlewares"
	"cybersafe-backend-api/internal/models"
	"net/http"

	"github.com/go-chi/chi"
)

func SetupRoutes(c *components.Components) http.Handler {

	subRouter := chi.NewMux()

	subRouter.Group(func(r chi.Router) {
		r.Use(middlewares.Authorizer(c, models.AdminUserRole, models.MasterUserRole))

		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			CreateCourseHandler(components.HttpComponents(w, r, c))
		})
		r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
			DeleteCourseHandler(components.HttpComponents(w, r, c))
		})
		r.Put("/{id}", func(w http.ResponseWriter, r *http.Request) {
			UpdateCourseHandler(components.HttpComponents(w, r, c))
		})
	})

	subRouter.Group(func(r chi.Router) {
		r.Use(middlewares.Authorizer(c))

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			ListCoursesHandler(components.HttpComponents(w, r, c))
		})
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			GetCourseByID(components.HttpComponents(w, r, c))
		})
		r.Post("/{id}/review", func(w http.ResponseWriter, r *http.Request) {
			CreateCourseReview(components.HttpComponents(w, r, c))
		})
		r.Post("/question", func(w http.ResponseWriter, r *http.Request) {
			AddAnswer(components.HttpComponents(w, r, c))
		})
		r.Post("/{id}/comment", func(w http.ResponseWriter, r *http.Request) {
			AddComment(components.HttpComponents(w, r, c))
		})
		r.Get("/{id}/comment", func(w http.ResponseWriter, r *http.Request) {
			GetCommentsByCourse(components.HttpComponents(w, r, c))
		})
		r.Get("/{id}/enrollment", func(w http.ResponseWriter, r *http.Request) {
			GetEnrollmentInfo(components.HttpComponents(w, r, c))
		})
		r.Post("/{id}/comment/like", func(w http.ResponseWriter, r *http.Request) {
			AddLikeToComment(components.HttpComponents(w, r, c))
		})
	})

	return subRouter

}
