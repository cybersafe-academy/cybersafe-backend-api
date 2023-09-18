package httpmodels

import (
	"cybersafe-backend-api/internal/models"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type CommentFields struct {
	Text string `json:"text"`
}

type CommentResponse struct {
	CommentFields

	ID uuid.UUID `json:"id" valid:"uuid, required"`

	UserID   uuid.UUID `json:"userID"`
	CourseID uuid.UUID `json:"courseID"`

	LikesCount int `json:"likesCount"`
}

type CommentRequest struct {
	CommentFields
}

func (cr *CommentRequest) Bind(_ *http.Request) error {

	_, err := govalidator.ValidateStruct(*cr)
	if err != nil {
		return err
	}

	return err
}

func (cr *CommentRequest) ToEntity() *models.Comment {

	return &models.Comment{
		Text: cr.Text,
	}
}
