package reviews

import (
	"cybersafe-backend-api/internal/models"

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
