package usecasefacades

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/google/uuid"
)

// AlertSingleGetter provides an interface implemented by GetSingleAlert usecase.
type AlertSingleGetter interface {
	Execute(ctx context.Context, id uuid.UUID) (entities.Alert, error)
}

// AlertFacade uses facade patter to wrap alert usecases to allow for managing other things such as DB transactions if needed.
type AlertFacade struct {
	getter AlertSingleGetter
}

// NewAlertFacade returns a new instance of AlertFacade.
func NewAlertFacade(getter AlertSingleGetter) *AlertFacade {
	return &AlertFacade{
		getter: getter,
	}
}

// GetAlert wraps the GetSingleAlert usecase.
func (s AlertFacade) GetAlert(ctx context.Context, id uuid.UUID) (entities.Alert, error) {
	return s.getter.Execute(ctx, id)
}
