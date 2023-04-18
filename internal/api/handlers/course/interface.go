package course

import (
	"cybersafe-backend-api/internal/models"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CourseFields struct {
	Name           string  `json:"name" valid:"type(string), required" example:"Example Course"`
	Description    string  `json:"description" valid:"type(string), required" example:"Course Description"`
	ContentInHours float64 `json:"contentInHours" valid:"type(float), required" example:"24.5"`
	ThumbnailURL   string  `json:"thumbnailURL" valid:"type(string), required" example:"https://image.com"`
	Level          string  `json:"level" valid:"type(string), required" example:"advanced"`

	Contents []ContentFields `json:"contents" valid:"type(string), required"`
}

type ContentFields struct {
	ID uuid.UUID `json:"id"`

	ContentType string `json:"contentType" valid:"type(string), required" example:"youtube"`
	YoutubeURL  string `json:"youtubeURL" valid:"type(string), required" example:"https://www.youtube.com/watch?v=mvV7tzRm8Pk"`
	PDFURL      string `json:"PDFURL" valid:"type(string), required" example:"https://pdf.com"`
	ImageURL    string `json:"imageURL" valid:"type(string), required" example:"https://image.com"`
}

type ResponseContent struct {
	CourseFields

	ID        uuid.UUID      `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

type RequestContent struct {
	CourseFields
}

func (re *RequestContent) Bind(_ *http.Request) error {
	_, err := govalidator.ValidateStruct(*re)
	if err != nil {
		return err
	}

	return err
}

func (re *RequestContent) ToEntity() *models.Course {
	return &models.Course{}
}
