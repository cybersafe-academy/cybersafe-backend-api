package companies

import (
	"cybersafe-backend-api/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CompaniesManagerDB struct {
	DBConnection *gorm.DB
}

func (cm *CompaniesManagerDB) GetByCNPJ(cnpj string) (models.Company, error) {
	company := models.Company{}
	result := cm.DBConnection.Where("CNPJ = ?", cnpj).First(&company)

	return company, result.Error
}

func (cm *CompaniesManagerDB) Create(company *models.Company) error {
	result := cm.DBConnection.Create(company)
	return result.Error
}

func (cm *CompaniesManagerDB) Delete(id uuid.UUID) error {
	result := cm.DBConnection.Delete(&models.Company{}, id)
	return result.Error
}

func (cm *CompaniesManagerDB) Update(company *models.Company) (int, error) {
	result := cm.DBConnection.Model(company).Clauses(clause.Returning{}).Updates(company)
	return int(result.RowsAffected), result.Error
}