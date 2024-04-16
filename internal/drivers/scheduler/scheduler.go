package scheduler

import (
	"context"

	"github.com/IgorSteps/easypark/internal/adapters/rest/handlers"
	"github.com/IgorSteps/easypark/internal/drivers/config"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

// Scheduler provides scheduling for cron jobs within the service,
// that allows for automatic tasks without relying on external triggers.
type Scheduler struct {
	cron            *cron.Cron
	schedulerConfig config.SchedulerConfig
	alertConfig     config.AlertConfig
	logger          *logrus.Logger
	facade          handlers.AlertFacade
}

// NewSchedulerFromConfig returns a new instance of Scheduler.
func NewSchedulerFromConfig(
	logger *logrus.Logger,
	facade handlers.AlertFacade,
	schedulerConfig config.SchedulerConfig,
	alertConfig config.AlertConfig,
) *Scheduler {
	return &Scheduler{
		cron:            cron.New(),
		logger:          logger,
		facade:          facade,
		schedulerConfig: schedulerConfig,
		alertConfig:     alertConfig,
	}
}

// Start starts the cron scheduler in its own go-routine.
func (s *Scheduler) Start() {
	s.logger.WithField("interval", s.schedulerConfig.Interval).Info("starting scheduler")
	s.cron.AddFunc(s.schedulerConfig.Interval, func() {
		_, err := s.facade.CheckForLateArrivals(context.TODO(), s.alertConfig.LateArrivalThresholdMinutes)
		if err != nil {
			s.logger.WithError(err).Error("failed to check for late arrivals")
			return
		}
	})
	s.cron.Start()
}

// Stop stops the cron scheduler.
func (s *Scheduler) Stop() {
	s.logger.Info("shutting down scheduler")
	s.cron.Stop()
}
