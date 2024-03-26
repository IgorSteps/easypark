package entities

import (
	"time"

	"github.com/google/uuid"
)

type ParkingSpaceStatus string

const (
	StatusAvailable ParkingSpaceStatus = "available"
	StatusOccupied  ParkingSpaceStatus = "occupied"
	StatusReserved  ParkingSpaceStatus = "reserved"
	StatusBlocked   ParkingSpaceStatus = "blocked"
)

type ParkingSpace struct {
	ID          uuid.UUID
	LotID       uuid.UUID
	Number      string
	Status      ParkingSpaceStatus
	ReservedFor *time.Time
	OccupiedAt  *time.Time
	UserID      uuid.UUID // Reference to the User who has reserved or occupied the space
}
