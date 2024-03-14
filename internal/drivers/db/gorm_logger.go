package db

import (
	"context"
	"time"

	"github.com/IgorSteps/easypark/internal/drivers/config"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
)

// GormLogrusLogger is a Logrus logger that implements GORM's logger interface.
type GormLogrusLogger struct {
	Logrus   *logrus.Logger
	LogLevel logger.LogLevel
}

// NewGormLogrusLogger initializes and returns a new instance of GormLogrusLogger.
func NewGormLogrusLoggerFromConfig(config config.LoggingConfig, logrus *logrus.Logger) *GormLogrusLogger {
	return &GormLogrusLogger{
		Logrus:   logrus,
		LogLevel: parseLogLevel(config.GormLevel),
	}
}

// LogMode sets log level.
func (l *GormLogrusLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

// Info logs info messages.
func (l *GormLogrusLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		l.Logrus.WithContext(ctx).Infof(msg, data...)
	}
}

// Warn logs warning messages.
func (l *GormLogrusLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		l.Logrus.WithContext(ctx).Warnf(msg, data...)
	}
}

// Error logs error messages.
func (l *GormLogrusLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		l.Logrus.WithContext(ctx).Errorf(msg, data...)
	}
}

// Trace logs SQL queries with data like execution time and affected row count.
func (l *GormLogrusLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= logger.Error:
		sql, _ := fc()
		l.Logrus.WithContext(ctx).WithError(err).WithFields(logrus.Fields{
			"elapsed": elapsed,
			"sql":     sql,
		}).Error("sql error")
	case elapsed > time.Second && l.LogLevel >= logger.Warn: // slow query threshold set to 1s. // TODO: Move to config?
		sql, rows := fc()
		l.Logrus.WithContext(ctx).WithFields(logrus.Fields{
			"elapsed": elapsed,
			"rows":    rows,
			"sql":     sql,
		}).Warn("slow query")
	case l.LogLevel >= logger.Info:
		sql, rows := fc()
		l.Logrus.WithContext(ctx).WithFields(logrus.Fields{
			"elapsed": elapsed,
			"rows":    rows,
			"sql":     sql,
		}).Info("sql query")
	}
}

func parseLogLevel(level string) logger.LogLevel {
	switch level {
	case "silent":
		return logger.Silent
	case "error":
		return logger.Error
	case "warn":
		return logger.Warn
	case "info":
		return logger.Info
	default:
		// info if unknown
		return logger.Info
	}
}
