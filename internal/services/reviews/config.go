package reviews

import (
	"gorm.io/gorm"
)

func Config(conn *gorm.DB) ReviewsManager {
	return &ReviewsManagerDB{
		DBConnection: conn,
	}
}
