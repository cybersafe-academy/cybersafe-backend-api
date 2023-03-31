package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	Shared struct {
		gorm.Model
		ID uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	}
)
