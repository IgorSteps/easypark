package logger

import (
	"github.com/IgorSteps/easypark/internal/drivers/config"
	"github.com/sirupsen/logrus"
)

// NewLoggerFromConfig returns new logrus logger from given config.
func NewLoggerFromConfig(config config.LoggingConfig) *logrus.Logger {
	logger := logrus.New()
	logLevel, err := logrus.ParseLevel(config.Level)
	if err != nil {
		logger.Errorf("Invalid log level: %s", config.Level)
	} else {
		logger.SetLevel(logLevel)
	}

	return logger
}
