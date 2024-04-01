package handlers

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// HandlerFactory implements HandlerFactory interface and helps provide dependencies for
// creating different REST Handlers.
type HandlerFactory struct {
	logger *logrus.Logger
	facade *Facade
}

// NewHandlerFactory provides new instance of the HandlerFactory.
func NewHandlerFactory(logger *logrus.Logger, facade *Facade) *HandlerFactory {
	return &HandlerFactory{
		logger: logger,
		facade: facade,
	}
}

// DriverCreate returns new REST handler for the creation of driver users.
func (s *HandlerFactory) DriverCreate() http.Handler {
	return NewDriverCreateHandler(s.facade.userFacade, s.logger)
}

// UserAuthorise returns new REST handler for user authentication.
func (s *HandlerFactory) UserAuthorise() http.Handler {
	return NewUserLoginHandler(s.facade.userFacade, s.logger)
}

// GetAllDrivers returns new REST handler for getting all driver users.
func (s *HandlerFactory) GetAllDrivers() http.Handler {
	return NewDriverUsersGetHandler(s.logger, s.facade.userFacade)
}

// DriverBan returns new REST handler to ban drivers.
func (s *HandlerFactory) DriverBan() http.Handler {
	return NewDriverStatusHandler(s.facade.userFacade, s.logger)
}

// ParkingRequestCreate returns new REST handler for creating parking requests.
func (s *HandlerFactory) ParkingRequestCreate() http.Handler {
	return NewParkingRequestCreateHandler(s.facade.parkingRequestFacade, s.logger)
}
