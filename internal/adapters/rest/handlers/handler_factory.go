package handlers

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// HandlerFactory implements HandlerFactory interface and helps provide implementation for
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

// UserCreate returns new REST handler for the creation of users.
func (s *HandlerFactory) UserCreate() http.Handler {
	return NewUserCreateHandler(s.facade, s.logger)
}

// UserAuthorise returns new REST handler for user authentication.
func (s *HandlerFactory) UserAuthorise() http.Handler {
	return NewUserLoginHandler(s.facade, s.logger)
}
