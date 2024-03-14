package main

import (
	"log"
)

// TODO: Move to config
const HTTPServerPort = "localhost:8080"

func main() {
	app, err := BuildDIForApp()
	if err != nil {
		log.Fatalf("failed to build DI for easpark app: %v", err)
	}

	app.logger.Info("starting Easypark")

	// This is blocking thread, nothing will run after this.
	app.logger.WithField("address", HTTPServerPort).Info("starting http server")
	err = app.server.Run(HTTPServerPort)
	if err != nil {
		log.Fatalf("failed to start REST server: %v", err)
	}

	app.logger.Info("shutting down Easypark")
}
