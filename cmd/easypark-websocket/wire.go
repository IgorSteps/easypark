//go:build wireinject
// +build wireinject

package main

import (
	"github.com/IgorSteps/easypark/internal/adapters/websocket/routes"
	"github.com/IgorSteps/easypark/internal/drivers/config"
	"github.com/IgorSteps/easypark/internal/drivers/logger"
	"github.com/IgorSteps/easypark/internal/drivers/websocketserver"
	"github.com/google/wire"
)

func SetupApp() (*App, error) {
	wire.Build(
		// config
		config.LoadConfig,
		wire.FieldsOf(new(*config.Config), "Logging", "HTTP"),

		// logger
		logger.NewLoggerFromConfig,
		//db.NewGormLogrusLoggerFromConfig,

		// websocket server
		routes.NewRouter,
		websocketserver.NewServerFromConfig,

		// app
		NewApp,
	)

	return &App{}, nil
}
