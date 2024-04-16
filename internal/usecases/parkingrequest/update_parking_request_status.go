package usecases

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// UpdateParkingRequestStatus provides business logic to update the status of a parking request to
// 'rejected' or 'pending'
type UpdateParkingRequestStatus struct {
	logger *logrus.Logger
	repo   repositories.ParkingRequestRepository
}

// NewUpdateParkingRequestStatus returns a new instance of UpdateParkingRequestStatus.
func NewUpdateParkingRequestStatus(lgr *logrus.Logger, r repositories.ParkingRequestRepository) *UpdateParkingRequestStatus {
	return &UpdateParkingRequestStatus{
		logger: lgr,
		repo:   r,
	}
}

// Execute runs the business logic.
func (s *UpdateParkingRequestStatus) Execute(ctx context.Context, id uuid.UUID, status string) error {
	domainStatus, err := parkingRequestFromString(status)
	if err != nil {
		s.logger.WithField("status", status).WithError(err).Error("unknown or not allowed parking request status")
		return err
	}

	parkingRequest, err := s.repo.GetSingle(ctx, id)
	if err != nil {
		return err
	}

	parkingRequest.Status = domainStatus

	err = s.repo.Save(ctx, &parkingRequest)
	if err != nil {
		return err
	}

	return nil
}

func parkingRequestFromString(status string) (entities.ParkingRequestStatus, error) {
	switch status {
	case "rejected":
		return entities.RequestStatusRejected, nil
	case "pending":
		return entities.RequestStatusRejected, nil
	default:
		return "", repositories.NewInvalidInputError("unknown or not allowed parking request status")
	}
}
