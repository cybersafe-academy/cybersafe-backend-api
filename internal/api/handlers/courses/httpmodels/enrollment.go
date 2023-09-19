package httpmodels

type EnrollmentFields struct {
	Status       string
	QuizProgress float64
}

type EnrollmentResponse struct {
	EnrollmentFields
}

type EnrollmentRequest struct {
	EnrollmentFields
}
