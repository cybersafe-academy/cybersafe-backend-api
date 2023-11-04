package httpmodels

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type AnswerRequest struct {
	IsCorrect bool `json:"isCorrect"`
	AnswerFields
}

type AnswerResponse struct {
	AnswerFields
	ID uuid.UUID `json:"id" valid:"uuid, required"`
}

type AnswerFields struct {
	Text     string `json:"text"`
	TextPtBr string `json:"text-pt-br"`
}

type AddAnswerRequest struct {
	QuestionID uuid.UUID `json:"questionID" valid:"required"`
	AnswerID   uuid.UUID `json:"answerID" valid:"required"`
}

type AddAnswersBatchRequest struct {
	Answers []AddAnswerRequest `json:"answers" valid:"required"`
}

func (aar *AddAnswerRequest) Bind(_ *http.Request) error {

	_, err := govalidator.ValidateStruct(*aar)
	if err != nil {
		return err
	}

	return err
}

func (aabr *AddAnswersBatchRequest) Bind(_ *http.Request) error {

	_, err := govalidator.ValidateStruct(*aabr)
	if err != nil {
		return err
	}

	return err
}
