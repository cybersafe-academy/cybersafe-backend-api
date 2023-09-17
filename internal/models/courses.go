package models

import "github.com/google/uuid"

const (
	BeginnerLevel     string = "beginner"
	IntermediateLevel string = "intermediate"
	AdvancedLevel     string = "advanced"
)

var (
	ValidCourseLevels []string = []string{
		BeginnerLevel,
		IntermediateLevel,
		AdvancedLevel,
	}
)

type Course struct {
	Shared

	Title          string
	Description    string
	ContentInHours float64
	ThumbnailURL   string
	Level          string

	CategoryID uuid.UUID `gorm:"default:null"`
	Category   Category  `gorm:"foreignKey:CategoryID"`

	Contents    []Content    `gorm:"foreignKey:CourseID"`
	Questions   []Question   `gorm:"foreignKey:CourseID"`
	Reviews     []Review     `gorm:"foreignKey:CourseID"`
	Enrollments []Enrollment `gorm:"foreignKey:CourseID"`
}

type CourseExtraFields struct {
	Course

	AvgRating float64
}

type CourseByCategoryFields struct {
	Course

	AvgRating float64
}

type Category struct {
	Shared

	Name string `gorm:"uniqueIndex:idx_course_name"`

	Course []Course
}

type Content struct {
	Shared

	Title       string
	ContentType string
	URL         string

	CourseID uuid.UUID
	Course   Course `gorm:"foreignKey:CourseID"`
}

type Review struct {
	Shared

	Comment string
	Rating  int

	UserID uuid.UUID `gorm:"uniqueIndex:idx_course_user"`
	User   User      `gorm:"foreignKey:UserID"`

	CourseID uuid.UUID `gorm:"uniqueIndex:idx_course_user"`
	Course   Course    `gorm:"foreignKey:CourseID"`
}

type Question struct {
	Shared

	Wording string

	CourseID uuid.UUID
	Course   Course `gorm:"foreignKey:CourseID"`

	Answers []Answer `gorm:"foreignKey:QuestionID"`
}

type Answer struct {
	Shared

	QuestionID uuid.UUID
	Question   Question `gorm:"foreignKey:QuestionID"`

	Text      string
	IsCorrect bool
}
