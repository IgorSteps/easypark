package handlers

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// HandlerFactory implements HandlerFactory interface and helps provide dependencies for
// creating different REST Handlers.
type HandlerFactory struct {
	logger *logrus.Logger
	facade UserFacade
}

// NewHandlerFactory provides new instance of the HandlerFactory.
func NewHandlerFactory(logger *logrus.Logger, facade UserFacade) *HandlerFactory {
	return &HandlerFactory{
		logger: logger,
		facade: facade,
	}
}

// DriverCreate returns new REST handler for the creation of driver users.
func (s *HandlerFactory) DriverCreate() http.Handler {
	return NewDriverCreateHandler(s.facade, s.logger)
}

// UserAuthorise returns new REST handler for user authentication.
func (s *HandlerFactory) UserAuthorise() http.Handler {
	return NewUserLoginHandler(s.facade, s.logger)
}

// GetAllDrivers returns new REST handler for getting all driver users.
func (s *HandlerFactory) GetAllDrivers() http.Handler {
	return NewDriverUsersGetHandler(s.logger, s.facade)
}
