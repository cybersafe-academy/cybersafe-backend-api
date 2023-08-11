package reviews

import "cybersafe-backend-api/internal/models"

type ReviewsManager interface {
	Create(*models.Review) error
}
