package models

import "github.com/google/uuid"

// ParkingRequestSpaceUpdateRequest represents an incoming HTTP request to assign a parking space to parking request.
type ParkingRequestSpaceUpdateRequest struct {
	ParkingSpaceID uuid.UUID `json:"parkingSpaceID"`
}

// ParkingRequestSpaceUpdateRequest represents an incoming HTTP request to automatically assign a parking space to parking request.
type ParkingRequestAutomaticSpaceUpdateRequest struct {
	ParkingRequestID uuid.UUID `json:"parkingRequestID"`
}

// ParkingRequestSpaceUpdateRequest represents a response to updating status of a parking request.
type ParkingRequestSpaceUpdateResponse struct {
	Message string `json:"message"`
}
