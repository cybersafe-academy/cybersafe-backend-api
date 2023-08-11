package reviews

import "cybersafe-backend-api/internal/models"

type ReviewsManagerMock struct {
	CreateMock func(*models.Review) error
}

func (rmm *ReviewsManagerMock) Create(review *models.Review) error {
	return rmm.CreateMock(review)
}
