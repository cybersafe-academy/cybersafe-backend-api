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

func (cm *CompaniesManagerDB) UpdateContentRecommendations(recommendations []models.CompanyContentRecommendation) error {
	// Start a transaction
	tx := cm.DBConnection.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	for _, recommendation := range recommendations {
		// Delete existing records that match mbti_type and company_id
		err := tx.Where("mbti_type = ? AND company_id = ?", recommendation.MBTIType, recommendation.CompanyID).Delete(&models.CompanyContentRecommendation{}).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// Insert new recommendations
	result := tx.Create(recommendations)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	// Commit the transaction
	return tx.Commit().Error
}

func (cm *CompaniesManagerDB) GetCompanyContentRecommendationsByMBTI(companyID uuid.UUID, mbti string) ([]models.CompanyContentRecommendation, error) {
	var recommendations []models.CompanyContentRecommendation
	result := cm.DBConnection.
		Preload("Category").
		Where("company_id = ? AND mbti_type = UPPER(?)", companyID, mbti).
		Find(&recommendations)

	return recommendations, result.Error
}
