package main

import (
	"log"

	"github.com/IgorSteps/easypark/internal/drivers/httpserver"
)

func main() {
	log.Printf("Starting Easypark")

	client, err := httpserver.NewClientFromConfig()
	if err != nil {
		log.Fatalf("Failed to create HTTP server client: %v", err)
	}

	if err := client.Run(); err != nil {
		log.Fatalf("Failed to run the HTTP server: %v", err)
	}
}
