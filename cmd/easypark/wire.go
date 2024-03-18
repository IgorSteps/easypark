//go:build wireinject
// +build wireinject

/***********************************************************************************
If you want to edit/activate InteliSense in this file:
 1) remove the build constraints above,
 2) edit file,
 3) put constraints back in
***********************************************************************************/

package main

import (
	"github.com/IgorSteps/easypark/internal/adapters/datastore"
	"github.com/IgorSteps/easypark/internal/adapters/rest/handlers"
	"github.com/IgorSteps/easypark/internal/adapters/rest/middleware"
	"github.com/IgorSteps/easypark/internal/adapters/rest/routes"
	"github.com/IgorSteps/easypark/internal/adapters/usecasefacades"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/IgorSteps/easypark/internal/drivers/auth"
	"github.com/IgorSteps/easypark/internal/drivers/config"
	"github.com/IgorSteps/easypark/internal/drivers/db"
	"github.com/IgorSteps/easypark/internal/drivers/httpserver"
	"github.com/IgorSteps/easypark/internal/drivers/logger"
	"github.com/IgorSteps/easypark/internal/usecases"
	"github.com/google/wire"
)

func BuildDIForApp() (*App, error) {
	wire.Build(
		// config
		config.LoadConfig,
		wire.FieldsOf(new(*config.Config), "Database", "Auth", "Logging", "HTTP"),

		// logger
		logger.NewLoggerFromConfig,
		db.NewGormLogrusLoggerFromConfig,

		// repositories
		datastore.NewUserPostgresRepository,
		wire.Bind(new(repositories.UserRepository), new(*datastore.UserPostgresRepository)),

		// db
		db.NewDatabaseFromConfig,
		db.NewGormWrapper,
		wire.Bind(new(datastore.Datastore), new(*db.GormWrapper)),

		// jwt
		auth.NewJWTTokenServiceFromConfig,
		wire.Bind(new(repositories.TokenRepository), new(*auth.JWTTokenService)),

		// usecase
		usecases.NewRegisterDriver,
		wire.Bind(new(usecasefacades.DriverCreator), new(*usecases.RegisterDriver)),
		usecases.NewAuthenticateUser,
		wire.Bind(new(usecasefacades.UserAuthenticator), new(*usecases.AuthenticateUser)),
		usecases.NewGetDrivers,
		wire.Bind(new(usecasefacades.DriversGetter), new(*usecases.GetDrivers)),

		// facades
		usecasefacades.NewUserFacade,
		wire.Bind(new(handlers.UserFacade), new(*usecasefacades.UserFacade)),

		// rest handlers
		middleware.NewAuthMiddleware,
		wire.Bind(new(routes.RequestAuthoriser), new(*middleware.AuthMiddleware)),
		handlers.NewHandlerFactory,
		wire.Bind(new(routes.HandlerFactory), new(*handlers.HandlerFactory)),

		// rest server
		routes.NewRouter,
		httpserver.NewServerFromConfig,

		// service
		NewApp,
	)

	return &App{}, nil
}
