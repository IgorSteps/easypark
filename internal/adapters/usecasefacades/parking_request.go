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

// ParkingRequestsGetter provides an interface implemented by the GetAllParkingRequests usecase.
type ParkingRequestsGetter interface {
	Execute(ctx context.Context) ([]entities.ParkingRequest, error)
}

// ParkingRequestUpdater provides an interface implemented by UpdateParkingRequest usecase.
type ParkingRequestStatusUpdater interface {
	Execute(ctx context.Context, id uuid.UUID, status string) error
}

// ParkingRequestSpaceAssigner provides an interface implemented by UpdateParkingRequestSpace usecase.
type ParkingRequestSpaceAssigner interface {
	Execute(ctx context.Context, requestID uuid.UUID, spaceID uuid.UUID) error
}

// ParkingRequestFacade uses facade pattern to wrap parking request' usecases to allow for managing other things such as DB transactions if needed.
type ParkingRequestFacade struct {
	parkingRequestCreator        ParkingRequestCreator
	parkgingRequestStatusUpdater ParkingRequestStatusUpdater
	parkingRequestSpaceAssigner  ParkingRequestSpaceAssigner
	parkingRequestGetter         ParkingRequestsGetter
}

// NewParkingRequestFacade creates a new instance of ParkingRequestFacade.
func NewParkingRequestFacade(
	creator ParkingRequestCreator,
	updater ParkingRequestStatusUpdater,
	assigner ParkingRequestSpaceAssigner,
	getter ParkingRequestsGetter,
) *ParkingRequestFacade {
	return &ParkingRequestFacade{
		parkingRequestCreator:        creator,
		parkgingRequestStatusUpdater: updater,
		parkingRequestSpaceAssigner:  assigner,
		parkingRequestGetter:         getter,
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

// AssignParkingSpace wraps the UpdateParkingRequestSpace usecase.
func (s *ParkingRequestFacade) AssignParkingSpace(ctx context.Context, requestID uuid.UUID, spaceID uuid.UUID) error {
	return s.parkingRequestSpaceAssigner.Execute(ctx, requestID, spaceID)
}

func (s *ParkingRequestFacade) GetAllParkingRequests(ctx context.Context) ([]entities.ParkingRequest, error) {
	return s.parkingRequestGetter.Execute(ctx)
}
