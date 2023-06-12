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

func (cm *CompaniesManagerDB) ListWithPagination(offset, limit int) ([]models.Company, int) {
	var companies []models.Company
	var count int64

	cm.DBConnection.Preload(clause.Associations).
		Offset(offset).
		Limit(limit).
		Find(&companies).
		Count(&count)

	return companies, int(count)
}

func (cm *CompaniesManagerDB) GetByID(id uuid.UUID) (models.Company, error) {
	var company models.Company
	result := cm.DBConnection.First(&company, id)
	return company, result.Error
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
