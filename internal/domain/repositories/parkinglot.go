package repositories

import (
    "context"

    "github.com/IgorSteps/easypark/internal/domain/entities"
    "github.com/google/uuid"
)

// ParkingLotRepository provides an interface for CRUD operations on parking lots.
type ParkingLotRepository interface {
    // CreateParkingLot adds a new parking lot to the database.
    CreateParkingLot(ctx context.Context, parkingLot *entities.ParkingLot) error

    // GetParkingLotByID retrieves a parking lot by its ID.
    GetParkingLotByID(ctx context.Context, id uuid.UUID) (*entities.ParkingLot, error)

    // GetAllParkingLots retrieves all parking lots from the database.
    GetAllParkingLots(ctx context.Context) ([]entities.ParkingLot, error)

    // UpdateParkingLot updates an existing parking lot's details in the database.
    UpdateParkingLot(ctx context.Context, parkingLot *entities.ParkingLot) error

    // DeleteParkingLot removes a parking lot from the database by its ID.
    DeleteParkingLot(ctx context.Context, id uuid.UUID) error
}
