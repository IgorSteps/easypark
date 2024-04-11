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

// ParkingSpace represents a parking space.
type ParkingSpace struct {
	ID              uuid.UUID `gorm:"primary_key"`
	ParkingLotID    uuid.UUID
	Name            string
	Status          ParkingSpaceStatus
	FreeAt          time.Time
	BusyAt          time.Time
	ParkingRequests []ParkingRequest `gorm:"constraint:OnDelete:CASCADE;"`
}

// OnCreate initilises internally managed fields on creation of a parking space.
func (s *ParkingSpace) OnCreate(name string, parkingLotID uuid.UUID) {
	s.ID = uuid.New()
	s.ParkingLotID = parkingLotID
	s.Status = StatusAvailable
	s.Name = name
}

// OnAssign updates a parking space with the request times and changes status to RESERVED.
func (s *ParkingSpace) OnAssign(busyAt, freeAt time.Time) {
	s.Status = StatusReserved
	s.BusyAt = busyAt
	s.FreeAt = freeAt
}

// IsAvailableFor checks if a new parking request times overlap with any existing requests for a parking space.
func (s *ParkingSpace) IsAvailableFor(requestStartTime, requestEndTime time.Time) bool {
	for _, pr := range s.ParkingRequests {
		if pr.StartTime.Before(requestEndTime) && pr.EndTime.After(requestStartTime) {
			return false
		}
	}
	return true
}
