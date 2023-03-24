package db

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"cybersafe-backend-api/pkg/helpers"
	"cybersafe-backend-api/pkg/models"
	"cybersafe-backend-api/pkg/settings"
)

var dbConnection *gorm.DB

func CreateDBConnection(config settings.Settings) *gorm.DB {

	args := fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		config.String("database.user"),
		config.String("database.password"),
		config.String("database.host"),
		config.String("database.name"),
	)

	db, err := gorm.Open(postgres.Open(args), &gorm.Config{
		PrepareStmt: true,
		NowFunc:     helpers.DefaultTimeZone,
		Logger:      logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Info().Err(err).Msg("Error occurred while connecting with the database")
		os.Exit(-1)
	}

	return db
}

func GetDatabaseConnection() (*gorm.DB, error) {
	sqlDB, err := dbConnection.DB()
	if err != nil {
		return dbConnection, err
	}
	if err := sqlDB.Ping(); err != nil {
		return dbConnection, err
	}

	return dbConnection, nil
}

func AutoMigrateDB() error {
	db, connErr := GetDatabaseConnection()
	if connErr != nil {
		return connErr
	}

	err := db.AutoMigrate(&models.User{})

	return err
}
