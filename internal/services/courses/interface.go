package courses

import (
	"cybersafe-backend-api/internal/models"

	"github.com/google/uuid"
)

type CoursesManager interface {
	ListWithPagination(int, int) ([]models.CourseExtraFields, int)
	ListByCategoryWithPagination(int, int) ([]models.CourseByCategoryFields, int)
	GetByID(uuid.UUID) (models.Course, error)
	Create(*models.Course) error
	Delete(uuid.UUID) error
	Update(*models.Course) (int, error)
	IsRightAnswer(*models.Answer) bool
	UpdateEnrollmentProgress(uuid.UUID, uuid.UUID)
	GetEnrollmentProgress(uuid.UUID, uuid.UUID) (models.Enrollment, error)
	GetQuestionsByCourseID(uuid.UUID) ([]models.Question, error)
	GetReviewsByCourseID(uuid.UUID) ([]models.Review, error)
}
