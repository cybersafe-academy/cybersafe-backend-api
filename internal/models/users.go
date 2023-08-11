package models

import (
	"time"
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
	CPF       string `gorm:"unique;default:null"`
	Password  string
	Enabled   bool `gorm:"default:false;not null"`
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
