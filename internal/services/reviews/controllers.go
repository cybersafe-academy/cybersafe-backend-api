package reviews

import (
	"cybersafe-backend-api/internal/models"

	"gorm.io/gorm"
)

type ReviewsManagerDB struct {
	DBConnection *gorm.DB
}

func (rmm *ReviewsManagerDB) Create(review *models.Review) error {
	result := rmm.DBConnection.Create(review)
	return result.Error
}
