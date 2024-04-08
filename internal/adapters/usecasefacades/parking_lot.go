package usecasefacades

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
)

// ParkingLotCreator provides an interface implemented by ParkingLotCreate usecase.
type ParkingLotCreator interface {
	Execute(ctx context.Context, name string, capacity int) (entities.ParkingLot, error)
}

// ParkingLotGetter provides an interface implemented by GetAllParkingLots usecase.
type ParkingLotGetter interface {
	Execute(ctx context.Context) ([]entities.ParkingLot, error)
}

// ParkingLotFacade uses facade pattern to wrap parking lots' usecases to allow for managing other things such as DB transactions if needed.
type ParkingLotFacade struct {
	parkingLotCreator ParkingLotCreator
	pakringLotGetter  ParkingLotGetter
}

// NewParkingLotFacade returns a new instance of ParkingLotFacade.
func NewParkingLotFacade(
	creator ParkingLotCreator,
	getter ParkingLotGetter,
) *ParkingLotFacade {
	return &ParkingLotFacade{
		parkingLotCreator: creator,
		pakringLotGetter:  getter,
	}
}

// CreateParkingLot wraps the CreateParingLot usecase.
func (s *ParkingLotFacade) CreateParkingLot(ctx context.Context, name string, capacity int) (entities.ParkingLot, error) {
	return s.parkingLotCreator.Execute(ctx, name, capacity)
}

// GetAllParkingLots wraps the GetAllParkingLots usease.
func (s *ParkingLotFacade) GetAllParkingLots(ctx context.Context) ([]entities.ParkingLot, error) {
	return s.pakringLotGetter.Execute(ctx)
}
