package repositories

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
)

type ParkingLotRepository interface {
	CreateParkingLot(ctx context.Context, parkingLot *entities.ParkingLot) error
}
