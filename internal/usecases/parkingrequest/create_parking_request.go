package usecases

import (
	"context"
	"time"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/sirupsen/logrus"
)

// CreateParkingRequest provides business logic to create a parking request.
type CreateParkingRequest struct {
	logger             *logrus.Logger
	parkingRequestRepo repositories.ParkingRequestRepository
}

// NewCreateParkingRequest creates a new instance of the CreateParkingRequest.
func NewCreateParkingRequest(l *logrus.Logger, r repositories.ParkingRequestRepository) *CreateParkingRequest {
	return &CreateParkingRequest{
		logger:             l,
		parkingRequestRepo: r,
	}
}

// Execute runs the business logic to create a parking request.
func (s *CreateParkingRequest) Execute(ctx context.Context, parkingRequest *entities.ParkingRequest) (*entities.ParkingRequest, error) {
	err := s.validate(parkingRequest)
	if err != nil {
		return nil, err
	}

	parkingRequest.OnCreate()

	err = s.parkingRequestRepo.Create(ctx, parkingRequest)
	if err != nil {
		return nil, err
	}

	return parkingRequest, nil
}

// validate validates the parking request.
// TODO: check if the desired parking lot exists.
func (s *CreateParkingRequest) validate(parkingRequest *entities.ParkingRequest) error {
	if parkingRequest.StartTime.Before(time.Now()) {
		s.logger.WithFields(logrus.Fields{
			"start time": parkingRequest.StartTime,
			"now time":   time.Now(),
		}).Warn("start time cannot be in the past")
		return repositories.NewInvalidInputError("start time cannot be in the past")
	}
	if parkingRequest.StartTime.After(parkingRequest.EndTime) {
		s.logger.WithFields(logrus.Fields{
			"start time": parkingRequest.StartTime,
			"end time":   parkingRequest.EndTime,
		}).Warn("start time cannot be after the end time")
		return repositories.NewInvalidInputError("start time cannot be after the end time")
	}

	return nil
}
