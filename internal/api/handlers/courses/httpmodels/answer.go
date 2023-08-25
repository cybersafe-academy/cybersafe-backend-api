package httpmodels

import "github.com/google/uuid"

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
