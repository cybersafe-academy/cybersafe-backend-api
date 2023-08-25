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
	AvgRating float64        `json:"avgRating"`

	Contents  []ContentResponse  `json:"contents"`
	Questions []QuestionResponse `json:"questions"`
}

type ContentFields struct {
	Title       string `json:"title" valid:"type(string), required"`
	ContentType string `json:"contentType" valid:"type(string), required"`
	URL         string `json:"URL" valid:"type(string)"`
}

type ContentRequest struct {
	ContentFields
}

type ContentResponse struct {
	ContentFields

	ID uuid.UUID `json:"id" valid:"uuid, required"`
}

type QuestionRequest struct {
	QuestionFields
	Answers []AnswerRequest `json:"answers" valid:"required"`
}

type QuestionResponse struct {
	QuestionFields

	Answers []AnswerResponse `json:"answers" valid:"required"`
	ID      uuid.UUID        `json:"id" valid:"uuid, required"`
}

type QuestionFields struct {
	Wording string `json:"wording"`
}

type AnswerRequest struct {
	AnswerFields
}

type AnswerResponse struct {
	AnswerFields
	ID uuid.UUID `json:"id" valid:"uuid, required"`
}

type AnswerFields struct {
	IsCorrect bool   `json:"isCorrect"`
	Text      string `json:"text"`
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

	Contents  []ContentRequest
	Questions []QuestionRequest
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
	Comment string `json:"comment" valid:"required"`
	Rating  int    `json:"rating" valid:"range(1|5), required"`

	CourseID uuid.UUID `json:"courseID" valid:"required"`
}

type ReviewResponse struct {
	ReviewFields

	ID        uuid.UUID      `json:"id" valid:"uuid, required"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`

	UserID uuid.UUID `json:"userID"`
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

func (rrc *ReviewRequestContent) ToEntityReview() *models.Review {
	review := &models.Review{
		Comment:  rrc.Comment,
		Rating:   rrc.Rating,
		CourseID: rrc.CourseID,
	}
	return review
}
