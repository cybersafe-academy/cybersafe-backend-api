package models

import "github.com/google/uuid"

const (
	BeginnerLevel     string = "beginner"
	IntermediateLevel string = "intermediate"
	AdvancedLevel     string = "advanced"
)

const (
	ContentTypeYoutube string = "youtube"
	ContentTypePDF     string = "pdf"
	ContentTypeImage   string = "image"
)

var (
	ValidContentTypes []string = []string{
		ContentTypeYoutube,
		ContentTypePDF,
		ContentTypeImage,
	}

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

	Contents []Content `gorm:"foreignKey:CourseID"`
}

type Content struct {
	Shared

	Title       string
	ContentType string
	URL         string

	CourseID uuid.UUID
	Course   Course `gorm:"foreignKey:CourseID"`
}
