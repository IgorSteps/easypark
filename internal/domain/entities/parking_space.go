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
	ID              uuid.UUID `gorm:"primary_key"`
	ParkingLotID    uuid.UUID
	Name            string
	Status          ParkingSpaceStatus
	ParkingRequests []ParkingRequest `gorm:"constraint:OnDelete:CASCADE;"`
}

func (s *ParkingSpace) OnCreate(name string, parkingLotID uuid.UUID) {
	s.ID = uuid.New()
	s.ParkingLotID = parkingLotID
	s.Status = StatusAvailable
	s.Name = name
}

// CheckForOverlap checks that the new request's time slot doesn't overlap with existing parking requests' time slots.
// TODO: TVery naive way of checking for overlap, hence the performance is shit, refactor in the future
func (s *ParkingSpace) CheckForOverlap(requestStartTime, requestEndTime time.Time) bool {
	for _, parkingRequest := range s.ParkingRequests {
		// Checks if the new request completely overlaps an existing request
		if requestStartTime.Before(parkingRequest.StartTime) && requestEndTime.After(parkingRequest.EndTime) {
			return true
		}
		// Checks if the start time of the new request is within an existing request
		if requestStartTime.After(parkingRequest.StartTime) && requestStartTime.Before(parkingRequest.EndTime) {
			return true
		}
		// Checks if the end time of the new request is within an existing request
		if requestEndTime.After(parkingRequest.StartTime) && requestEndTime.Before(parkingRequest.EndTime) {
			return true
		}
		// Checks if the new request is exactly the same as an existing request
		if requestStartTime.Equal(parkingRequest.StartTime) && requestEndTime.Equal(parkingRequest.EndTime) {
			return true
		}
	}
	// No overlap
	return false
}
