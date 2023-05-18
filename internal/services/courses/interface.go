package courses

import (
	"cybersafe-backend-api/internal/models"

	"github.com/google/uuid"
)

type CoursesManager interface {
	ListWithPagination(int, int) ([]models.Course, int)
	GetByID(uuid.UUID) (models.Course, error)
	Create(*models.Course) error
	Delete(uuid.UUID) error
	Update(*models.Course) (int, error)
}
