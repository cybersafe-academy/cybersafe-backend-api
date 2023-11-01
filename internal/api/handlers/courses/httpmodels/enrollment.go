package httpmodels

type EnrollmentFields struct {
	Status string `json:"status"`
}

type EnrollmentResponse struct {
	EnrollmentFields
	HitsPercentage float64 `json:"hitsPercentage"`
}

type EnrollmentRequest struct {
	EnrollmentFields
}
