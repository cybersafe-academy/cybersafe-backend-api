package categories

import (
	"gorm.io/gorm"
)

func Config(conn *gorm.DB) CategoriesManager {
	return &CategoriesManagerDB{
		DBConnection: conn,
	}
}
