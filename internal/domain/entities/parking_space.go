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
	ID           uuid.UUID `gorm:"primary_key"`
	ParkingLotID uuid.UUID
	Number       string
	Status       ParkingSpaceStatus
	ReservedFor  *time.Time
	OccupiedAt   *time.Time
	UserID       *uuid.UUID
}
