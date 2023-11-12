package companies

import (
	"cybersafe-backend-api/internal/models"
	"cybersafe-backend-api/pkg/errutil"
	"cybersafe-backend-api/pkg/helpers"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompanyFields struct {
		LegalName string `json:"legalName" valid:"required"`
		TradeName string `json:"tradeName"`
		CNPJ      string `json:"cnpj" valid:"cnpj, required"`
		Email     string `json:"email" valid:"email, required"`
		Phone     string `json:"phone"`
	}

type ResponseContent struct {
	CompanyFields

	ID        uuid.UUID      `json:"id" valid:"uuid, required"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

type RequestContent struct {
	CompanyFields
}

type CompanyContentRecommendationFields struct {
	MBTIType   string    `json:"mbtiType" valid:"required"`
	Categories []string  `json:"categories" valid:"uuid, required"`
	CompanyID  uuid.UUID `json:"-"`
}

type CompanyContentRecommendationResponseContent struct {
	CompanyContentRecommendationFields

	CompanyID uuid.UUID `json:"companyID"`

	ID        uuid.UUID      `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

type CategoryResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type CompanyContentRecommendationByMBTIResponseContent struct {
	CompanyID  uuid.UUID          `json:"companyID"`
	MbtiType   string             `json:"mbtiType"`
	Categories []CategoryResponse `json:"categories"`

	ID        uuid.UUID      `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

type CompanyContentRecommendationRequestContent struct {
	CompanyContentRecommendationFields
}

type DefaultContentRecommendationFields struct {
	MBTIType   string      `json:"mbtiType" valid:"required"`
	Categories []uuid.UUID `json:"categories" valid:"uuid, required"`
}

type DefaultContentRecommendationResponseContent struct {
	DefaultContentRecommendationFields

	ID        uuid.UUID      `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

type DefaultContentRecommendationRequestContent struct {
	DefaultContentRecommendationFields
}

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

func (re *CompanyContentRecommendationRequestContent) Bind(_ *http.Request) error {
	isValidMBTIType := helpers.IsValidMBTIType(re.MBTIType)
	if !isValidMBTIType {
		return errutil.ErrInvalidMBTIType
	}

	_, err := govalidator.ValidateStruct(*re)
	if err != nil {
		return err
	}

	return err
}

func (re *CompanyContentRecommendationRequestContent) ToEntity() []models.CompanyContentRecommendation {
	formattedCompanyContentRecommendations := []models.CompanyContentRecommendation{}

	for _, categoryID := range re.Categories {
		formattedCompanyContentRecommendations = append(
			formattedCompanyContentRecommendations, models.CompanyContentRecommendation{
				MBTIType:   re.MBTIType,
				CompanyID:  re.CompanyID,
				CategoryID: uuid.MustParse(categoryID),
			},
		)
	}

	return formattedCompanyContentRecommendations
}

func (re *DefaultContentRecommendationRequestContent) Bind(_ *http.Request) error {
	isValidMBTIType := helpers.IsValidMBTIType(re.MBTIType)
	if !isValidMBTIType {
		return errutil.ErrInvalidMBTIType
	}

	_, err := govalidator.ValidateStruct(*re)
	if err != nil {
		return err
	}

	return err
}

func (re *DefaultContentRecommendationRequestContent) ToEntity() *[]models.DefaultContentRecommendation {
	formattedDefaultContentRecommendation := make([]models.DefaultContentRecommendation, len(re.Categories))

	for _, categoryID := range re.Categories {
		formattedDefaultContentRecommendation = append(
			formattedDefaultContentRecommendation, models.DefaultContentRecommendation{
				MBTIType:   re.MBTIType,
				CategoryID: categoryID,
			},
		)
	}

	return &formattedDefaultContentRecommendation
}
