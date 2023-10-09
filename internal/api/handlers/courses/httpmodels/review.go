package httpmodels

import (
	"cybersafe-backend-api/internal/models"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReviewFields struct {
	Comment string `json:"comment" valid:"required"`
	Rating  int    `json:"rating" valid:"range(1|5), required"`
}

type UserReviewFields struct {
	ID   uuid.UUID `json:"id" valid:"uuid"`
	Name string    `json:"name" valid:"type(string)"`
}

type ReviewResponse struct {
	ReviewFields

	ID        uuid.UUID      `json:"id" valid:"uuid, required"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`

	CourseID uuid.UUID `json:"courseID" valid:"required"`

	User UserReviewFields `json:"user"`
}

type ReviewRequestContent struct {
	ReviewFields
}

func (rre *ReviewRequestContent) Bind(_ *http.Request) error {
	_, err := govalidator.ValidateStruct(*rre)

	if err != nil {
		return err
	}

	return err
}

func (rrc *ReviewRequestContent) ToEntityReview() *models.Review {
	review := &models.Review{
		Comment: rrc.Comment,
		Rating:  rrc.Rating,
	}
	return review
}
