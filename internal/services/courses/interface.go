package courses

import (
	"cybersafe-backend-api/internal/models"

	"github.com/google/uuid"
)

type CoursesManager interface {
	ListWithPagination(int, int) ([]models.CourseExtraFields, int)
	ListCoursesWithRecommendation(userID, companyID uuid.UUID, userMBTIType string) ([]models.CourseExtraFields, error)
	GetEnrolledCourses(uuid.UUID) []models.Course
	GetByID(uuid.UUID) (models.Course, error)
	Create(*models.Course) error
	Delete(uuid.UUID) error
	Update(*models.Course) (int, error)
	IsRightAnswer(*models.Answer) bool
	UpdateEnrollmentProgress(uuid.UUID, uuid.UUID)
	UpdateEnrollmentStatus(uuid.UUID, uuid.UUID) (float64, error)
	Enroll(*models.Enrollment) error
	Withdraw(uuid.UUID, uuid.UUID) error
	GetEnrollmentProgress(uuid.UUID, uuid.UUID) (models.Enrollment, error)
	GetQuestionsByCourseID(uuid.UUID) ([]models.Question, error)
	GetReviewsByCourseID(uuid.UUID) ([]models.Review, error)
	AddComment(*models.Comment) error
	ListCommentsByCourse(uuid.UUID) []models.Comment
	AddLikeToComment(uuid.UUID, uuid.UUID) error
	ExistsEnrollmentByUserIDAndCourseID(uuid.UUID, uuid.UUID) bool
}
