package httpmodels

import (
	"github.com/google/uuid"
)

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
	Wording     string `json:"wording"`
	WordingPtBr string `json:"wording-pt-br"`
}
