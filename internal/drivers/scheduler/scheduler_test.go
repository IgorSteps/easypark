package scheduler_test

import (
	"testing"
	"time"

	"github.com/IgorSteps/easypark/internal/drivers/config"
	"github.com/IgorSteps/easypark/internal/drivers/scheduler"
	mocks "github.com/IgorSteps/easypark/mocks/adapters/rest/handlers"
	"github.com/sirupsen/logrus/hooks/test"
)

func TestScheduler_NewScheduler(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	schedulerConfig := config.SchedulerConfig{
		Interval: "@hourly",
	}
	alertConfig := config.AlertConfig{
		LateArrivalThresholdMinutes: time.Hour,
		OverStayThresholdMinutes:    time.Hour,
	}
	facade := &mocks.AlertFacade{}

	// --------
	// ACT
	// --------
	scheduler.NewSchedulerFromConfig(testLogger, facade, schedulerConfig, alertConfig)

	// --------
	// ASSERT
	// --------
}
