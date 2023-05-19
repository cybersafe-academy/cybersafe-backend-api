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
	Name      string `json:"name" valid:"type(string),"`
	Role      string `json:"role" valid:"type(string),"`
	Email     string `json:"email" valid:"type(string), email, required"`
	BirthDate string `json:"birthDate" valid:"date"`
	CPF       string `json:"cpf" valid:"type(string), cpf,"`
}

type UserFieldsUpdate struct {
	Name      string `json:"name" valid:"type(string),"`
	Role      string `json:"role" valid:"type(string),"`
	Email     string `json:"email" valid:"type(string), email, required"`
	BirthDate string `json:"birthDate" valid:"type(date)"`
	CPF       string `json:"cpf" valid:"type(string), cpf,"`
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

type RequestContentUpdate struct {
	UserFieldsUpdate
	Password string `json:"password" valid:"stringlength(8|24)"`
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
func (re *RequestContentUpdate) Bind(_ *http.Request) error {

	if !govalidator.IsIn(re.Role, models.ValidUserRoles...) {
		return errutil.ErrInvalidUserRole
	}

	_, err := govalidator.ValidateStruct(*re)
	if err != nil {
		return err
	}

	return err
}

func (re *RequestContent) ToEntity() *models.User {

	birthDate, _ := time.Parse(helpers.DefaultDateFormat(), re.BirthDate)

	return &models.User{
		Name:      re.Name,
		Email:     re.Email,
		BirthDate: birthDate,
		CPF:       re.CPF,
		Role:      re.Role,
		Password:  re.Password,
	}
}

func (re *RequestContentUpdate) ToEntity() *models.User {

	birthDate, _ := time.Parse(helpers.DefaultDateFormat(), re.BirthDate)

	return &models.User{
		Name:      re.Name,
		Email:     re.Email,
		BirthDate: birthDate,
		CPF:       re.CPF,
		Role:      re.Role,
		Password:  re.Password,
	}
}
