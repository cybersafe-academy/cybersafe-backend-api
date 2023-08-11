package courses

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
	ContentInHours float64 `json:"contentInHours" valid:"required"`
	ThumbnailURL   string  `json:"thumbnailURL" valid:"type(string), required"`
	Level          string  `json:"level" valid:"type(string), required"`
}

type CourseResponse struct {
	CourseFields

	ID        uuid.UUID      `json:"id" valid:"uuid, required"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`

	Contents []ContentResponse `json:"contents"`
}

type ContentRequest struct {
	ContentFields
}

type ContentResponse struct {
	ContentFields

	ID uuid.UUID `json:"id" valid:"uuid, required"`
}

type ContentFields struct {
	Title       string `json:"title" valid:"type(string), required"`
	ContentType string `json:"contentType" valid:"type(string), required"`
	URL         string `json:"URL" valid:"type(string)"`
}

type ResponseContent struct {
	CourseResponse

	ID        uuid.UUID      `json:"id" valid:"uuid, required"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

type RequestContent struct {
	CourseFields

	Contents []ContentRequest
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

type ReviewFields struct {
	Comment  string    `json:"comment"`
	Rating   int       `json:"rating"`
	UserID   uuid.UUID `json:"userID"`
	CourseID uuid.UUID `json:"courseID"`
}

type ReviewResponse struct {
	ReviewFields

	ID        uuid.UUID      `json:"id" valid:"uuid, required"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

type ReviewRequestContent struct {
	ReviewFields
}

func (rre *ReviewRequestContent) Bind(_ *http.Request) error {
	_, err := govalidator.ValidateStruct(*rre)
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
			Title:       content.Title,
			ContentType: content.ContentType,
			URL:         content.URL,
		})
	}

	return course
}

func (rrc *ReviewRequestContent) ToEntityReview() *models.Review {
	review := &models.Review{
		Comment:  rrc.Comment,
		Rating:   rrc.Rating,
		UserID:   rrc.UserID,
		CourseID: rrc.CourseID,
	}
	return review
}
