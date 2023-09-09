package categories

import (
	"cybersafe-backend-api/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CategoriesManagerDB struct {
	DBConnection *gorm.DB
}

func (cm *CategoriesManagerDB) ListWithPagination(offset, limit int) ([]models.Category, int) {
	var categories []models.Category
	var count int64

	cm.DBConnection.
		Offset(offset).
		Limit(limit).
		Find(&categories).
		Count(&count)

	return categories, int(count)
}

func (cm *CategoriesManagerDB) GetByID(id uuid.UUID) (models.Category, error) {
	var category models.Category

	result := cm.DBConnection.First(&category, id)

	return category, result.Error
}

func (cm *CategoriesManagerDB) Create(category *models.Category) error {
	result := cm.DBConnection.Create(category)

	return result.Error
}

func (cm *CategoriesManagerDB) Delete(id uuid.UUID) error {
	result := cm.DBConnection.Delete(&models.Category{}, id)

	return result.Error
}

func (cm *CategoriesManagerDB) Update(category *models.Category) (int, error) {
	result := cm.DBConnection.Model(category).Clauses(clause.Returning{}).Updates(category)

	return int(result.RowsAffected), result.Error
}
