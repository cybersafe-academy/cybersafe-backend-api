package answers

import (
	"gorm.io/gorm"
)

func Config(conn *gorm.DB) AnswersManager {
	return &AnswersManagerDB{
		DBConnection: conn,
	}
}
