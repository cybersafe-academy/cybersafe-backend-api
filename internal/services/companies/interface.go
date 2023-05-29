package companies

import (
	"cybersafe-backend-api/internal/models"

	"github.com/google/uuid"
)

type CompaniesManager interface {
	GetByCNPJ(string) (models.Company, error)
	Create(*models.Company) error
	Delete(uuid.UUID) error
	Update(*models.Company) (int, error)
}
