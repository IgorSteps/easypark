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

// ParkingSpaceFacade uses facade pattern to wrap parking space usecases to allow for managing other things such as DB transactions if needed.
type ParkingSpaceFacade struct {
	parkingSpaceStatusUpdater ParkingSpaceStatusUpdater
}

// NewParkingSpaceFacade returns new instance of ParkingSpaceFacade.
func NewParkingSpaceFacade(updater ParkingSpaceStatusUpdater) *ParkingSpaceFacade {
	return &ParkingSpaceFacade{
		parkingSpaceStatusUpdater: updater,
	}
}

// UpdateParkingSpaceStatus wrpas the UpdateParkingSpaceStatus usecase.
func (s *ParkingSpaceFacade) UpdateParkingSpaceStatus(ctx context.Context, id uuid.UUID, status string) (entities.ParkingSpace, error) {
	return s.parkingSpaceStatusUpdater.Execute(ctx, id, status)
}
