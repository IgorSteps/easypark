package usecases

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// DeassignParkingSpace provides business logic to deassign parking spaces.
type DeassignParkingSpace struct {
	logger             *logrus.Logger
	parkingRequestRepo repositories.ParkingRequestRepository
}

// NewDeassignParkingSpace returns new instance of DeassignParkingSpace.
func NewDeassignParkingSpace(l *logrus.Logger, reqRepo repositories.ParkingRequestRepository) *DeassignParkingSpace {
	return &DeassignParkingSpace{
		logger:             l,
		parkingRequestRepo: reqRepo,
	}
}

// Execute runs the business logic.
func (s *DeassignParkingSpace) Execute(ctx context.Context, requestID uuid.UUID) error {
	parkingRequest, err := s.parkingRequestRepo.GetSingle(ctx, requestID)
	if err != nil {
		return err
	}
	if parkingRequest.Status != entities.RequestStatusApproved {
		s.logger.Error("cannot deassign parking space from a parking request that is not approved")
		return repositories.NewInvalidInputError("cannot deassign parking space from a parking request that is not approved")
	}

	parkingRequest.OnSpaceDeassign()

	err = s.parkingRequestRepo.Save(ctx, &parkingRequest)
	if err != nil {
		return err
	}

	return nil
}
