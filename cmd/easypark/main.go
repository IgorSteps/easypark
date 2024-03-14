package main

import (
	"log"

	"github.com/sirupsen/logrus"
)

const HTTPServerPort = "localhost:8080"

func main() {
	app, err := BuildDIForApp()
	if err != nil {
		log.Fatalf("failed to build DI for easpark app: %v", err)
	}

	app.logger.Info("starting Easypark")

	app.logger.Level = logrus.DebugLevel
	//app.logger.Formatter = new(logrus.JSONFormatter) // if we want JSON looking logs

	// This is blocking thread, nothing will run after this.
	app.logger.WithField("address", HTTPServerPort).Info("starting http server")
	err = app.server.Run(HTTPServerPort)
	if err != nil {
		log.Fatalf("failed to start REST server: %v", err)
	}

	app.logger.Info("shutting down Easypark")
}
