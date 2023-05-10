package users

import "gorm.io/gorm"

func Config(conn *gorm.DB) Users {
	return &UserDB{
		Conn: conn,
	}
}
