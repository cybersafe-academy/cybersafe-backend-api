package models

import (
	"time"

	"github.com/google/uuid"
)

const (
	DefaultUserRole string = "default"
	AdminUserRole   string = "admin"
	MasterUserRole  string = "master"
)

var (
	ValidUserRoles []string = []string{
		DefaultUserRole,
		AdminUserRole,
		MasterUserRole,
	}
)

type User struct {
	Shared

	Name              string
	Role              string `gorm:"default:'default'"`
	Email             string `gorm:"unique"`
	BirthDate         time.Time
	CPF               string `gorm:"unique;default:null"`
	ProfilePictureURL string
	Password          string
	MBTIType          string
	Enabled           bool `gorm:"default:false;not null"`

	Enrollments []Enrollment `gorm:"foreignKey:UserID"`
}

func (u *User) SetImageURL(url string) {
	u.ProfilePictureURL = url
}

type UserAnswer struct {
	Shared

	UserID uuid.UUID `gorm:"uniqueIndex:idx_question_answer_user"`
	User   User      `gorm:"foreignKey:UserID"`

	QuestionID uuid.UUID `gorm:"uniqueIndex:idx_question_answer_user"`
	Question   Question  `gorm:"foreignKey:QuestionID"`

	AnswerID uuid.UUID `gorm:"uniqueIndex:idx_question_answer_user"`
	Answer   Answer    `gorm:"foreignKey:AnswerID"`
}

func RoleToHierarchyNumber(role string) int {
	switch role {
	case DefaultUserRole:
		return 0
	case AdminUserRole:
		return 1
	case MasterUserRole:
		return 2
	default:
		return -1
	}
}
