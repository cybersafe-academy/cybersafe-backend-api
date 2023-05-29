package companies

import (
	"gorm.io/gorm"
)

func Config(conn *gorm.DB) CompaniesManager {
	return &CompaniesManagerDB{
		DBConnection: conn,
	}
}
