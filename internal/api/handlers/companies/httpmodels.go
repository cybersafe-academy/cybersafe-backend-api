package companies

import (
	"cybersafe-backend-api/internal/models"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	CompanyFields struct {
		LegalName string `json:"legalName" valid:"required"`
		TradeName string `json:"tradeName"`
		CNPJ      string `json:"cnpj" valid:"cnpj, required"`
		Email     string `json:"email" valid:"email, required"`
		Phone     string `json:"phone"`
	}

	ResponseContent struct {
		CompanyFields

		ID        uuid.UUID      `json:"id" valid:"uuid, required"`
		CreatedAt time.Time      `json:"createdAt"`
		UpdatedAt time.Time      `json:"updatedAt"`
		DeletedAt gorm.DeletedAt `json:"deletedAt"`
	}

	RequestContent struct {
		CompanyFields
	}
)

func (re *RequestContent) Bind(_ *http.Request) error {
	_, err := govalidator.ValidateStruct(*re)
	if err != nil {
		return err
	}

	return err
}

func (re *RequestContent) ToEntity() *models.Company {
	return &models.Company{
		LegalName: re.LegalName,
		TradeName: re.TradeName,
		CNPJ:      re.CNPJ,
		Email:     re.Email,
		Phone:     re.Phone,
	}
}
