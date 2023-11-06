package analytics

import (
	"gorm.io/gorm"
)

func Config(conn *gorm.DB) AnalyticsManager {
	return &AnalyticsManagerDB{
		DBConnection: conn,
	}
}
