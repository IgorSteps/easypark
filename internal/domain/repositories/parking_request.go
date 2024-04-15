package repositories

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/google/uuid"
)

// ParkingRequestRepository provides an interface for CRUD operations on parking requests.
type ParkingRequestRepository interface {
	// Create creates a parking request.
	Create(ctx context.Context, parkReq *entities.ParkingRequest) error

	// GetMany gets many parking requests that match given query.
	GetMany(ctx context.Context, query map[string]interface{}) ([]entities.ParkingRequest, error)

	// GetSingle gets a single parking request using its ID.
	GetSingle(ctx context.Context, id uuid.UUID) (entities.ParkingRequest, error)

	// Save saves the parking request when performing any updating.
	Save(ctx context.Context, request *entities.ParkingRequest) error
}
