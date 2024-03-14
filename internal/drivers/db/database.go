package db

import (
	"fmt"
	"log"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/drivers/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDatabaseFromConfig creates our databse from config.
func NewDatabaseFromConfig(config config.DatabaseConfig, logger *GormLogrusLogger) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=UTC",
		config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}

	db.AutoMigrate(&entities.User{})

	return db, nil
}
