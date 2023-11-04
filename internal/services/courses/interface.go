package courses

import (
	"cybersafe-backend-api/internal/api/handlers/courses/httpmodels"
	"cybersafe-backend-api/internal/models"

	"github.com/google/uuid"
)

type CoursesManager interface {
	ListWithPagination(int, int) ([]models.CourseExtraFields, int)
	ListByCategory() *httpmodels.CourseByCategoryResponse
	GetEnrolledCourses(uuid.UUID) []models.Course
	GetByID(uuid.UUID) (models.Course, error)
	Create(*models.Course) error
	Delete(uuid.UUID) error
	Update(*models.Course) (int, error)
	IsRightAnswer(*models.Answer) bool
	UpdateEnrollmentProgress(uuid.UUID, uuid.UUID)
	UpdateEnrollmentStatus(uuid.UUID, uuid.UUID) (float64, error)
	Enroll(*models.Enrollment) error
	GetEnrollmentProgress(uuid.UUID, uuid.UUID) (models.Enrollment, error)
	GetQuestionsByCourseID(uuid.UUID) ([]models.Question, error)
	GetReviewsByCourseID(uuid.UUID) ([]models.Review, error)
	AddComment(*models.Comment) error
	ListCommentsByCourse(uuid.UUID) []models.Comment
	AddLikeToComment(uuid.UUID, uuid.UUID) error
}
