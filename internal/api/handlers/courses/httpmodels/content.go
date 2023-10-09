package httpmodels

import "github.com/google/uuid"

type ContentRequest struct {
	ContentFields
}

type ContentResponse struct {
	ContentFields

	ID uuid.UUID `json:"id" valid:"uuid, required"`
}

type ContentFields struct {
	Title string `json:"title" valid:"type(string), required"`
	URL   string `json:"URL" valid:"type(string)"`
}
