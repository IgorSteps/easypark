package repositories

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
)

// ParkingLotRepository describes an interface for CRUD operations on parking lots.
type ParkingLotRepository interface {
	// CreateParkingLot creates a parking lot.
	CreateParkingLot(ctx context.Context, parkingLot *entities.ParkingLot) error

	// GetAllParkingLots gets all parking lots.
	GetAllParkingLots(ctx context.Context) ([]entities.ParkingLot, error)
}
