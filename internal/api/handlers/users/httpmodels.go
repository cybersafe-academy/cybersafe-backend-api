package users

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

type UserFields struct {
	Name              string `json:"name" valid:"required"`
	Role              string `json:"role"`
	Email             string `json:"email" valid:"email, required"`
	BirthDate         string `json:"birthDate" valid:"date, required"`
	CPF               string `json:"cpf" valid:"cpf, required"`
	ProfilePictureURL string `json:"profilePictureURL"`
	MbtiType          string `json:"mbtiType"`
}

type ResponseContent struct {
	UserFields
	CompanyResponse CompanyResponse `json:"company"`

	ID        uuid.UUID      `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

type CompanyResponse struct {
	ID        uuid.UUID `json:"id"`
	LegalName string    `json:"legalName"`
	TradeName string    `json:"tradeName"`
	CNPJ      string    `json:"cnpj"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
}

type RequestContent struct {
	UserFields

	CompanyID uuid.UUID `json:"companyID"`
	Password  string    `json:"password" valid:"stringlength(8|24)"`
}

type PreSignupRequest struct {
	Role      string    `json:"role" valid:"required"`
	Email     string    `json:"email" valid:"email, required"`
	CompanyID uuid.UUID `json:"companyID"`
}

type PersonalityTestRequest struct {
	MBTIType string `json:"mbtiType" valid:"required"`
}

func (re *RequestContent) Bind(_ *http.Request) error {
	if !govalidator.IsIn(re.Role, models.ValidUserRoles...) {
		return errutil.ErrInvalidUserRole
	}

	_, err := govalidator.ValidateStruct(*re)
	if err != nil {
		return err
	}

	return err
}

func (re *PreSignupRequest) Bind(_ *http.Request) error {
	if !govalidator.IsIn(re.Role, models.ValidUserRoles...) {
		return errutil.ErrInvalidUserRole
	}

	_, err := govalidator.ValidateStruct(*re)
	if err != nil {
		return err
	}

	return err
}

func (ptr *PersonalityTestRequest) Bind(_ *http.Request) error {
	isValidMBTIType := helpers.IsValidMBTIType(ptr.MBTIType)
	if !isValidMBTIType {
		return errutil.ErrInvalidMBTIType
	}

	_, err := govalidator.ValidateStruct(*ptr)
	if err != nil {
		return err
	}

	return err
}

func (re *RequestContent) ToEntity() *models.User {
	birthDate, _ := time.ParseInLocation(
		helpers.DefaultDateFormat(),
		re.BirthDate,
		helpers.MustGetAmericaSPTimeZone(),
	)

	return &models.User{
		Name:              re.Name,
		Email:             re.Email,
		BirthDate:         birthDate,
		CPF:               re.CPF,
		ProfilePictureURL: re.ProfilePictureURL,
		Role:              re.Role,
		MBTIType:          re.MbtiType,
		Password:          re.Password,
		CompanyID:         re.CompanyID,
	}
}
