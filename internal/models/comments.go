package models

import "github.com/google/uuid"

type Comment struct {
	Shared

	Text string `gorm:"unique"`

	CourseID uuid.UUID
	Course   Course `gorm:"foreignKey:CourseID"`

	UserID uuid.UUID
	User   User `gorm:"foreignKey:UserID"`

	Likes []Like `gorm:"foreignKey:CommentID"`
}
