package models

import "github.com/google/uuid"

type Like struct {
	Shared

	CommentID uuid.UUID `gorm:"uniqueIndex:idx_comment_user"`
	Comment   Comment   `gorm:"foreignKey:CommentID"`

	UserID uuid.UUID `gorm:"uniqueIndex:idx_comment_user"`
	User   User      `gorm:"foreignKey:UserID"`
}
