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
	parkingRequestUsecases "github.com/IgorSteps/easypark/internal/usecases/parkingrequest"
	userUsecases "github.com/IgorSteps/easypark/internal/usecases/user"
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
		datastore.NewParkingRequestPostgresRepository,
		wire.Bind(new(repositories.ParkingRequestRepository), new(*datastore.ParkingRequestPostgresRepository)),

		// db
		db.NewDatabaseFromConfig,
		db.NewGormWrapper,
		wire.Bind(new(datastore.Datastore), new(*db.GormWrapper)),

		// jwt
		auth.NewJWTTokenServiceFromConfig,
		wire.Bind(new(repositories.TokenRepository), new(*auth.JWTTokenService)),

		// usecase
		userUsecases.NewRegisterDriver,
		wire.Bind(new(usecasefacades.DriverCreator), new(*userUsecases.RegisterDriver)),
		userUsecases.NewAuthenticateUser,
		wire.Bind(new(usecasefacades.UserAuthenticator), new(*userUsecases.AuthenticateUser)),
		userUsecases.NewGetDrivers,
		wire.Bind(new(usecasefacades.DriversGetter), new(*userUsecases.GetDrivers)),
		userUsecases.NewBanDriver,
		wire.Bind(new(usecasefacades.DriverBanner), new(*userUsecases.BanDriver)),
		userUsecases.NewCheckDriverStatus,
		wire.Bind(new(middleware.StatusChecker), new(*userUsecases.CheckDriverStatus)),
		parkingRequestUsecases.NewCreateParkingRequest,
		wire.Bind(new(usecasefacades.ParkingRequestCreator), new(*parkingRequestUsecases.CreateParkingRequest)),

		// facades
		usecasefacades.NewUserFacade,
		usecasefacades.NewParkingRequestFacade,
		wire.Bind(new(handlers.UserFacade), new(*usecasefacades.UserFacade)),
		wire.Bind(new(handlers.ParkingRequestFacade), new(*usecasefacades.ParkingRequestFacade)),
		handlers.NewFacade,

		// rest handlers and middleware
		middleware.NewMiddleware,
		wire.Bind(new(routes.Middleware), new(*middleware.Middleware)),
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
