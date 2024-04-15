package usecases

import (
	"context"
	"time"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// AssignParkingSpace provides business logic to assign a parking space to a parking request.
type AssignParkingSpace struct {
	logger             *logrus.Logger
	parkingRequestRepo repositories.ParkingRequestRepository
	parkingSpaceRepo   repositories.ParkingSpaceRepository
}

// NewAssignParkingSpace returns new instance of AssignParkingSpace.
func NewAssignParkingSpace(
	l *logrus.Logger,
	reqRepo repositories.ParkingRequestRepository,
	spaceRepo repositories.ParkingSpaceRepository,
) *AssignParkingSpace {
	return &AssignParkingSpace{
		logger:             l,
		parkingRequestRepo: reqRepo,
		parkingSpaceRepo:   spaceRepo,
	}
}

// Execute runs the business logic to assign a parking space to a parking request.
func (s *AssignParkingSpace) Execute(ctx context.Context, requestID uuid.UUID, parkingSpaceID uuid.UUID) error {
	parkingSpace, err := s.parkingSpaceRepo.GetSingle(ctx, parkingSpaceID)
	if err != nil {
		return err
	}

	parkingRequest, err := s.parkingRequestRepo.GetSingle(ctx, requestID)
	if err != nil {
		return err
	}

	// Check if the parking request has passed its desired start time.
	if parkingRequest.StartTime.Before(time.Now()) {
		s.logger.Warn("not allowed to assign a parking space to a parking request with the desired start time in the past")
		return repositories.NewInvalidInputError("not allowed to assign a parking space to a parking request with the desired start time in the past")

	}

	// Check if the status of parking request is 'rejected'.
	if parkingRequest.Status == entities.RequestStatusRejected {
		s.logger.Warn("not allowed to assign parking space to a 'rejected' parking request")
		return repositories.NewInvalidInputError("not allowed to assign parking space to a 'rejected' parking request")
	}

	// Check if the parking space belongs in the driver's chosen destination(parking lot).
	if parkingRequest.DestinationParkingLotID != parkingSpace.ParkingLotID {
		s.logger.WithFields(logrus.Fields{
			"desired": parkingRequest.DestinationParkingLotID,
			"actual":  parkingSpace.ParkingLotID,
		}).Warn("parking space is not in the desired parking lot")
		return repositories.NewInvalidInputError("parking space is not in the desired parking lot")
	}

	// Check the parking space status and act accordingly.
	switch parkingSpace.Status {
	case entities.ParkingSpaceStatusBlocked:
		s.logger.Warn("not allowed to assign blocked parking space")
		return repositories.NewInvalidInputError("not allowed to assign blocked parking space")
	case entities.ParkingSpaceStatusReserved:
		// TODO: Ideally, we would allow to have reservations for parking spaces like we do with parking requests.
		// but there a lot of considerations...
		s.logger.Warn("not allowed to assign reserved parking space")
		return repositories.NewInvalidInputError("not allowed to assign reserved parking space")
	case entities.ParkingSpaceStatusOccupied: // If space is occupied now, it can be available in the future, so still check for overlap.
	case entities.ParkingSpaceStatusAvailable:
		if parkingSpace.CheckForOverlap(parkingRequest.StartTime, parkingRequest.EndTime) {
			s.logger.Warn("there is an overlap with existing parking requests time slots")
			return repositories.NewInvalidInputError("there is an overlap with existing parking requests time slots")
		}

		s.logger.Debug("no overlap with existing parking requests, continuing...")
	}

	// All checks have passed, now we can approve this parking request and
	// associate it with the selected parking space.
	parkingRequest.OnSpaceAssign(parkingSpaceID)

	// Save updated parking request.
	err = s.parkingRequestRepo.Save(ctx, &parkingRequest)
	if err != nil {
		return err
	}

	return nil
}
