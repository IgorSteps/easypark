package models

import (
	"time"

	"github.com/IgorSteps/easypark/internal/domain/entities"
)

// CreateParkingRequestRequest represent the data in an incoming HTTP request to create a parking request.
type CreateParkingRequestRequest struct {
	Destination string    `json:"destination"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
}

// ToDomain converts CreateParkingRequestRequest into our domain type.
func (s *CreateParkingRequestRequest) ToDomain() *entities.ParkingRequest {
	return &entities.ParkingRequest{
		Destination: s.Destination,
		StartTime:   s.StartTime,
		EndTime:     s.EndTime,
	}
}

// CreateParkingRequestResponse represent the data in an outgoing HTTP response toa  create parking request request.
type CreateParkingRequestResponse struct {
	Destination string                        `json:"destination"`
	StartTime   time.Time                     `json:"starttime"`
	EndTime     time.Time                     `json:"endtime"`
	Status      entities.ParkingRequestStatus `json:"status"`
}
