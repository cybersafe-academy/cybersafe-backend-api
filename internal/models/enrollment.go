package models

import "github.com/google/uuid"

const (
	ApprovedStatus   string = "beginner"
	FailedStatus     string = "intermediate"
	InProgressStatus string = "advanced"
)

var (
	ValidCourseLeves []string = []string{
		ApprovedStatus,
		FailedStatus,
		InProgressStatus,
	}
)

type Enrollment struct {
	Shared

	CourseID uuid.UUID
	Course   Course `gorm:"foreignKey:CourseID"`

	UserID uuid.UUID
	User   User `gorm:"foreignKey:UserID"`

	Status       string
	QuizProgress float64
}
