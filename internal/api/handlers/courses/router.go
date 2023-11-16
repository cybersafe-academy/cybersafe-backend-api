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

		r.Get("/management", func(w http.ResponseWriter, r *http.Request) {
			ListCoursesHandler(components.HttpComponents(w, r, c))
		})

		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			CreateCourseHandler(components.HttpComponents(w, r, c))
		})

		r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
			DeleteCourseHandler(components.HttpComponents(w, r, c))
		})

		r.Put("/{id}", func(w http.ResponseWriter, r *http.Request) {
			UpdateCourseHandler(components.HttpComponents(w, r, c))
		})

		r.Post("/categories", func(w http.ResponseWriter, r *http.Request) {
			CreateCourseCategory(components.HttpComponents(w, r, c))
		})

		r.Put("/categories/{id}", func(w http.ResponseWriter, r *http.Request) {
			UpdateCategoryHandler(components.HttpComponents(w, r, c))
		})

		r.Delete("/categories/{id}", func(w http.ResponseWriter, r *http.Request) {
			DeleteCategoryHandler(components.HttpComponents(w, r, c))
		})
	})

	subRouter.Group(func(r chi.Router) {
		r.Use(middlewares.Authorizer(c))

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			FetchCoursesHandler(components.HttpComponents(w, r, c))
		})

		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			GetCourseByID(components.HttpComponents(w, r, c))
		})

		r.Get("/{id}/questions", func(w http.ResponseWriter, r *http.Request) {
			GetQuestionsByCourseID(components.HttpComponents(w, r, c))
		})

		r.Get("/{id}/reviews", func(w http.ResponseWriter, r *http.Request) {
			GetReviewsByCourseID(components.HttpComponents(w, r, c))
		})

		r.Get("/categories", func(w http.ResponseWriter, r *http.Request) {
			ListCategoriesHandler(components.HttpComponents(w, r, c))
		})

		r.Post("/{id}/reviews", func(w http.ResponseWriter, r *http.Request) {
			CreateCourseReview(components.HttpComponents(w, r, c))
		})

		r.Post("/questions", func(w http.ResponseWriter, r *http.Request) {
			AddAnswer(components.HttpComponents(w, r, c))
		})

		r.Post("/{id}/questions/batch", func(w http.ResponseWriter, r *http.Request) {
			AddAnswersBatch(components.HttpComponents(w, r, c))
		})

		r.Post("/{id}/comments", func(w http.ResponseWriter, r *http.Request) {
			AddComment(components.HttpComponents(w, r, c))
		})

		r.Get("/{id}/comments", func(w http.ResponseWriter, r *http.Request) {
			GetCommentsByCourse(components.HttpComponents(w, r, c))
		})

		r.Get("/{id}/enrollments", func(w http.ResponseWriter, r *http.Request) {
			GetEnrollmentInfo(components.HttpComponents(w, r, c))
		})

		r.Get("/enrolled", func(w http.ResponseWriter, r *http.Request) {
			GetEnrolledCourses(components.HttpComponents(w, r, c))
		})

		r.Post("/{id}/enroll", func(w http.ResponseWriter, r *http.Request) {
			Enroll(components.HttpComponents(w, r, c))
		})

		r.Post("/{id}/withdraw", func(w http.ResponseWriter, r *http.Request) {
			Withdraw(components.HttpComponents(w, r, c))
		})

		r.Post("/{courseID}/comments/{commentID}/likes", func(w http.ResponseWriter, r *http.Request) {
			AddLikeToComment(components.HttpComponents(w, r, c))
		})
	})

	return subRouter
}
