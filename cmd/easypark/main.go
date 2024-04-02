package main

import (
	"log"
)

func main() {
	app, err := SetupApp()
	if err != nil {
		log.Fatalf("failed to build DI for easpark app: %v", err)
	}

	app.logger.Info("starting Easypark")

	// This is blocking thread, nothing will run after this.
	app.logger.WithField("address", app.server.Address).Info("starting http server")
	err = app.server.Run()
	if err != nil {
		log.Fatalf("failed to start REST server: %v", err)
	}

	app.logger.Info("shutting down Easypark")
}
