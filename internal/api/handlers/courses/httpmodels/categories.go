package httpmodels

import (
	"cybersafe-backend-api/internal/models"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryRequest struct {
	CategoryFields
}

type CategoryFields struct {
	Name string `json:"name" valid:"required"`
}

type CategoryResponse struct {
	CategoryFields

	ID        uuid.UUID      `json:"id"`
	CreatedAt time.Time      `json:"createdAt,omitempty"`
	UpdatedAt time.Time      `json:"updatedAt,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty"`
}

type CourseCategoryResponse struct {
	CategoryFields

	ID uuid.UUID `json:"id"`
}

func (cr *CategoryRequest) Bind(_ *http.Request) error {
	_, err := govalidator.ValidateStruct(*cr)

	if err != nil {
		return err
	}

	return err
}

func (cr *CategoryRequest) ToEntity() *models.Category {
	category := &models.Category{
		Name: cr.Name,
	}

	return category
}
