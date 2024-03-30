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

	// GetAllParkingRequests gets all parking requests for a particular user.
	GetAllParkingRequestsForUser(ctx context.Context, userID uuid.UUID) ([]entities.ParkingRequest, error)

	// UpdateParkingRequest updates a parking request.
	UpdateParkingRequest(ctx context.Context, parkingRequestToUpdateID uuid.UUID, updatedRequest *entities.ParkingRequest) error
}
