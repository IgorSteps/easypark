package usecasefacades

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
)

// ParkingLotCreator provides an interface implemented by UpdateParkingRequestSpace usecase.
type ParkingLotCreator interface {
	Execute(ctx context.Context, name string, capacity int) (entities.ParkingLot, error)
}

// ParkingLotFacade uses facade pattern to wrap parking lots' usecases to allow for managing other things such as DB transactions if needed.
type ParkingLotFacade struct {
	parkingLotCreator ParkingLotCreator
}

// NewParkingLotFacade returns a new instance of ParkingLotFacade.
func NewParkingLotFacade(
	creator ParkingLotCreator,
) *ParkingLotFacade {
	return &ParkingLotFacade{
		parkingLotCreator: creator,
	}
}

// CreateParkingLot wraps the UpdateParkingRequestSpace usecase.
func (s *ParkingLotFacade) CreateParkingLot(ctx context.Context, name string, capacity int) (entities.ParkingLot, error) {
	return s.parkingLotCreator.Execute(ctx, name, capacity)
}
