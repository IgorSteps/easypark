//go:build wireinject
// +build wireinject

package main

import (
	"github.com/IgorSteps/easypark/internal/adapters/websocket/handlers"
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

		// repos
		// datastore.NewUserPostgresRepository,
		// wire.Bind(new(repositories.UserRepository), new(*datastore.UserPostgresRepository)),

		// usecase
		//usecases.NewCreateMessage,

		handlers.NewHub,
		// websocket server
		routes.NewRouter,
		websocketserver.NewServerFromConfig,

		// app
		NewApp,
	)

	return &App{}, nil
}
