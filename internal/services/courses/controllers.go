package courses

import (
	"cybersafe-backend-api/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CoursesManagerDB struct {
	DBConnection *gorm.DB
}

func (cm *CoursesManagerDB) ListWithPagination(offset, limit int) ([]models.CourseExtraFields, int) {

	var courses []models.CourseExtraFields
	var count int64

	cm.DBConnection.
		Table("courses").
		Preload("Contents").
		Preload("Reviews").
		Joins("LEFT JOIN reviews ON reviews.course_id = courses.id").
		Select("courses.*, avg(reviews.rating) as avg_rating").
		Where("courses.deleted_at IS NULL").
		Group("courses.id").
		Offset(offset).
		Limit(limit).
		Find(&courses).
		Count(&count)

	return courses, int(count)
}

func (cm *CoursesManagerDB) GetByID(id uuid.UUID) (models.Course, error) {
	var course models.Course

	result := cm.DBConnection.First(&course, id)

	return course, result.Error
}

func (cm *CoursesManagerDB) Create(course *models.Course) error {
	result := cm.DBConnection.Create(course)
	return result.Error
}

func (cm *CoursesManagerDB) Delete(id uuid.UUID) error {
	result := cm.DBConnection.Delete(&models.Course{}, id)
	return result.Error
}

func (cm *CoursesManagerDB) Update(course *models.Course) (int, error) {
	result := cm.DBConnection.Model(course).Clauses(clause.Returning{}).Updates(course)
	return int(result.RowsAffected), result.Error
}
