package entities

import (
	"time"

	"github.com/google/uuid"
)

// ParkingRequestStatus represents status of a parking request.
type ParkingRequestStatus string

const (
	RequestStatusPending   ParkingRequestStatus = "pending"
	RequestStatusApproved  ParkingRequestStatus = "approved"
	RequestStatusRejected  ParkingRequestStatus = "rejected"
	RequestStatusActive    ParkingRequestStatus = "active"
	RequestStatusCompleted ParkingRequestStatus = "completed"
)

// ParkingRequest represents a parking request.
type ParkingRequest struct {
	ID                      uuid.UUID `gorm:"primary_key"`
	UserID                  uuid.UUID
	ParkingSpaceID          *uuid.UUID // Can be nil, because the Admin chooses it after request is created.
	DestinationParkingLotID uuid.UUID
	StartTime               time.Time
	EndTime                 time.Time
	Status                  ParkingRequestStatus
}

// OnCreate sets requests's fields on create.
func (s *ParkingRequest) OnCreate() {
	s.ID = uuid.New()
	s.Status = RequestStatusPending
}

// OnArrivalNotification sets parking requests's status to 'active'.
func (s *ParkingRequest) OnArrivalNotification() {
	s.Status = RequestStatusActive
}

// OnDepartureNotification sets parking requests's status to 'completed'.
func (s *ParkingRequest) OnDepartureNotification() {
	s.Status = RequestStatusCompleted
}

// OnSpaceAssign approves and associates this parking request with the selected parking space.
func (s *ParkingRequest) OnSpaceAssign(spaceID uuid.UUID) {
	s.ParkingSpaceID = &spaceID
	s.Status = RequestStatusApproved
}
