package reviews

import (
	"cybersafe-backend-api/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ReviewsManagerDB struct {
	DBConnection *gorm.DB
}

func (rmm *ReviewsManagerDB) Create(review *models.Review) error {
	result := rmm.DBConnection.Clauses(clause.Returning{}).Create(review)
	return result.Error
}

func (rmm *ReviewsManagerDB) ExistsByUserIDAndCourseID(userID, courseID uuid.UUID) bool {
	result := rmm.DBConnection.
		Where("user_id = ?", userID).
		Where("course_id = ?", courseID).
		First(&models.Review{})

	return result.Error != nil
}
