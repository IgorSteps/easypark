package httpserver

import (
	"net/http"

	"github.com/IgorSteps/easypark/internal/drivers/config"
	"github.com/go-chi/chi/v5"
)

// Server represents the HTTP server that handles the requests using chi.
type Server struct {
	Router  chi.Router
	Address string
}

// NewServer creates a new Server instance.
func NewServerFromConfig(r chi.Router, config config.HTTPConfig) *Server {
	return &Server{
		Router:  r,
		Address: config.Address,
	}
}

// Run starts the HTTP server on the given address.
func (s *Server) Run() error {
	return http.ListenAndServe(s.Address, s.Router)
}
