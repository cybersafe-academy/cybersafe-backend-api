package answers

import (
	"cybersafe-backend-api/internal/models"

	"gorm.io/gorm"
)

type AnswersManagerDB struct {
	DBConnection *gorm.DB
}

func (am *AnswersManagerDB) SaveUserAnswer(userAnswer *models.UserAnswer) error {
	result := am.DBConnection.Create(userAnswer)
	return result.Error
}
