package user

import (
	"cybersafe-backend-api/internal/models"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserFields struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	CPF  string `json:"cpf"`
}

type ResponseContent struct {
	UserFields

	ID        uuid.UUID      `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

type RequestContent struct {
	UserFields
	Password string `json:"password"`
}

func (re *RequestContent) Bind(_ *http.Request) error {
	_, err := govalidator.ValidateStruct(*re)
	if err != nil {
		return err
	}

	return err
}

func (re *RequestContent) ToEntity() *models.User {
	return &models.User{
		Name:     re.Name,
		Age:      re.Age,
		CPF:      re.CPF,
		Password: re.Password,
	}
}
