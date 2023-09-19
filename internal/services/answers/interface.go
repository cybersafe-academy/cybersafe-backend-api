package answers

import "cybersafe-backend-api/internal/models"

type AnswersManager interface {
	SaveUserAnswer(*models.UserAnswer) error
}
