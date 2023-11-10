package reviews

import (
	"cybersafe-backend-api/internal/models"

	"github.com/google/uuid"
)

type ReviewsManager interface {
	Create(*models.Review) error
	ExistsByUserIDAndCourseID(uuid.UUID, uuid.UUID) bool
}
