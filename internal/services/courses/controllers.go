package courses

import (
	"gorm.io/gorm"
)

type CoursesManagerDB struct {
	DBConnection *gorm.DB
}
