package usecases

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// UpdateParkingRequestSpace provides business logic to update a space in a parking request.
type UpdateParkingRequestSpace struct {
	logger             *logrus.Logger
	parkingRequestRepo repositories.ParkingRequestRepository
	parkingSpaceRepo   repositories.ParkingSpaceRepository
}

// NewUpdateParkingRequestSpace returns new instance of UpdateParkingRequestSpace.
func NewUpdateParkingRequestSpace(
	l *logrus.Logger,
	reqRepo repositories.ParkingRequestRepository,
	spaceRepo repositories.ParkingSpaceRepository,
) *UpdateParkingRequestSpace {
	return &UpdateParkingRequestSpace{
		logger:             l,
		parkingRequestRepo: reqRepo,
		parkingSpaceRepo:   spaceRepo,
	}
}

// Execute runs the business logic to assign a space to a parking request.
func (s *UpdateParkingRequestSpace) Execute(ctx context.Context, requestID uuid.UUID, parkingSpaceID uuid.UUID) error {
	parkingSpace, err := s.parkingSpaceRepo.GetParkingSpaceByID(ctx, parkingSpaceID)
	if err != nil {
		return err
	}

	parkingRequest, err := s.parkingRequestRepo.GetParkingRequestByID(ctx, requestID)
	if err != nil {
		return err
	}

	// Check if the parking space belongs in the user chosen destination(parking lot).
	if parkingRequest.DestinationParkingLotID != parkingSpace.ParkingLotID {
		s.logger.WithFields(logrus.Fields{
			"desired": parkingRequest.DestinationParkingLotID,
			"actual":  parkingSpace.ParkingLotID,
		}).Warn("parking space is not in the desired parking lot")
		return repositories.NewInvalidInputError("parking space is not in the desired parking lot")
	}

	// Check if the space is available.
	if parkingSpace.Status != entities.StatusAvailable {
		s.logger.WithField("status", parkingSpace.Status).Warn("parking space isn't available")
		return repositories.NewInvalidInputError("parking space isn't available")
	}

	// Update our parking request.
	parkingRequest.ParkingSpaceID = &parkingSpace.ID

	// Update and save parking space.
	parkingSpace.OnAssign(parkingRequest.StartTime, parkingRequest.EndTime, parkingRequest.UserID)
	err = s.parkingSpaceRepo.Save(ctx, &parkingSpace)
	if err != nil {
		return err
	}

	// Save udpated parking request
	err = s.parkingRequestRepo.Save(ctx, &parkingRequest)
	if err != nil {
		return err
	}

	return nil
}
