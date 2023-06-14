package authentication

import (
	"cybersafe-backend-api/pkg/errutil"
	"cybersafe-backend-api/pkg/helpers"
	"net/http"
	"time"

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
	BirthDate string `json:"birthDate" valid:"date"`
	CPF       string `json:"cpf" valid:"type(string), cpf,"`
	Password  string `json:"password" valid:"stringlength(8|24)"`
}

func (re *FinishSignupRequest) Bind(_ *http.Request) error {
	_, err := govalidator.ValidateStruct(*re)
	if err != nil {
		return err
	}

	birthDate, _ := time.ParseInLocation(
		helpers.DefaultDateFormat(),
		re.BirthDate,
		helpers.MustGetAmericaSPTimeZone(),
	)

	if birthDate.After(time.Now()) {
		return errutil.ErrFutureDateNotAllowed
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
