package usecases

import (
	"context"
	"time"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/sirupsen/logrus"
)

// CheckLateArrival provides business logic to check when an arrival notification hasn't been received within a given threshold
// from the parking request start time.
type CheckLateArrival struct {
	logger             *logrus.Logger
	parkingRequestRepo repositories.ParkingRequestRepository
	alertCreator       repositories.AlertCreator
}

// NewCheckLateArrival returns a new instance of CheckLateArrival.
func NewCheckLateArrival(
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

// Execute runs the business logic for CheckLateArrival.
func (s *CheckLateArrival) Execute(ctx context.Context, threshold time.Duration) ([]entities.Alert, error) {
	// All requests in the past must have 'completed', current requests have either 'active' if received arrival notification
	// or 'approved' if not, hence we look for requests with 'approved' status and with start time less than 1 hour ago.
	query := map[string]interface{}{
		"status": entities.RequestStatusApproved,
	}
	reqs, err := s.parkingRequestRepo.GetMany(ctx, query)
	if err != nil {
		return nil, err
	}

	// TODO:
	// GORM's map conditionals only support equalities, so will have to filter by time manually.
	// We need to move this to a method on the parking request repository.
	var timeFilteredReqs []entities.ParkingRequest
	for _, req := range reqs {
		if req.StartTime.Before(time.Now().Add(-threshold)) {
			timeFilteredReqs = append(timeFilteredReqs, req)
		}
	}

	s.logger.WithFields(logrus.Fields{
		"time now": time.Now(),
		"requests": timeFilteredReqs,
	}).Debug("got requests that trigger late arrival alert")

	var alerts []entities.Alert
	for _, req := range timeFilteredReqs {
		alert, err := s.alertCreator.Execute(
			ctx,
			entities.LateArrival,
			"arrival notification hasn't been received within one hour from the parking request start time",
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
