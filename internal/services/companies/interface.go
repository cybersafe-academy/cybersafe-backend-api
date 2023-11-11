package companies

import (
	"cybersafe-backend-api/internal/models"

	"github.com/google/uuid"
)

type CompaniesManager interface {
	ListWithPagination(int, int) ([]models.Company, int)
	GetByID(uuid.UUID) (models.Company, error)
	GetByCNPJ(string) (models.Company, error)
	Create(*models.Company) error
	Delete(uuid.UUID) error
	Update(*models.Company) (int, error)
	UpdateContentRecommendations(recommendations []*models.CompanyContentRecommendation) error
}
