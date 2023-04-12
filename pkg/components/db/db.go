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
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		config.String("database.host"),
		config.String("database.user"),
		config.String("database.password"),
		config.String("database.name"),
		config.String("database.port"),
	)

	db, err := gorm.Open(postgres.Open(args), &gorm.Config{
		PrepareStmt: true,
		NowFunc:     helpers.DefaultTimeZone,
		Logger:      logger.Default.LogMode(logger.Info),
	})

	dbConnection = db

	if err != nil {
		log.Info().Err(err).Msg("Error occurred while connecting with the database")
		os.Exit(-1)
	}

	return db
}

func GetDatabaseConnection() (*gorm.DB, error) {

	if dbConnection != nil {
		return dbConnection, nil
	}

	err := fmt.Errorf("database connection was not configured")

	log.Info().Err(err).Msg("Error occurred while connecting with the database")

	return nil, err
}

func MustGetDbConn() *gorm.DB {

	db, _ := GetDatabaseConnection()

	return db
}

func AutoMigrateDB() error {
	db, connErr := GetDatabaseConnection()
	if connErr != nil {
		return connErr
	}

	modelsSlice := []any{
		&models.User{},
	}

	err := db.AutoMigrate(modelsSlice...)

	return err
}
