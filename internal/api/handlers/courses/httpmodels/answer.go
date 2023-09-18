package httpmodels

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type AnswerRequest struct {
	AnswerFields
}

type AnswerResponse struct {
	AnswerFields
	ID uuid.UUID `json:"id" valid:"uuid, required"`
}

type AnswerFields struct {
	IsCorrect bool   `json:"-"`
	Text      string `json:"text"`
}

type AddAnswerRequest struct {
	QuestionID uuid.UUID `json:"questionID" valid:"required"`
	AnswerID   uuid.UUID `json:"answerID" valid:"required"`
}

func (aar *AddAnswerRequest) Bind(_ *http.Request) error {

	_, err := govalidator.ValidateStruct(*aar)
	if err != nil {
		return err
	}

	return err
}
