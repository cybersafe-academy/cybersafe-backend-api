package users

import "gorm.io/gorm"

func Config(conn *gorm.DB) UsersManager {
	return &UsersManagerDB{
		DBConnection: conn,
	}
}
