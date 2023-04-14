package authentication

import (
	"net/http"

	"github.com/asaskevich/govalidator"
)

type LoginRequest struct {
	Password string `json:"password"`
	CPF      string `json:"cpf" valid:"type(string), cpf, required"`
}

type TokenContent struct {
	AccessToken string  `json:"accessToken"`
	TokenType   string  `json:"tokenType"`
	ExpiresIn   float64 `json:"expiresIn"`
}

func (re *LoginRequest) Bind(_ *http.Request) error {
	_, err := govalidator.ValidateStruct(*re)
	if err != nil {
		return err
	}

	return err
}
