package logger_test

import (
	"testing"

	"github.com/IgorSteps/easypark/internal/drivers/config"
	"github.com/IgorSteps/easypark/internal/drivers/logger"
	"github.com/stretchr/testify/assert"
)

func TestNewLoggerFromConfig(t *testing.T) {
	// -------
	// ASSEMLE
	// -------
	config := config.LoggingConfig{Level: "info"}

	// ----
	// ACT
	// ----
	logger := logger.NewLoggerFromConfig(config)

	// -------
	// ASSERT
	// -------
	assert.NotNil(t, logger, "Logger must not be nil")
	assert.Equal(t, "info", logger.GetLevel().String(), "Log level must be info")
}

func TestNewLoggerFromConfig_InvalidLogLevel(t *testing.T) {
	// -------
	// ASSEMLE
	// -------
	config := config.LoggingConfig{Level: "blob"}

	// ----
	// ACT
	// ----
	logger := logger.NewLoggerFromConfig(config)

	// -------
	// ASSERT
	// -------
	assert.NotNil(t, logger, "Logger must not be nil")
	assert.Equal(t, "info", logger.GetLevel().String(), "Log level must be info")

}
