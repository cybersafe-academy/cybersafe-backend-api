package course

import (
	"cybersafe-backend-api/internal/models"
	"cybersafe-backend-api/pkg/errutil"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CourseFields struct {
	Title          string  `json:"title" valid:"type(string), required"`
	Description    string  `json:"description" valid:"type(string), required"`
	ContentInHours float64 `json:"contentInHours" valid:"type(float), required"`
	ThumbnailURL   string  `json:"thumbnailURL" valid:"type(string), required"`
	Level          string  `json:"level" valid:"type(string), required"`

	Contents []ContentFields `json:"contents"`
}

type ContentFields struct {
	ID uuid.UUID `json:"id" valid:"uuid, required"`

	Title       string `json:"title" valid:"type(string), required"`
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

	if !govalidator.IsIn(re.Level, models.ValidCourseLevels...) {
		return errutil.ErrInvalidCourseLevel
	}

	for _, content := range re.Contents {
		if !govalidator.IsIn(content.ContentType, models.ValidContentTypes...) {
			return errutil.ErrInvalidContentType
		}
	}

	_, err := govalidator.ValidateStruct(*re)
	if err != nil {
		return err
	}

	return err
}

func (re *RequestContent) ToEntity() *models.Course {
	course := &models.Course{
		Title:          re.Title,
		Description:    re.Description,
		ContentInHours: re.ContentInHours,
		ThumbnailURL:   re.ThumbnailURL,
		Level:          re.Level,
	}

	for _, content := range re.Contents {
		course.Contents = append(course.Contents, models.Content{
			Title:       re.Title,
			ContentType: content.ContentType,
			YoutubeURL:  content.YoutubeURL,
			PDFURL:      content.PDFURL,
			ImageURL:    content.ImageURL,
		})
	}

	return course
}
