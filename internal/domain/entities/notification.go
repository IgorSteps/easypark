package entities

import (
	"time"

	"github.com/google/uuid"
)

// NotificationType defines the type of notification being sent.
type NotificationType int

const (
	ArrivalNotification NotificationType = iota
	DepartureNotification
	GPSCoordinateMismatchNotification
	ArrivalDelayNotification
	DepartureDelayNotification
)

// Notification represents a notification.
type Notification struct {
	ID             uuid.UUID
	DriverID       uuid.UUID
	ParkingSpaceID uuid.UUID
	Location       string
	Timestamp      time.Time
	Message        string
}
