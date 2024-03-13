//go:build wireinject
// +build wireinject

/***********************************************************************************
If you want to edit/activate Intelisense in this file:
 1) remove the build constraints,
 2) edit file,
 3) put constraints back in
***********************************************************************************/

package main

import (
	"github.com/IgorSteps/easypark/internal/adapters/datastore"
	"github.com/IgorSteps/easypark/internal/adapters/rest/handlers"
	"github.com/IgorSteps/easypark/internal/adapters/rest/routes"
	"github.com/IgorSteps/easypark/internal/adapters/usecasefacades"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/IgorSteps/easypark/internal/drivers/db"
	"github.com/IgorSteps/easypark/internal/drivers/httpserver"
	"github.com/IgorSteps/easypark/internal/usecases"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

func BuildDIForApp() (*App, error) {
	wire.Build(
		// logger
		logrus.New,

		// repositories
		datastore.NewUserPostgresRepository,
		wire.Bind(new(repositories.UserRepository), new(*datastore.UserPostgresRepository)),

		// db
		db.NewDatabaseFromConfig,
		db.NewGormWrapper,
		wire.Bind(new(datastore.Datastore), new(*db.GormWrapper)),

		// usecase
		usecases.NewRegisterUser,
		wire.Bind(new(usecasefacades.UserCreator), new(*usecases.RegisterUser)),

		// facades
		usecasefacades.NewUserFacade,
		wire.Bind(new(handlers.UserFacade), new(*usecasefacades.UserFacade)),

		// rest handlers
		handlers.NewHandlerFactory,
		wire.Bind(new(routes.HandlerFactory), new(*handlers.HandlerFactory)),

		// rest server
		routes.NewRouter,
		httpserver.NewServer,

		// service
		NewApp,
	)

	return &App{}, nil
}
