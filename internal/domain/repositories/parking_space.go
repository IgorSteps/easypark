package repositories

import (
	"context"
	"time"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/google/uuid"
)

// ParkingSpaceRepository describes an interfaces that provides CRUD functionality on parking space entity.
type ParkingSpaceRepository interface {
	// GetSingle gets a parking space by ID.
	GetSingle(ctx context.Context, id uuid.UUID) (entities.ParkingSpace, error)

	// GetMany returns parking spaces that match given query.
	GetMany(ctx context.Context, query map[string]interface{}) ([]entities.ParkingSpace, error)

	// Save saves an updated parking space.
	Save(ctx context.Context, space *entities.ParkingSpace) error

	// FindAvailableSpaces finds suitable parking space using given parameters.§
	FindAvailableSpaces(ctx context.Context, lotID uuid.UUID, startTime time.Time, endTime time.Time) ([]entities.ParkingSpace, error)
}
