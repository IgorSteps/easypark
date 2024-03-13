package httpserver

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Server represents the HTTP server that handles the requests using chi.
type Server struct {
	router chi.Router
}

// NewServer creates a new Server instance.
func NewServer(r chi.Router) *Server {
	return &Server{router: r}
}

// Run starts the HTTP server on the given address.
func (s *Server) Run(addr string) error {
	log.Printf("Starting HTTP server on %s", addr)
	return http.ListenAndServe(addr, s.router)
}
