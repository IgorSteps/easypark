package entities

import (
	"time"

	"github.com/google/uuid"
)

type ParkingRequestStatus string

const (
	RequestStatusPending  ParkingRequestStatus = "pending"
	RequestStatusApproved ParkingRequestStatus = "approved"
	RequestStatusRejected ParkingRequestStatus = "rejected"
)

type ParkingRequest struct {
	ID               uuid.UUID `gorm:"primary_key"`
	UserID           uuid.UUID
	ParkingSpaceID   *uuid.UUID // Can be nil, because Admin chooses it after request is created.
	ParkingSpaceName *string    // Same as above.
	Destination      string     // Parking Lot name.
	StartTime        time.Time
	EndTime          time.Time
	Status           ParkingRequestStatus
}

func (s *ParkingRequest) OnCreate() {
	s.ID = uuid.New()
	s.Status = RequestStatusPending
}
