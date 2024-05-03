package usecasefacades

import (
	"context"
	"time"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/google/uuid"
)

// AlertSingleGetter provides an interface implemented by GetSingleAlert usecase.
type AlertSingleGetter interface {
	Execute(ctx context.Context, id uuid.UUID) (entities.Alert, error)
}

// AlertAllGetter provides an interface implemented by GetAllAlerts usecase.
type AlertAllGetter interface {
	Execute(ctx context.Context) ([]entities.Alert, error)
}

// AlertLateArrivalChecker provides an interface implemented by the CheckLaterAlert usecase.
type AlertLateArrivalChecker interface {
	Execute(ctx context.Context, threshold time.Duration) ([]entities.Alert, error)
}

// AlertOverStayChecker provides an interface implemented by the CheckOverStay usecase.
type AlertOverStayChecker interface {
	Execute(ctx context.Context, threshold time.Duration) ([]entities.Alert, error)
}

// AlertFacade uses facade patter to wrap alert usecases to allow for managing other things such as DB transactions if needed.
type AlertFacade struct {
	singleGetter       AlertSingleGetter
	lateArrivalChecker AlertLateArrivalChecker
	overStayChecker    AlertOverStayChecker
	allGetter          AlertAllGetter
}

// NewAlertFacade returns a new instance of AlertFacade.
func NewAlertFacade(getter AlertSingleGetter, lateChecker AlertLateArrivalChecker, overStayCheck AlertOverStayChecker, getAll AlertAllGetter) *AlertFacade {
	return &AlertFacade{
		singleGetter:       getter,
		lateArrivalChecker: lateChecker,
		overStayChecker:    overStayCheck,
		allGetter:          getAll,
	}
}

// GetAlert wraps the GetSingleAlert usecase.
func (s *AlertFacade) GetAlert(ctx context.Context, id uuid.UUID) (entities.Alert, error) {
	return s.singleGetter.Execute(ctx, id)
}

// CheckForLateArrivals wraps the CheckLateArrival usecase.
func (s *AlertFacade) CheckForLateArrivals(ctx context.Context, threshold time.Duration) ([]entities.Alert, error) {
	return s.lateArrivalChecker.Execute(ctx, threshold)
}

// CheckForOverStay wraps the CheckOverStay usecase.
func (s *AlertFacade) CheckForOverStays(ctx context.Context, threshold time.Duration) ([]entities.Alert, error) {
	return s.overStayChecker.Execute(ctx, threshold)
}

// GetAlert wraps the GetAllAlerts usecase.
func (s *AlertFacade) GetAllAlerts(ctx context.Context) ([]entities.Alert, error) {
	return s.allGetter.Execute(ctx)
}
