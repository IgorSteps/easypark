//go:build wireinject
// +build wireinject

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
	"github.com/IgorSteps/easypark/internal/drivers/scheduler"
	alertUsecases "github.com/IgorSteps/easypark/internal/usecases/alert"
	notificationUsecases "github.com/IgorSteps/easypark/internal/usecases/notification"
	parkingLotUsecases "github.com/IgorSteps/easypark/internal/usecases/parkinglot"
	parkingRequestUsecases "github.com/IgorSteps/easypark/internal/usecases/parkingrequest"
	parkingSpaceUsecases "github.com/IgorSteps/easypark/internal/usecases/parkingspace"
	userUsecases "github.com/IgorSteps/easypark/internal/usecases/user"
	"github.com/google/wire"
)

func SetupApp() (*App, error) {
	wire.Build(
		// config
		config.LoadConfig,
		wire.FieldsOf(new(*config.Config), "Database", "Auth", "Logging", "HTTP", "Scheduler", "Alert"),

		// logger
		logger.NewLoggerFromConfig,
		db.NewGormLogrusLoggerFromConfig,

		// repositories
		datastore.NewUserPostgresRepository,
		wire.Bind(new(repositories.UserRepository), new(*datastore.UserPostgresRepository)),
		datastore.NewParkingRequestPostgresRepository,
		wire.Bind(new(repositories.ParkingRequestRepository), new(*datastore.ParkingRequestPostgresRepository)),
		datastore.NewParkingSpacePostgresRepository,
		wire.Bind(new(repositories.ParkingSpaceRepository), new(*datastore.ParkingSpacePostgresRepository)),
		datastore.NewParkingParkingLotPostgresRepository,
		wire.Bind(new(repositories.ParkingLotRepository), new(*datastore.ParkingLotPostgresRepository)),
		datastore.NewNotificationPostgresRepository,
		wire.Bind(new(repositories.NotificationRepository), new(*datastore.NotificationPostgresRepository)),
		datastore.NewAlertPostgresRepository,
		wire.Bind(new(repositories.AlertRepository), new(*datastore.AlertPostgresRepository)),

		// db
		db.NewDatabaseFromConfig,
		db.NewGormWrapper,
		wire.Bind(new(datastore.Datastore), new(*db.GormWrapper)),

		// jwt
		auth.NewJWTTokenServiceFromConfig,
		wire.Bind(new(repositories.TokenRepository), new(*auth.JWTTokenService)),

		// usecases:
		// user
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
		// parking request
		parkingRequestUsecases.NewCreateParkingRequest,
		wire.Bind(new(usecasefacades.ParkingRequestCreator), new(*parkingRequestUsecases.CreateParkingRequest)),
		parkingRequestUsecases.NewUpdateParkingRequestStatus,
		wire.Bind(new(usecasefacades.ParkingRequestStatusUpdater), new(*parkingRequestUsecases.UpdateParkingRequestStatus)),
		parkingRequestUsecases.NewAssignParkingSpace,
		wire.Bind(new(usecasefacades.ParkingRequestSpaceAssigner), new(*parkingRequestUsecases.AssignParkingSpace)),
		parkingRequestUsecases.NewGetAllParkingRequests,
		wire.Bind(new(usecasefacades.ParkingRequestsAllGetter), new(*parkingRequestUsecases.GetAllParkingRequests)),
		parkingRequestUsecases.NewGetDriversParkingRequests,
		wire.Bind(new(usecasefacades.ParkingRequestDriversGetter), new(*parkingRequestUsecases.GetDriversParkingRequests)),
		// parking lot
		parkingLotUsecases.NewCreateParkingLot,
		wire.Bind(new(usecasefacades.ParkingLotCreator), new(*parkingLotUsecases.CreateParkingLot)),
		parkingLotUsecases.NewGetAllParkingLots,
		wire.Bind(new(usecasefacades.ParkingLotGetter), new(*parkingLotUsecases.GetAllParkingLots)),
		parkingLotUsecases.NewDeleteParkingLot,
		wire.Bind(new(usecasefacades.ParkingLotDeleter), new(*parkingLotUsecases.DeteleParkingLot)),
		parkingLotUsecases.NewGetSingleParkingLot,
		wire.Bind(new(usecasefacades.ParkingLotSingleGetter), new(*parkingLotUsecases.GetSingleParkingLot)),
		// parking space
		parkingSpaceUsecases.NewUpdateParkingSpaceStatus,
		wire.Bind(new(usecasefacades.ParkingSpaceStatusUpdater), new(*parkingSpaceUsecases.UpdateParkingSpaceStatus)),
		parkingSpaceUsecases.NewGetSingleParkingSpace,
		wire.Bind(new(usecasefacades.ParkingSpaceGetter), new(*parkingSpaceUsecases.GetSingleParkingSpace)),
		// alert
		alertUsecases.NewCreateAlert,
		wire.Bind(new(repositories.AlertCreator), new(*alertUsecases.CreateAlert)),
		alertUsecases.NewGetSingleAlert,
		wire.Bind(new(usecasefacades.AlertSingleGetter), new(*alertUsecases.GetSingleAlert)),
		alertUsecases.NewGetAllAlerts,
		wire.Bind(new(usecasefacades.AlertAllGetter), new(*alertUsecases.GetAllAlerts)),
		alertUsecases.NewCheckLateArrival,
		wire.Bind(new(usecasefacades.AlertLateArrivalChecker), new(*alertUsecases.CheckLateArrival)),
		// notification
		notificationUsecases.NewCreateNotification,
		wire.Bind(new(usecasefacades.NotificationCreator), new(*notificationUsecases.CreateNotification)),
		notificationUsecases.NewGetAllNotifications,
		wire.Bind(new(usecasefacades.NotificationGetter), new(*notificationUsecases.GetAllNotifications)),

		// facades
		usecasefacades.NewUserFacade,
		wire.Bind(new(handlers.UserFacade), new(*usecasefacades.UserFacade)),
		usecasefacades.NewParkingRequestFacade,
		wire.Bind(new(handlers.ParkingRequestFacade), new(*usecasefacades.ParkingRequestFacade)),
		usecasefacades.NewParkingLotFacade,
		wire.Bind(new(handlers.ParkingLotFacade), new(*usecasefacades.ParkingLotFacade)),
		usecasefacades.NewParkingSpaceFacade,
		wire.Bind(new(handlers.ParkingSpaceFacade), new(*usecasefacades.ParkingSpaceFacade)),
		usecasefacades.NewNotificationFacade,
		wire.Bind(new(handlers.NotificationFacade), new(*usecasefacades.NotificationFacade)),
		usecasefacades.NewAlertFacade,
		wire.Bind(new(handlers.AlertFacade), new(*usecasefacades.AlertFacade)),
		handlers.NewFacade,

		// scheduler
		scheduler.NewSchedulerFromConfig,

		// rest handlers and middleware
		middleware.NewMiddleware,
		wire.Bind(new(routes.Middleware), new(*middleware.Middleware)),
		handlers.NewHandlerFactory,
		wire.Bind(new(routes.HandlerFactory), new(*handlers.HandlerFactory)),

		// rest server
		routes.NewRouter,
		httpserver.NewServerFromConfig,

		// app
		NewApp,
	)

	return &App{}, nil
}
