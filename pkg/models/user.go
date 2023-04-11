package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Shared
	Name     string
	Age      int
	CPF      string
	Password string
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	return nil
}
