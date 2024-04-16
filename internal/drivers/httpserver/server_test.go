package httpserver_test

import (
	"testing"

	"github.com/IgorSteps/easypark/internal/drivers/httpserver"
	"github.com/sirupsen/logrus/hooks/test"

	"github.com/IgorSteps/easypark/internal/drivers/config"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	router := chi.NewRouter()
	config := config.HTTPConfig{Address: ":8080"}
	logger, _ := test.NewNullLogger()
	// ----
	// ACT
	// ----
	server := httpserver.NewServerFromConfig(router, config, logger)

	// -------
	// ASSERT
	// -------
	assert.NotNil(t, server, "Server must not be nil")
	assert.Equal(t, router, server.Router, "Server router must be the same as the provided one")
}
