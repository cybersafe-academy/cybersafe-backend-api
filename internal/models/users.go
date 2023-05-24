package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	DefaultUserRole string = "default"
	AdminUserRole   string = "admin"
	MasterUserRole  string = "master"
)

var (
	ValidUserRoles []string = []string{
		DefaultUserRole,
		AdminUserRole,
		MasterUserRole,
	}
)

type User struct {
	Shared

	Name      string
	Role      string `gorm:"default:'default'"`
	Email     string `gorm:"unique"`
	BirthDate time.Time
	CPF       string `gorm:"unique"`
	Password  string
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	return nil
}

func RoleToHierarchyNumber(role string) int {
	switch role {
	case DefaultUserRole:
		return 0
	case AdminUserRole:
		return 1
	case MasterUserRole:
		return 2
	default:
		return -1
	}
}
