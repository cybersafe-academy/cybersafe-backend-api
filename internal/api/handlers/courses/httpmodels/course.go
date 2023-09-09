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
	AvgRating float64        `json:"avgRating"`

	Contents  []ContentResponse  `json:"contents"`
	Questions []QuestionResponse `json:"questions"`
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

	Contents  []ContentRequest  `json:"contents"`
	Questions []QuestionRequest `json:"questions"`
}

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
	}

	for _, content := range re.Contents {
		course.Contents = append(course.Contents, models.Content{
			Title: content.Title,
			URL:   content.URL,
		})
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
