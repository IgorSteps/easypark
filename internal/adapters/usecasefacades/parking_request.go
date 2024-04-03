package usecasefacades

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
)

// ParkingRequestCreator provides an interface implemented by CreateParkingRequest usecase.
type ParkingRequestCreator interface {
	Execute(ctx context.Context, parkingRequest *entities.ParkingRequest) (*entities.ParkingRequest, error)
	CreateParkingRequest(ctx context.Context, parkingRequest *entities.ParkingRequest) (*entities.ParkingRequest, error)
	GetAllRequests(ctx context.Context) ([]entities.ParkingRequest, error)
	UpdateRequest(ctx context.Context, parkingRequest *entities.ParkingRequest) error
	RemoveRequest(ctx context.Context, parkingRequest *entities.ParkingRequest) error
}

// ParkingRequestFacade uses facade pattern to wrap parking request' usecases to allow for managing other things such as DB transactions if needed.
type ParkingRequestFacade struct {
	parkingRequestCreator ParkingRequestCreator
}

// NewParkingRequestFacade creates a new instance of ParkingRequestFacade.
func NewParkingRequestFacade(creator ParkingRequestCreator) *ParkingRequestFacade {
	return &ParkingRequestFacade{
		parkingRequestCreator: creator,
	}
}

// CreateParkingRequest wraps the CreateParkingRequest usecase.
func (s *ParkingRequestFacade) CreateParkingRequest(ctx context.Context, parkingRequest *entities.ParkingRequest) (*entities.ParkingRequest, error) {
	return s.parkingRequestCreator.Execute(ctx, parkingRequest)
}
