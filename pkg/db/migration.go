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
		&models.Content{},
	}

	err := db.AutoMigrate(modelsSlice...)

	return err
}
