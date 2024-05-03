package entities

import "github.com/google/uuid"

// AlertType is an enum for the types of alerts that can be created.
type AlertType int

const (
	// LocationMismatch alert type is when the location in the notification doesn't match actual parking space
	// location.
	LocationMismatch AlertType = iota

	// LateArrival alert type is when an arrival notification hasn't been received within one hour
	// from the requests start time.
	LateArrival

	// OverStay alert type is when an exit notification hasn't been received after 30 minutes
	// from the requests end time.
	OverStay
)

// Alert represents an alert that is sent to the admin.
type Alert struct {
	ID             uuid.UUID
	Type           AlertType
	Message        string
	UserID         uuid.UUID
	ParkingSpaceID uuid.UUID
}

func (s *Alert) CreateLocationMismatchAlert(msg string, driverID, spaceID uuid.UUID) {
	s.ID = uuid.New()
	s.Type = LocationMismatch
	s.Message = msg
	s.UserID = driverID
	s.ParkingSpaceID = spaceID
}

func (s *Alert) CreateLateArrivalAlert(msg string, driverID, spaceID uuid.UUID) {
	s.ID = uuid.New()
	s.Type = LateArrival
	s.Message = msg
	s.UserID = driverID
	s.ParkingSpaceID = spaceID
}

func (s *Alert) CreateOverStayAlert(msg string, driverID, spaceID uuid.UUID) {
	s.ID = uuid.New()
	s.Type = OverStay
	s.Message = msg
	s.UserID = driverID
	s.ParkingSpaceID = spaceID
}
