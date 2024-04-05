package models

import (
	"time"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/google/uuid"
)

// CreateParkingRequestRequest represent the data in an incoming HTTP request to create a parking request.
type CreateParkingRequestRequest struct {
	DestinationParkingLotID uuid.UUID `json:"destination"`
	StartTime               time.Time `json:"startTime"`
	EndTime                 time.Time `json:"endTime"`
}

// ToDomain converts CreateParkingRequestRequest into our domain type.
func (s *CreateParkingRequestRequest) ToDomain() *entities.ParkingRequest {
	return &entities.ParkingRequest{
		DestinationParkingLotID: s.DestinationParkingLotID,
		StartTime:               s.StartTime,
		EndTime:                 s.EndTime,
	}
}

// CreateParkingRequestResponse represent the data in an outgoing HTTP response toa  create parking request request.
type CreateParkingRequestResponse struct {
	ID          uuid.UUID                     `json:"id"`
	UserID      uuid.UUID                     `json:"userId"`
	Destination uuid.UUID                     `json:"destination"`
	StartTime   time.Time                     `json:"starttime"`
	EndTime     time.Time                     `json:"endtime"`
	Status      entities.ParkingRequestStatus `json:"status"`
}
