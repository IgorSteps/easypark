package websocketserver

import (
	"net/http"

	"github.com/IgorSteps/easypark/internal/drivers/config"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

// Server represents the Websocket server that handles the requests using chi.
type Server struct {
	Router  chi.Router
	Logger  *logrus.Logger
	Address string
}

// NewServer creates a new Server instance.
func NewServerFromConfig(r chi.Router, config config.HTTPConfig, l *logrus.Logger) *Server {
	return &Server{
		Router:  r,
		Address: config.WebsocketAddress,
		Logger:  l,
	}
}

// Run starts the websocket server on the given address.
func (s *Server) Run() error {
	s.Logger.WithField("address", s.Address).Info("starting websocket server")
	return http.ListenAndServe(s.Address, s.Router)
}
