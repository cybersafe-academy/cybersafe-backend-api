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
	Name              string `json:"name"`
	Role              string `json:"role"`
	Email             string `json:"email" valid:"email, required"`
	BirthDate         string `json:"birthDate" valid:"date"`
	CPF               string `json:"cpf" valid:"cpf"`
	ProfilePictureURL string `json:"profilePictureURL" valid:"url"`
}

type ResponseContent struct {
	UserFields

	ID        uuid.UUID      `json:"id" valid:"uuid, required"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

type RequestContent struct {
	UserFields
	Password string `json:"password" valid:"stringlength(8|24)"`
}

type PreSignupRequest struct {
	Role  string `json:"role"`
	Email string `json:"email" valid:"email, required"`
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
	isValidMBTIType := isValidMBTIType(ptr.MBTIType)
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
		Password:          re.Password,
	}
}
