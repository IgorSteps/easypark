package entities

import "github.com/google/uuid"

// AlertType is an enum for the types of alerts that can be created.
type AlertType int

const (
	// LocationMismatch alert type is when the location in the notification doesn't match actual parking space
	// location.
	LocationMismatch AlertType = iota
)

// Alert represents an alert that is sent to the admin.
type Alert struct {
	ID             uuid.UUID
	Type           AlertType
	Message        string
	UserID         uuid.UUID
	ParkingSpaceID uuid.UUID
}

func (s *Alert) OnLocationMismatchAlertCreate(msg string, driverID, spaceID uuid.UUID) {
	s.ID = uuid.New()
	s.Type = LocationMismatch
	s.Message = msg
	s.UserID = driverID
	s.ParkingSpaceID = spaceID
}
