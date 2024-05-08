package usecases

import (
	"context"
	"time"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// AssignParkingSpace provides business logic to automatically assign a parking space to a parking request.
type AutomaticAssignParkingSpace struct {
	logger             *logrus.Logger
	parkingRequestRepo repositories.ParkingRequestRepository
	parkingSpaceRepo   repositories.ParkingSpaceRepository
}

// NewAssignParkingSpace returns new instance of AutomaticAssignParkingSpace.
func NewAutomaticAssignParkingSpace(
	l *logrus.Logger,
	reqRepo repositories.ParkingRequestRepository,
	spaceRepo repositories.ParkingSpaceRepository,
) *AutomaticAssignParkingSpace {
	return &AutomaticAssignParkingSpace{
		logger:             l,
		parkingRequestRepo: reqRepo,
		parkingSpaceRepo:   spaceRepo,
	}
}

// Execute runs the business logic to assign a parking space to a parking request.
func (s *AutomaticAssignParkingSpace) Execute(ctx context.Context, requestID uuid.UUID) (*entities.ParkingSpace, error) {
	parkingRequest, err := s.parkingRequestRepo.GetSingle(ctx, requestID)
	if err != nil {
		return nil, err
	}

	// Check if the parking request has passed its desired start time.
	if parkingRequest.StartTime.Before(time.Now()) {
		s.logger.Warn("not allowed to assign a parking space to a parking request with the desired start time in the past")
		return nil, repositories.NewInvalidInputError("not allowed to assign a parking space to a parking request with the desired start time in the past")

	}

	// Check if the status of parking request is 'rejected'.
	if parkingRequest.Status == entities.RequestStatusRejected {
		s.logger.Warn("not allowed to assign parking space to a 'rejected' parking request")
		return nil, repositories.NewInvalidInputError("not allowed to assign parking space to a 'rejected' parking request")
	}

	availableSpaces, err := s.parkingSpaceRepo.FindAvailableSpaces(ctx, parkingRequest.DestinationParkingLotID, parkingRequest.StartTime, parkingRequest.EndTime)
	if err != nil {
		return nil, err
	}

	if len(availableSpaces) == 0 {
		return nil, repositories.NewInvalidInputError("no available parking spaces at the desired time")
	}

	// Select the first available space
	// Can be extended to choose using some criteria...
	selectedSpace := availableSpaces[0]

	// Assign to parking request.
	parkingRequest.OnSpaceAssign(selectedSpace.ID)

	// Save updated parking request.
	err = s.parkingRequestRepo.Save(ctx, &parkingRequest)
	if err != nil {
		return nil, err
	}

	return &selectedSpace, nil
}
