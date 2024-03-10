package main

import (
	"fmt"
	"log"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	drivers "github.com/IgorSteps/easypark/internal/drivers/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dbConfig, err := drivers.LoadDatabaseConfig()
	if err != nil {
		// NOTE:
		// Best practice is we must never fatal or panic, and instead load default values,
		// but it is a univercity project do I don't mind.
		log.Fatalf("Failed to load database config: %v", err)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=UTC",
		dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.Port, dbConfig.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Automigrate here:
	err = db.AutoMigrate(&entities.User{})
	if err != nil {
		log.Fatalf("Failed to automigrate: %v", err)
	}
}
