package models

import (
	"time"

	"github.com/IgorSteps/easypark/internal/domain/entities"
)

type CreateParkingRequestRequest struct {
	Destination string    `json:"destination"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
}

func (s *CreateParkingRequestRequest) ToDomain() *entities.ParkingRequest {
	return &entities.ParkingRequest{
		Destination: s.Destination,
		StartTime:   s.StartTime,
		EndTime:     s.EndTime,
	}
}

type CreateParkingRequestResponse struct {
	Destination string                        `json:"destination"`
	StartTime   time.Time                     `json:"starttime"`
	EndTime     time.Time                     `json:"endtime"`
	Status      entities.ParkingRequestStatus `json:"status"`
}
