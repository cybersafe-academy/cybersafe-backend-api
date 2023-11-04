package models

import "github.com/google/uuid"

const (
	ApprovedStatus   string = "approved"
	FailedStatus     string = "failed"
	InProgressStatus string = "in_progress"
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

	CourseID uuid.UUID `gorm:"uniqueIndex:idx_user_course_enroll"`
	Course   Course    `gorm:"foreignKey:CourseID"`

	UserID uuid.UUID `gorm:"uniqueIndex:idx_user_course_enroll"`
	User   User      `gorm:"foreignKey:UserID"`

	Status       string
	QuizProgress float64
}
