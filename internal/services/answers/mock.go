package answers

import "cybersafe-backend-api/internal/models"

type AnswersManagerMock struct {
	SaveUserAnswerMock func(*models.UserAnswer) error
}

func (amm *AnswersManagerMock) SaveUserAnswer(answer *models.UserAnswer) error {
	return amm.SaveUserAnswerMock(answer)
}
