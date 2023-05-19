package authentication

import (
	"cybersafe-backend-api/internal/models"
	"cybersafe-backend-api/pkg/errutil"
	"net/http"

	"github.com/asaskevich/govalidator"
)

type LoginRequest struct {
	Password string `json:"password"`
	CPF      string `json:"cpf"`
}

type TokenContent struct {
	AccessToken string  `json:"accessToken"`
	TokenType   string  `json:"tokenType"`
	ExpiresIn   float64 `json:"expiresIn"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" valid:"type(string), email, required"`
}

type FirstAccessRequest struct {
	Email string `json:"email" valid:"type(string), email, required"`
}

type UpdatePasswordRequest struct {
	Password string `json:"password" valid:"stringlength(8|24)"`
}

type FinishSignupRequest struct {
	Name      string `json:"name" valid:"type(string),"`
	Role      string `json:"role" valid:"type(string),"`
	BirthDate string `json:"birthDate" valid:"type(date)"`
	CPF       string `json:"cpf" valid:"type(string), cpf,"`
	Password  string `json:"password" valid:"stringlength(8|24)"`
}

func (re *FinishSignupRequest) Bind(_ *http.Request) error {

	if !govalidator.IsIn(re.Role, models.ValidUserRoles...) {
		return errutil.ErrInvalidUserRole
	}

	_, err := govalidator.ValidateStruct(*re)
	if err != nil {
		return err
	}

	return err
}

func (re *LoginRequest) Bind(_ *http.Request) error {
	_, err := govalidator.ValidateStruct(*re)
	if err != nil {
		return err
	}

	return err
}

func (re *ForgotPasswordRequest) Bind(_ *http.Request) error {
	_, err := govalidator.ValidateStruct(*re)
	if err != nil {
		return err
	}

	return err
}

func (re *UpdatePasswordRequest) Bind(_ *http.Request) error {
	_, err := govalidator.ValidateStruct(*re)
	if err != nil {
		return err
	}

	return err
}

func (re *FirstAccessRequest) Bind(_ *http.Request) error {
	_, err := govalidator.ValidateStruct(*re)
	if err != nil {
		return err
	}

	return err
}
