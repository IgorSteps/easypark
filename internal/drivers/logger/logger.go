package logger

import (
	"github.com/IgorSteps/easypark/internal/drivers/config"
	"github.com/sirupsen/logrus"
)

const (
	jsonFormat = "json"
	textFormat = "text"
)

// NewLoggerFromConfig returns new logrus logger from given config.
func NewLoggerFromConfig(config config.LoggingConfig) *logrus.Logger {
	logger := logrus.New()

	logLevel, err := logrus.ParseLevel(config.Level)
	if err != nil {
		logger.WithError(err).Error("invalid log level, setting to info")
		logger.SetLevel(logrus.InfoLevel)
	} else {
		logger.SetLevel(logLevel)
	}

	switch config.Format {
	case textFormat:
		logger.SetFormatter(&logrus.TextFormatter{})
	case jsonFormat:
		logger.SetFormatter(&logrus.JSONFormatter{})
	default:
		logger.WithField("given format", config.Format).Info("unknown log format, setting default 'text' format")
	}

	return logger
}
