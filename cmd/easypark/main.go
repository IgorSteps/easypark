package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	app, err := SetupApp()
	if err != nil {
		log.Fatalf("failed to setup Easpark app: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	app.logger.Info("starting Easypark")

	// Start the Scheduler.
	app.scheduler.Start()
	defer app.scheduler.Stop()

	// Start t heREST server.
	go func() {
		if err := app.server.Run(); err != nil {
			log.Fatalf("failed to start REST server: %v", err)
		}
	}()

	<-ctx.Done()
	app.logger.Info("shutting down Easypark")
}
