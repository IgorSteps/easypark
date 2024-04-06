package usecases

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/sirupsen/logrus"
)

type GetAllParkingRequests struct {
	logger *logrus.Logger
	repo   repositories.ParkingRequestRepository
}

func NewGetAllParkingRequests(l *logrus.Logger, r repositories.ParkingRequestRepository) *GetAllParkingRequests {
	return &GetAllParkingRequests{
		logger: l,
		repo:   r,
	}
}

func (s *GetAllParkingRequests) Execute(ctx context.Context) ([]entities.ParkingRequest, error) {
	parkingRequests, err := s.repo.GetAllParkingRequests(ctx)
	if err != nil {
		return []entities.ParkingRequest{}, err
	}

	return parkingRequests, nil
}
