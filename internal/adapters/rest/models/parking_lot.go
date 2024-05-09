package models

import (
	"time"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/google/uuid"
)

// CreateParkingLotRequest represents an incoming HTTP request body to create a parking lot.
type CreateParkingLotRequest struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
}

// CreateParkingLotResponse represents a response to creating a parking lot.
type CreateParkingLotResponse struct {
	ID            uuid.UUID               `json:"id"`
	Name          string                  `json:"name"`
	Capacity      int                     `json:"capacity"`
	PakringSpaces []entities.ParkingSpace `json:"parkingSpaces"`
	CreatedAt     time.Time               `json:"createdAt"`
	UpdatedAt     time.Time               `json:"updatedAt"`
}
