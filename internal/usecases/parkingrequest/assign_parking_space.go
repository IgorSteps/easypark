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

	// Check if the parking request is already rejected.
	if parkingRequest.Status == entities.RequestStatusRejected {
		s.logger.Warn("rejected parking request cannot be assigned a space")
		return repositories.NewInvalidInputError("rejected parking request cannot be assigned a space")
	}

	// Check if the parking space belongs in the user chosen destination(parking lot).
	if parkingRequest.DestinationParkingLotID != parkingSpace.ParkingLotID {
		s.logger.WithFields(logrus.Fields{
			"desired": parkingRequest.DestinationParkingLotID,
			"actual":  parkingSpace.ParkingLotID,
		}).Warn("parking space is not in the desired parking lot")
		return repositories.NewInvalidInputError("parking space is not in the desired parking lot")
	}

	// Check if space is blocked.
	if parkingSpace.Status == entities.StatusBlocked {
		s.logger.Warn("not allowed to assign blocked parking spaces")
		return repositories.NewInvalidInputError("not allowed to assign a blocked parking space")
	}

	// When parking request is assigned a place, it's status is changed to 'reserved'
	// When driver sends arrival notification, it's status is changed to 'occupied'
	// When driver sends departure notification, it's status is changed to either 'available' or
	// 'reserved' given that there might be other bookins made.
	// If status is 'reserved' or 'occupied', check if the space is available at the
	// requested time(because times might not overlap).
	if parkingSpace.Status == entities.StatusOccupied || parkingSpace.Status == entities.StatusReserved {
		if !parkingSpace.IsAvailableFor(parkingRequest.StartTime, parkingRequest.EndTime) {
			s.logger.Warn("requested parking space is not available due to an overlap with existing bookings")
			return repositories.NewInvalidInputError("parking space isn't available at the requested time")
		}

		s.logger.Debug("parking space can be assigned, because there is no overlap with existing bookings")
	}

	// Update our parking request.
	parkingRequest.ParkingSpaceID = &parkingSpace.ID
	// We assume that if a parking request is assigned a space, it becomes approved.
	parkingRequest.Status = entities.RequestStatusApproved

	// Update and save parking space.
	parkingSpace.OnAssign(parkingRequest.StartTime, parkingRequest.EndTime)
	err = s.parkingSpaceRepo.Save(ctx, &parkingSpace)
	if err != nil {
		return err
	}

	// Save updated parking request
	err = s.parkingRequestRepo.Save(ctx, &parkingRequest)
	if err != nil {
		return err
	}

	return nil
}
