package usecasefacades

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/google/uuid"
)

// ParkingLotCreator provides an interface implemented by ParkingLotCreate usecase.
type ParkingLotCreator interface {
	Execute(ctx context.Context, name string, capacity int) (entities.ParkingLot, error)
}

// ParkingLotGetter provides an interface implemented by GetAllParkingLots usecase.
type ParkingLotGetter interface {
	Execute(ctx context.Context) ([]entities.ParkingLot, error)
}

// ParkingLotGetter provides an interface implemented by GetSingleParkingLot usecase.
type ParkingLotSingleGetter interface {
	Execute(ctx context.Context, id uuid.UUID) (*entities.ParkingLot, error)
}

// ParkingLotDeleter provides an interface implemented by DeleteParkingLot usecase.
type ParkingLotDeleter interface {
	Execute(ctx context.Context, id uuid.UUID) error
}

// ParkingLotFacade uses facade pattern to wrap parking lots' usecases to allow for managing other things such as DB transactions if needed.
type ParkingLotFacade struct {
	parkingLotCreator      ParkingLotCreator
	pakringLotGetter       ParkingLotGetter
	parkingLotDeleter      ParkingLotDeleter
	parkingLotSingleGetter ParkingLotSingleGetter
}

// NewParkingLotFacade returns a new instance of ParkingLotFacade.
func NewParkingLotFacade(
	creator ParkingLotCreator,
	getter ParkingLotGetter,
	deleter ParkingLotDeleter,
	sGetter ParkingLotSingleGetter,
) *ParkingLotFacade {
	return &ParkingLotFacade{
		parkingLotCreator:      creator,
		pakringLotGetter:       getter,
		parkingLotDeleter:      deleter,
		parkingLotSingleGetter: sGetter,
	}
}

// CreateParkingLot wraps the CreateParingLot usecase.
func (s *ParkingLotFacade) CreateParkingLot(ctx context.Context, name string, capacity int) (entities.ParkingLot, error) {
	return s.parkingLotCreator.Execute(ctx, name, capacity)
}

// GetAllParkingLots wraps the GetAllParkingLots usecase.
func (s *ParkingLotFacade) GetAllParkingLots(ctx context.Context) ([]entities.ParkingLot, error) {
	return s.pakringLotGetter.Execute(ctx)
}

// DeleteParkingLot wraps the DeleteParkingLot usecase.
func (s *ParkingLotFacade) DeleteParkingLot(ctx context.Context, id uuid.UUID) error {
	return s.parkingLotDeleter.Execute(ctx, id)
}

// GetSingleParkingLot wraps the GetSingleParkingLot usecase.
func (s *ParkingLotFacade) GetSingleParkingLot(ctx context.Context, lotID uuid.UUID) (*entities.ParkingLot, error) {
	return s.parkingLotSingleGetter.Execute(ctx, lotID)
}
