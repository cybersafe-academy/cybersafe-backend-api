package models

import "github.com/google/uuid"

type Enrollment struct {
	Shared

	CourseID uuid.UUID
	Course   Course `gorm:"foreignKey:CourseID"`

	UserID uuid.UUID
	User   User `gorm:"foreignKey:UserID"`
}
