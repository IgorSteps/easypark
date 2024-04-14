package repositories

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/google/uuid"
)

// AlertRepository defines an interface that provides CRUD operations on Alert entity.
type AlertRepository interface {
	// Create creates a new alert in the database.
	Create(ctx context.Context, alert *entities.Alert) error

	// GetSingle returns a specific alert using its ID.
	GetSingle(ctx context.Context, alertID uuid.UUID) (entities.Alert, error)
}

// AlertCreator provides an interface implemented by the CreateAlert usecase.
type AlertCreator interface {
	Execute(ctx context.Context, alertType entities.AlertType, msg string, driverID, spaceID uuid.UUID) (*entities.Alert, error)
}
