package httpserver

import (
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	router := chi.NewRouter()

	// ----
	// ACT
	// ----
	server := NewServer(router)

	// -------
	// ASSERT
	// -------
	assert.NotNil(t, server, "Server should not be nil")
	assert.Equal(t, router, server.router, "Server router should match the given router")
}
