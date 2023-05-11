package courses

import (
	"gorm.io/gorm"
)

func Config(conn *gorm.DB) CoursesManager {
	return &CoursesManagerDB{
		DBConnection: conn,
	}
}
