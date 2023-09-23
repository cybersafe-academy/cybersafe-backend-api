package httpmodels

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
	Title          string  `json:"title" valid:"required"`
	Description    string  `json:"description" valid:"required"`
	ContentInHours float64 `json:"contentInHours" valid:"required"`
	ThumbnailURL   string  `json:"thumbnailURL" valid:"required"`
	Level          string  `json:"level" valid:"required"`
	ContentURL     string  `json:"contentURL" valid:"required"`
}

type CourseResponse struct {
	CourseFields

	ID        uuid.UUID      `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
	AvgRating float64        `json:"avgRating"`

	Category CourseCategoryResponse `json:"category,omitempty"`

	Questions []QuestionResponse `json:"questions"`
}

type ResponseContent struct {
	CourseResponse

	ID        uuid.UUID      `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

type RequestContent struct {
	CourseFields

	CategoryID uuid.UUID `valid:"required"`

	Questions []QuestionRequest `json:"questions"`
}

type CourseExtraFields struct {
	ID        uuid.UUID `json:"id"`
	AvgRating float64   `json:"avgRating"`

	CourseFields
}

type CourseByCategory struct {
	CategoryName string              `json:"name"`
	Courses      []CourseExtraFields `json:"courses"`
}

type RawCoursesByCategory struct {
	CourseID uuid.UUID

	CourseTitle          string
	CourseThumbnailURL   string
	CourseContentURL     string
	AvgRating            float64
	CourseDescription    string
	CourseLevel          string
	CourseContentInHours float64

	CategoryName string
}

type CourseByCategoryResponse map[string][]map[string]any

func (re *RequestContent) Bind(_ *http.Request) error {
	if !govalidator.IsIn(re.Level, models.ValidCourseLevels...) {
		return errutil.ErrInvalidCourseLevel
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
		CategoryID:     re.CategoryID,
		ContentURL:     re.ContentURL,
	}

	for _, question := range re.Questions {
		questionModel := models.Question{
			Wording: question.Wording,
		}

		for _, answer := range question.Answers {
			questionModel.Answers = append(questionModel.Answers, models.Answer{
				Text:      answer.Text,
				IsCorrect: answer.IsCorrect,
			})
		}

		course.Questions = append(course.Questions, questionModel)
	}

	return course
}
