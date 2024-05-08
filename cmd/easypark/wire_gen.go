// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/IgorSteps/easypark/internal/adapters/datastore"
	"github.com/IgorSteps/easypark/internal/adapters/rest/handlers"
	"github.com/IgorSteps/easypark/internal/adapters/rest/middleware"
	"github.com/IgorSteps/easypark/internal/adapters/rest/routes"
	"github.com/IgorSteps/easypark/internal/adapters/usecasefacades"
	"github.com/IgorSteps/easypark/internal/drivers/auth"
	"github.com/IgorSteps/easypark/internal/drivers/config"
	"github.com/IgorSteps/easypark/internal/drivers/db"
	"github.com/IgorSteps/easypark/internal/drivers/httpserver"
	"github.com/IgorSteps/easypark/internal/drivers/logger"
	"github.com/IgorSteps/easypark/internal/drivers/scheduler"
	usecases5 "github.com/IgorSteps/easypark/internal/usecases/alert"
	usecases6 "github.com/IgorSteps/easypark/internal/usecases/notification"
	usecases3 "github.com/IgorSteps/easypark/internal/usecases/parkinglot"
	usecases2 "github.com/IgorSteps/easypark/internal/usecases/parkingrequest"
	usecases4 "github.com/IgorSteps/easypark/internal/usecases/parkingspace"
	"github.com/IgorSteps/easypark/internal/usecases/user"
)

// Injectors from wire.go:

func SetupApp() (*App, error) {
	configConfig, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}
	loggingConfig := configConfig.Logging
	logrusLogger := logger.NewLoggerFromConfig(loggingConfig)
	databaseConfig := configConfig.Database
	gormLogrusLogger := db.NewGormLogrusLoggerFromConfig(loggingConfig, logrusLogger)
	gormDB, err := db.NewDatabaseFromConfig(databaseConfig, gormLogrusLogger)
	if err != nil {
		return nil, err
	}
	gormWrapper := db.NewGormWrapper(gormDB)
	userPostgresRepository := datastore.NewUserPostgresRepository(gormWrapper, logrusLogger)
	registerDriver := usecases.NewRegisterDriver(logrusLogger, userPostgresRepository)
	authConfig := configConfig.Auth
	jwtTokenService, err := auth.NewJWTTokenServiceFromConfig(authConfig)
	if err != nil {
		return nil, err
	}
	authenticateUser := usecases.NewAuthenticateUser(logrusLogger, userPostgresRepository, jwtTokenService)
	getDrivers := usecases.NewGetDrivers(logrusLogger, userPostgresRepository)
	banDriver := usecases.NewBanDriver(logrusLogger, userPostgresRepository)
	userFacade := usecasefacades.NewUserFacade(registerDriver, authenticateUser, getDrivers, banDriver)
	parkingRequestPostgresRepository := datastore.NewParkingRequestPostgresRepository(gormWrapper, logrusLogger)
	createParkingRequest := usecases2.NewCreateParkingRequest(logrusLogger, parkingRequestPostgresRepository)
	updateParkingRequestStatus := usecases2.NewUpdateParkingRequestStatus(logrusLogger, parkingRequestPostgresRepository)
	parkingSpacePostgresRepository := datastore.NewParkingSpacePostgresRepository(logrusLogger, gormWrapper)
	assignParkingSpace := usecases2.NewAssignParkingSpace(logrusLogger, parkingRequestPostgresRepository, parkingSpacePostgresRepository)
	getAllParkingRequests := usecases2.NewGetAllParkingRequests(logrusLogger, parkingRequestPostgresRepository)
	getDriversParkingRequests := usecases2.NewGetDriversParkingRequests(logrusLogger, parkingRequestPostgresRepository)
	parkingRequestFacade := usecasefacades.NewParkingRequestFacade(createParkingRequest, updateParkingRequestStatus, assignParkingSpace, getAllParkingRequests, getDriversParkingRequests)
	parkingLotPostgresRepository := datastore.NewParkingParkingLotPostgresRepository(logrusLogger, gormWrapper)
	createParkingLot := usecases3.NewCreateParkingLot(logrusLogger, parkingLotPostgresRepository)
	getAllParkingLots := usecases3.NewGetAllParkingLots(logrusLogger, parkingLotPostgresRepository)
	deteleParkingLot := usecases3.NewDeleteParkingLot(logrusLogger, parkingLotPostgresRepository)
	getSingleParkingLot := usecases3.NewGetSingleParkingLot(logrusLogger, parkingLotPostgresRepository)
	parkingLotFacade := usecasefacades.NewParkingLotFacade(createParkingLot, getAllParkingLots, deteleParkingLot, getSingleParkingLot)
	updateParkingSpaceStatus := usecases4.NewUpdateParkingSpaceStatus(logrusLogger, parkingSpacePostgresRepository)
	getSingleParkingSpace := usecases4.NewGetSingleParkingSpace(logrusLogger, parkingSpacePostgresRepository)
	parkingSpaceFacade := usecasefacades.NewParkingSpaceFacade(updateParkingSpaceStatus, getSingleParkingSpace)
	notificationPostgresRepository := datastore.NewNotificationPostgresRepository(logrusLogger, gormWrapper)
	alertPostgresRepository := datastore.NewAlertPostgresRepository(logrusLogger, gormWrapper)
	createAlert := usecases5.NewCreateAlert(logrusLogger, alertPostgresRepository)
	createNotification := usecases6.NewCreateNotification(logrusLogger, notificationPostgresRepository, parkingSpacePostgresRepository, parkingRequestPostgresRepository, createAlert)
	getAllNotifications := usecases6.NewGetAllNotifications(logrusLogger, notificationPostgresRepository)
	notificationFacade := usecasefacades.NewNotificationFacade(createNotification, getAllNotifications)
	getSingleAlert := usecases5.NewGetSingleAlert(logrusLogger, alertPostgresRepository)
	checkLateArrival := usecases5.NewCheckLateArrival(logrusLogger, parkingRequestPostgresRepository, createAlert)
	getAllAlerts := usecases5.NewGetAllAlerts(logrusLogger, alertPostgresRepository)
	alertFacade := usecasefacades.NewAlertFacade(getSingleAlert, checkLateArrival, getAllAlerts)
	facade := handlers.NewFacade(userFacade, parkingRequestFacade, parkingLotFacade, parkingSpaceFacade, notificationFacade, alertFacade)
	handlerFactory := handlers.NewHandlerFactory(logrusLogger, facade)
	checkDriverStatus := usecases.NewCheckDriverStatus(logrusLogger, userPostgresRepository)
	middlewareMiddleware := middleware.NewMiddleware(jwtTokenService, logrusLogger, checkDriverStatus)
	router := routes.NewRouter(handlerFactory, middlewareMiddleware, logrusLogger)
	httpConfig := configConfig.HTTP
	server := httpserver.NewServerFromConfig(router, httpConfig, logrusLogger)
	schedulerConfig := configConfig.Scheduler
	alertConfig := configConfig.Alert
	schedulerScheduler := scheduler.NewSchedulerFromConfig(logrusLogger, alertFacade, schedulerConfig, alertConfig)
	app := NewApp(server, logrusLogger, schedulerScheduler)
	return app, nil
}
