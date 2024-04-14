package db

import (
	"fmt"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/drivers/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDatabaseFromConfig creates our database from config.
func NewDatabaseFromConfig(config config.DatabaseConfig, logger *GormLogrusLogger) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=UTC",
		config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode,
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
