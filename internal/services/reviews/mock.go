package reviews

import (
	"cybersafe-backend-api/internal/models"

	"github.com/google/uuid"
)

type ReviewsManagerMock struct {
	CreateMock                    func(*models.Review) error
	ExistsByUserIDAndCourseIDMock func(uuid.UUID, uuid.UUID) bool
}

func (rmm *ReviewsManagerMock) Create(review *models.Review) error {
	return rmm.CreateMock(review)
}

func (rmm *ReviewsManagerMock) ExistsByUserIDAndCourseID(userID, courseID uuid.UUID) bool {
	return rmm.ExistsByUserIDAndCourseIDMock(userID, courseID)
}
