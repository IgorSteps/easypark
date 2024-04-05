package models

import "github.com/google/uuid"

// ParkingRequestSpaceUpdateRequest represents an incoming HTTP request to update status of a parking request.
type ParkingRequestSpaceUpdateRequest struct {
	ParkingSpaceID uuid.UUID `json:"parkingSpaceID"`
}

// ParkingRequestSpaceUpdateRequest represents a response to updating status of a parking request.
type ParkingRequestSpaceUpdateResponse struct {
	Message string `json:"message"`
}
