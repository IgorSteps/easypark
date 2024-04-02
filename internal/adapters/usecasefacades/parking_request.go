package usecasefacades

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/google/uuid"
)

// ParkingRequestCreator provides an interface implemented by CreateParkingRequest usecase.
type ParkingRequestCreator interface {
	Execute(ctx context.Context, parkingRequest *entities.ParkingRequest) (*entities.ParkingRequest, error)
}

// ParkingRequestUpdater provides an interface implemented by UpdateParkingRequest usecase.
type ParkingRequestStatusUpdater interface {
	Execute(ctx context.Context, id uuid.UUID, status string) error
}

// ParkingRequestFacade uses facade pattern to wrap parking request' usecases to allow for managing other things such as DB transactions if needed.
type ParkingRequestFacade struct {
	parkingRequestCreator        ParkingRequestCreator
	parkgingRequestStatusUpdater ParkingRequestStatusUpdater
}

// NewParkingRequestFacade creates a new instance of ParkingRequestFacade.
func NewParkingRequestFacade(creator ParkingRequestCreator, updater ParkingRequestStatusUpdater) *ParkingRequestFacade {
	return &ParkingRequestFacade{
		parkingRequestCreator:        creator,
		parkgingRequestStatusUpdater: updater,
	}
}

// CreateParkingRequest wraps the CreateParkingRequest usecase.
func (s *ParkingRequestFacade) CreateParkingRequest(ctx context.Context, parkingRequest *entities.ParkingRequest) (*entities.ParkingRequest, error) {
	return s.parkingRequestCreator.Execute(ctx, parkingRequest)
}

// UpdateParkingRequestStatus wraps the UpdateParkingRequestStatus usecase.
func (s *ParkingRequestFacade) UpdateParkingRequestStatus(ctx context.Context, id uuid.UUID, status string) error {
	return s.parkgingRequestStatusUpdater.Execute(ctx, id, status)
}
