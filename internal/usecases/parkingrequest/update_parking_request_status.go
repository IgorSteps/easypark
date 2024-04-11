package usecases

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type UpdateParkingRequestStatus struct {
	logger *logrus.Logger
	repo   repositories.ParkingRequestRepository
}

func NewUpdateParkingRequestStatus(lgr *logrus.Logger, r repositories.ParkingRequestRepository) *UpdateParkingRequestStatus {
	return &UpdateParkingRequestStatus{
		logger: lgr,
		repo:   r,
	}
}

func (s *UpdateParkingRequestStatus) Execute(ctx context.Context, id uuid.UUID, status string) error {
	domainStatus, err := parkingRequestFromString(status)
	if err != nil {
		s.logger.WithField("status", status).WithError(err).Error("unknown parking request status")
		return err
	}

	parkingRequest, err := s.repo.GetParkingRequestByID(ctx, id)
	if err != nil {
		return err
	}

	// We don't allow updating status of approved requests, because they have a parking space assigned.
	if parkingRequest.Status == entities.RequestStatusApproved && (domainStatus == entities.RequestStatusRejected || domainStatus == entities.RequestStatusPending) {
		s.logger.Warn("not allowed to change status of an approved request to 'rejected' or 'pending'")
		return repositories.NewInvalidInputError("not allowed to change approved request status to reject or pending")
	}

	// Udpate status.
	parkingRequest.Status = domainStatus

	err = s.repo.Save(ctx, &parkingRequest)
	if err != nil {
		return err
	}

	return nil
}

func parkingRequestFromString(status string) (entities.ParkingRequestStatus, error) {
	switch status {
	case "approved":
		return "", repositories.NewInvalidInputError("not allowed to directly approve parking requests")
	case "rejected":
		return entities.RequestStatusRejected, nil
	case "pending":
		return entities.RequestStatusRejected, nil
	default:
		return "", repositories.NewInvalidInputError("unknown parking request status")
	}
}
