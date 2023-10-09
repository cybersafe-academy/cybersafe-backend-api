package categories

import (
	"cybersafe-backend-api/internal/models"

	"github.com/google/uuid"
)

type CategoriesManager interface {
	ListWithPagination(int, int) ([]models.Category, int)
	GetByID(uuid.UUID) (models.Category, error)
	Create(*models.Category) error
	Delete(uuid.UUID) error
	Update(*models.Category) (int, error)
}
