package services

import (
	"cybersafe-backend-api/internal/services/answers"
	"cybersafe-backend-api/internal/services/categories"
	"cybersafe-backend-api/internal/services/companies"
	"cybersafe-backend-api/internal/services/courses"
	"cybersafe-backend-api/internal/services/reviews"
	"cybersafe-backend-api/internal/services/users"
)

type Resources struct {
	Users      users.UsersManager
	Courses    courses.CoursesManager
	Companies  companies.CompaniesManager
	Reviews    reviews.ReviewsManager
	Categories categories.CategoriesManager
	Answers    answers.AnswersManager
}
