package usecasefacades

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/google/uuid"
)

// ParkingSpaceStatusUpdated provides an interface implemented by UpdateParkingSpaceStatus usecase.
type ParkingSpaceStatusUpdater interface {
	Execute(ctx context.Context, id uuid.UUID, status string) (entities.ParkingSpace, error)
}

// ParkingSpaceGetter provides an interface implemented by GetSingleParkingSpace usecase.
type ParkingSpaceGetter interface {
	Execute(ctx context.Context, parkingSpaceID uuid.UUID) (entities.ParkingSpace, error)
}

// ParkingSpaceFacade uses facade pattern to wrap parking space usecases to allow for managing other things such as DB transactions if needed.
type ParkingSpaceFacade struct {
	parkingSpaceStatusUpdater ParkingSpaceStatusUpdater
	parkingSpaceGetter        ParkingSpaceGetter
}

// NewParkingSpaceFacade returns new instance of ParkingSpaceFacade.
func NewParkingSpaceFacade(updater ParkingSpaceStatusUpdater, getter ParkingSpaceGetter) *ParkingSpaceFacade {
	return &ParkingSpaceFacade{
		parkingSpaceStatusUpdater: updater,
		parkingSpaceGetter:        getter,
	}
}

// UpdateParkingSpaceStatus wraps the UpdateParkingSpaceStatus usecase.
func (s *ParkingSpaceFacade) UpdateParkingSpaceStatus(ctx context.Context, id uuid.UUID, status string) (entities.ParkingSpace, error) {
	return s.parkingSpaceStatusUpdater.Execute(ctx, id, status)
}

// GetSingleParkingSpace wraps the GetSingleParkingSpace usecase.
func (s *ParkingSpaceFacade) GetSingleParkingSpace(ctx context.Context, parkingSpaceID uuid.UUID) (entities.ParkingSpace, error) {
	return s.parkingSpaceGetter.Execute(ctx, parkingSpaceID)
}
