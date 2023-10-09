package categories

import (
	"cybersafe-backend-api/internal/models"

	"github.com/google/uuid"
)

type CategoriesManagerMock struct {
	ListWithPaginationMock func(int, int) ([]models.Course, int)
	GetByIDMock            func(uuid.UUID) (models.Course, error)
	CreateMock             func(*models.Course) error
	DeleteMock             func(uuid.UUID) error
	UpdateMock             func(*models.Course) (int, error)
}

func (cm *CategoriesManagerMock) ListWithPagination(limit, offset int) ([]models.Course, int) {
	return cm.ListWithPaginationMock(limit, offset)
}
func (cm *CategoriesManagerMock) GetByID(id uuid.UUID) (models.Course, error) {
	return cm.GetByIDMock(id)
}
func (cm *CategoriesManagerMock) Create(course *models.Course) error {
	return cm.CreateMock(course)
}
func (cm *CategoriesManagerMock) Delete(id uuid.UUID) error {
	return cm.DeleteMock(id)
}
func (cm *CategoriesManagerMock) Update(course *models.Course) (int, error) {
	return cm.UpdateMock(course)
}
