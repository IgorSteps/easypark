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

func (s *UpdateParkingRequestStatus) Execute(ctx context.Context, id uuid.UUID, status entities.ParkingRequestStatus) error {
	parkingRequest, err := s.repo.GetParkingRequestByID(ctx, id)
	if err != nil {
		return err
	}

	parkingRequest.Status = status

	err = s.repo.Save(ctx, &parkingRequest)
	if err != nil {
		return err
	}

	return nil
}
