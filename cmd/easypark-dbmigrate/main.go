package main

import (
	"log"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/drivers/db"
)

func main() {
	db, err := db.NewDatabaseFromConfig()
	if err != nil {
		log.Fatalf("failed to created dabasae from config: %v", err)
	}

	// Automigrate here:
	err = db.AutoMigrate(&entities.User{})
	if err != nil {
		log.Fatalf("failed to automigrate: %v", err)
	}

	log.Printf("DB is all setup and ready to go")
}
