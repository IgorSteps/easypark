package entities

import (
	"time"

	"github.com/google/uuid"
)

// ParkingSpaceStatus represents status of a parking space.
type ParkingSpaceStatus string

const (
	ParkingSpaceStatusAvailable ParkingSpaceStatus = "available"
	ParkingSpaceStatusOccupied  ParkingSpaceStatus = "occupied"
	ParkingSpaceStatusReserved  ParkingSpaceStatus = "reserved"
	ParkingSpaceStatusBlocked   ParkingSpaceStatus = "blocked"
)

// ParkingSpace represents a parking space.
type ParkingSpace struct {
	ID              uuid.UUID `gorm:"primary_key"`
	ParkingLotID    uuid.UUID
	Name            string
	Status          ParkingSpaceStatus
	ParkingRequests []ParkingRequest `gorm:"constraint:OnDelete:CASCADE;"`
}

// OnCreate sets parking space fields on creation.
func (s *ParkingSpace) OnCreate(name string, parkingLotID uuid.UUID) {
	s.ID = uuid.New()
	s.ParkingLotID = parkingLotID
	s.Status = ParkingSpaceStatusAvailable
	s.Name = name
}

// OnArrival changes the status of parking space to 'occupied'.
func (s *ParkingSpace) OnArrival() {
	s.Status = ParkingSpaceStatusOccupied
}

// OnDeparture changes the status of parking space to 'available'.
func (s *ParkingSpace) OnDeparture() {
	s.Status = ParkingSpaceStatusAvailable
}

// CheckForOverlap checks that the new request's time slot doesn't overlap with existing parking requests' time slots.
func (s *ParkingSpace) CheckForOverlap(requestStartTime, requestEndTime time.Time) bool {
	// TODO: Very naive way of checking for overlap, hence the performance is not great for large number of parking requests,
	// but I think it is okay for an MVP.
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
