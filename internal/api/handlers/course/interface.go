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
	Name           string  `json:"name" valid:"type(string), required"`
	Description    string  `json:"description" valid:"type(string), required"`
	ContentInHours float64 `json:"contentInHours" valid:"type(float), required"`
	ThumbnailURL   string  `json:"thumbnailURL" valid:"type(string), required"`
	Level          string  `json:"level" valid:"type(string), required"`

	Contents []ContentFields `json:"contents"`
}

type ContentFields struct {
	ID uuid.UUID `json:"id" valid:"uuid, required"`

	ContentType string `json:"contentType" valid:"type(string), required"`
	YoutubeURL  string `json:"youtubeURL" valid:"type(string), required"`
	PDFURL      string `json:"PDFURL" valid:"type(string), required"`
	ImageURL    string `json:"imageURL" valid:"type(string), required"`
}

type ResponseContent struct {
	CourseFields

	ID        uuid.UUID      `json:"id" valid:"uuid, required"`
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
	course := &models.Course{
		Name:           re.Name,
		Description:    re.Description,
		ContentInHours: re.ContentInHours,
		ThumbnailURL:   re.ThumbnailURL,
		Level:          re.Level,
	}

	for _, content := range re.Contents {
		course.Contents = append(course.Contents, models.Content{
			ContentType: content.ContentType,
			YoutubeURL:  content.YoutubeURL,
			PDFURL:      content.PDFURL,
			ImageURL:    content.ImageURL,
		})
	}

	return course
}
