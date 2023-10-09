package db

import "cybersafe-backend-api/internal/models"

func AutoMigrateDB() error {
	db, connErr := GetDatabaseConnection()
	if connErr != nil {
		return connErr
	}

	modelsSlice := []any{
		&models.User{},
		&models.Course{},
		&models.Company{},
		&models.Review{},
		&models.Enrollment{},
		&models.Question{},
		&models.Answer{},
		&models.UserAnswer{},
		&models.Category{},
		&models.Comment{},
		&models.Like{},
	}

	err := db.AutoMigrate(modelsSlice...)

	return err
}
