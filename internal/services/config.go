package services

import (
	"cybersafe-backend-api/internal/services/companies"
	"cybersafe-backend-api/internal/services/courses"
	"cybersafe-backend-api/internal/services/users"

	"gorm.io/gorm"
)

func Config(conn *gorm.DB) Resources {
	return Resources{
		Users:     users.Config(conn),
		Courses:   courses.Config(conn),
		Companies: companies.Config(conn),
	}
}
