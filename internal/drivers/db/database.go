package db

import (
	"fmt"
	"os"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/drivers/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDatabaseFromConfig creates our database from config.
func NewDatabaseFromConfig(config config.DatabaseConfig, logger *GormLogrusLogger) (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	logger.Logrus.WithField("db host", host).Debug("got db host from os env")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=UTC",
		host, config.User, config.Password, config.DBName, config.Port, config.SSLMode,
	)

	// Init db session and configure GORM to use our logger.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		logger.Logrus.WithError(err).Error("failed to connect to the database")
		return nil, err
	}

	// Create/Migrate our db tables.
	db.AutoMigrate(
		&entities.User{},
		&entities.ParkingLot{},
		&entities.ParkingSpace{},
		&entities.ParkingRequest{},
		&entities.Notification{},
		&entities.Alert{},
	)

	return db, nil
}
