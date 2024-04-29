//go:build wireinject
// +build wireinject

package main

import (
	"github.com/IgorSteps/easypark/internal/adapters/datastore"
	"github.com/IgorSteps/easypark/internal/adapters/usecasefacades"
	"github.com/IgorSteps/easypark/internal/adapters/websocket/client"
	"github.com/IgorSteps/easypark/internal/adapters/websocket/routes"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/IgorSteps/easypark/internal/drivers/config"
	"github.com/IgorSteps/easypark/internal/drivers/db"
	"github.com/IgorSteps/easypark/internal/drivers/logger"
	"github.com/IgorSteps/easypark/internal/drivers/websocketserver"
	usecases "github.com/IgorSteps/easypark/internal/usecases/message"
	"github.com/google/wire"
)

func SetupApp() (*App, error) {
	wire.Build(
		// config
		config.LoadConfig,
		wire.FieldsOf(new(*config.Config), "Database", "Logging", "HTTP"),

		// logger
		logger.NewLoggerFromConfig,
		db.NewGormLogrusLoggerFromConfig,

		// db
		db.NewDatabaseFromConfig,
		db.NewGormWrapper,
		wire.Bind(new(datastore.Datastore), new(*db.GormWrapper)),

		// repos
		datastore.NewUserPostgresRepository,
		wire.Bind(new(repositories.UserRepository), new(*datastore.UserPostgresRepository)),
		datastore.NewMessagePostgresRepository,
		wire.Bind(new(repositories.MessageRepository), new(*datastore.MessagePostgresRepository)),

		// usecase
		usecases.NewQueueMessage,
		wire.Bind(new(usecasefacades.MessageEnqueuer), new(*usecases.QueueMessage)),
		usecases.NewDequeueMessages,
		wire.Bind(new(usecasefacades.MessageDequeuer), new(*usecases.DequeueMessages)),

		// facades
		usecasefacades.NewMessageFacade,
		wire.Bind(new(client.MessageFacade), new(*usecasefacades.MessageFacade)),

		// hub
		client.NewHub,

		// websocket server
		routes.NewRouter,
		websocketserver.NewServerFromConfig,

		// app
		NewApp,
	)

	return &App{}, nil
}
