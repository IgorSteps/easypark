package entities

import (
	"time"

	"github.com/google/uuid"
)

// NotificationType defines the type of notification being sent.
type NotificationType int

const (
	// notifications:
	ArrivalNotification NotificationType = iota
	DepartureNotification
	// alerts:
	LocationMismatchNotification
	ArrivalDelayNotification
	DepartureDelayNotification
)

// Notification represents a notification.
type Notification struct {
	ID             uuid.UUID
	Type           NotificationType
	DriverID       uuid.UUID
	ParkingSpaceID uuid.UUID
	Location       string
	Timestamp      time.Time
	Message        string
}

func (s *Notification) OnCreate(driverID, spaceID uuid.UUID, location string, notificationType NotificationType) *Notification {
	return &Notification{
		ID:             uuid.New(),
		Type:           notificationType,
		DriverID:       driverID,
		ParkingSpaceID: spaceID,
		Location:       location,
		Timestamp:      time.Now(),
	}
}
