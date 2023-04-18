package models

import "github.com/google/uuid"

type LevelString string
type ContentTypeString string

const (
	BeginnerLevel     LevelString = "beginner"
	IntermediateLevel LevelString = "intermediate"
	AdvancedLevel     LevelString = "advanced"
)

const (
	ContentTypeYoutube ContentTypeString = "youtube"
	ContentTypePDF     ContentTypeString = "pdf"
	ContentTypeImage   ContentTypeString = "image"
)

type Course struct {
	Shared

	Name           string
	Description    string
	ContentInHours float64
	ThumbnailURL   string
	Level          string

	Contents []Content
}

type Content struct {
	Shared

	ContentType string
	YoutubeURL  string
	PDFURL      string
	ImageURL    string

	CourseID uuid.UUID
	Course   Course
}
