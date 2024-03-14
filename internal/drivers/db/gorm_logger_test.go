package db_test

import (
	"context"
	"testing"
	"time"

	"github.com/IgorSteps/easypark/internal/drivers/config"
	"github.com/IgorSteps/easypark/internal/drivers/db"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/logger"
)

func TestGormLogrusLogger_Info(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	gormLogger, hook := NewTestGormLogrusLogger("info")
	ctx := context.Background()

	// --------
	// ACT
	// --------
	gormLogger.Info(ctx, "boom %s", "badaboom")

	// --------
	// ASSERT
	// --------
	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "boom badaboom", hook.LastEntry().Message)

	hook.Reset()
	assert.Nil(t, hook.LastEntry())
}

func TestGormLogrusLogger_Warn(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	gormLogger, hook := NewTestGormLogrusLogger("warn")
	ctx := context.Background()

	// --------
	// ACT
	// --------
	gormLogger.Warn(ctx, "boom %s", "badaboom")

	// --------
	// ASSERT
	// --------
	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.WarnLevel, hook.LastEntry().Level)
	assert.Equal(t, "boom badaboom", hook.LastEntry().Message)

	hook.Reset()
	assert.Nil(t, hook.LastEntry())
}

func TestGormLogrusLogger_Error(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	gormLogger, hook := NewTestGormLogrusLogger("error")
	ctx := context.Background()

	// --------
	// ACT
	// --------
	gormLogger.Error(ctx, "boom %s", "badaboom")

	// --------
	// ASSERT
	// --------
	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "boom badaboom", hook.LastEntry().Message)

	hook.Reset()
	assert.Nil(t, hook.LastEntry())
}

func TestGormLogrusLogger_Trace(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	ctx := context.Background()

	// Define test cases
	tests := []struct {
		name        string
		execTime    time.Duration
		logLevel    string
		err         error
		expectedLog string
	}{
		{
			name:        "Trace info",
			execTime:    10 * time.Millisecond,
			logLevel:    "info",
			err:         nil,
			expectedLog: "sql query",
		},
		{
			name:        "Trace warn for slow query",
			execTime:    2 * time.Second, // our threshold is 1s
			logLevel:    "warn",
			err:         nil,
			expectedLog: "slow query",
		},
		{
			name:        "Trace error",
			execTime:    10 * time.Millisecond,
			logLevel:    "error",
			err:         logger.ErrRecordNotFound,
			expectedLog: "sql error",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gormLogger, hook := NewTestGormLogrusLogger(tc.logLevel)
			startTime := time.Now()

			// --------
			// ACT
			// --------
			gormLogger.Trace(ctx, startTime.Add(-tc.execTime), func() (string, int64) {
				return "SELECT * FROM users;", 1
			}, tc.err)

			// --------
			// ASSERT
			// --------
			assert.NotEmpty(t, hook.Entries, "There should be log entries")
			entry := hook.LastEntry()
			assert.Contains(t, entry.Message, tc.expectedLog, "Log message should match expected")

			hook.Reset()
			assert.Nil(t, hook.LastEntry())
		})
	}
}

// Helper function to create a GormLogrusLogger with a mocked Logrus logger
func NewTestGormLogrusLogger(logLevel string) (*db.GormLogrusLogger, *test.Hook) {
	logrusLogger, hook := test.NewNullLogger()
	config := config.LoggingConfig{GormLevel: logLevel}
	gormLogger := db.NewGormLogrusLoggerFromConfig(config, logrusLogger)
	return gormLogger, hook
}
