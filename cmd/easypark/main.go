package main

import (
	"log"
)

const httpServerPort = ":8081"

func main() {
	app, err := BuildDIForApp()
	if err != nil {
		log.Fatalf("failed to build DI for easpark app: %v", err)
	}

	log.Printf("Starting Easypark")

	//app.logger.Formatter = new(logrus.JSONFormatter)

	// This is blocking thread, nothing will run after this.
	err = app.server.Run(httpServerPort)
	if err != nil {
		log.Fatalf("failed to start REST server: %v", err)
	}

	log.Print("Shutting down Easypark")
}