// get all active parking requests, check within threshold, make note of overdue stay
package usecases

import (
	"context"
	"time"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/sirupsen/logrus"
)

type CheckOverStays struct {
	logger             *logrus.Logger
	parkingRequestRepo repositories.ParkingRequestRepository
	alertCreator       repositories.AlertCreator
}

// NewCheckOverStays returns a new instance of CheckOverStays.
func NewCheckOverStays(
	l *logrus.Logger,
	r repositories.ParkingRequestRepository,
	aCreator repositories.AlertCreator,
) *CheckLateArrival {
	return &CheckLateArrival{
		logger:             l,
		parkingRequestRepo: r,
		alertCreator:       aCreator,
	}
}

// Execute runs the business logic for CheckOverStay.
func (s *CheckOverStays) Execute(ctx context.Context, threshold time.Duration) ([]entities.Alert, error) {
	//look for requests with current time 1+ hour from request end time
	query := map[string]interface{}{
		"status": entities.RequestStatusActive,
	}
	reqs, err := s.parkingRequestRepo.GetMany(ctx, query)
	if err != nil {
		return nil, err
	}

	var timeFilteredReqs []entities.ParkingRequest
	for _, req := range reqs {
		if (time.Now().Second() - req.EndTime.Second()) > int(threshold.Seconds()) { //filter overstays
			timeFilteredReqs = append(timeFilteredReqs, req)
		}
	}

	s.logger.WithFields(logrus.Fields{
		"time now": time.Now(),
		"requests": timeFilteredReqs,
	}).Debug("got requests that trigger overstay alerts")

	var alerts []entities.Alert
	for _, req := range timeFilteredReqs {
		alert, err := s.alertCreator.Execute(
			ctx,
			entities.OverStay,
			"exit notification hasn't been received within one hour from the parking request end time",
			req.UserID,
			*req.ParkingSpaceID,
		)
		if err != nil {
			return nil, err
		}

		alerts = append(alerts, *alert)
	}

	return alerts, nil
}
