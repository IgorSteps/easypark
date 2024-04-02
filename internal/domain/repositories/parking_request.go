package repositories

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/google/uuid"
)

// ParkingRequestRepository provides an interface for CRUD operations on parking requests.
type ParkingRequestRepository interface {
	// CreateParkingRequest creates a parking request.
	CreateParkingRequest(ctx context.Context, parkReq *entities.ParkingRequest) error

	// GetAllParkingRequests gets all parking requests.
	GetAllParkingRequests(ctx context.Context) ([]entities.ParkingRequest, error)

	// GetParkingRequestByID gets a single parking request using its ID.
	GetParkingRequestByID(ctx context.Context, id uuid.UUID) (entities.ParkingRequest, error)

	// GetAllParkingRequests gets all parking requests for a particular user.
	GetAllParkingRequestsForUser(ctx context.Context, userID uuid.UUID) ([]entities.ParkingRequest, error)

	// Save saves the parking request when performing any updating.
	Save(ctx context.Context, request *entities.ParkingRequest) error
}
