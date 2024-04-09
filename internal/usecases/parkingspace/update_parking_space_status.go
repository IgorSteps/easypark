package usecases

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// UpdateParkingSpaceStatus provides business logic to update a parking space status.
type UpdateParkingSpaceStatus struct {
	logger *logrus.Logger
	repo   repositories.ParkingSpaceRepository
}

// NewUpdateParkingSpaceStatus returns a new instance of the UpdateParkingSpaceStatus.
func NewUpdateParkingSpaceStatus(l *logrus.Logger, r repositories.ParkingSpaceRepository) *UpdateParkingSpaceStatus {
	return &UpdateParkingSpaceStatus{
		logger: l,
		repo:   r,
	}
}

// Execute runs the business logic.
func (s *UpdateParkingSpaceStatus) Execute(ctx context.Context, id uuid.UUID, status string) (entities.ParkingSpace, error) {
	domainStatus, err := parseSpaceStatus(status)
	if err != nil {
		s.logger.WithField("status", status).WithError(err).Error("failed to parse given status")
		return entities.ParkingSpace{}, err
	}

	parkSpace, err := s.repo.GetParkingSpaceByID(ctx, id)
	if err != nil {
		return entities.ParkingSpace{}, err
	}

	// Update status.
	parkSpace.Status = domainStatus

	err = s.repo.Save(ctx, &parkSpace)
	if err != nil {
		return entities.ParkingSpace{}, err
	}

	return parkSpace, nil
}

func parseSpaceStatus(status string) (entities.ParkingSpaceStatus, error) {
	switch status {
	case "available":
		return entities.StatusAvailable, nil
	case "occupied":
		return entities.StatusOccupied, nil
	case "blocked":
		return entities.StatusBlocked, nil
	case "reserved":
		return entities.StatusReserved, nil
	default:
		return "", repositories.NewInvalidInputError("failed to parse given status")
	}
}
