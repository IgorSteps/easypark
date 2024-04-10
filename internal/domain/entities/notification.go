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
)

// Notification represents a notification.
type Notification struct {
	ID             uuid.UUID `gorm:"primary_key"`
	Type           NotificationType
	DriverID       uuid.UUID
	ParkingSpaceID uuid.UUID
	Location       string
	Timestamp      time.Time
}

func (s *Notification) OnCreate(driverID, spaceID uuid.UUID, location string, notificationType NotificationType) {
	s.ID = uuid.New()
	s.Type = notificationType
	s.DriverID = driverID
	s.ParkingSpaceID = spaceID
	s.Location = location
	s.Timestamp = time.Now()
}
