package usecases

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/sirupsen/logrus"
)

// GetAllParkingRequests provides business logic to get all parking requests.
// TODO: Ideally this would be refactored to get filtered pakring requests.
type GetAllParkingRequests struct {
	logger *logrus.Logger
	repo   repositories.ParkingRequestRepository
}

// NewGetAllParkingRequests returns new instance of GetAllParkingRequests.
func NewGetAllParkingRequests(l *logrus.Logger, r repositories.ParkingRequestRepository) *GetAllParkingRequests {
	return &GetAllParkingRequests{
		logger: l,
		repo:   r,
	}
}

// Execute runs the business logic to get all parking requests.
func (s *GetAllParkingRequests) Execute(ctx context.Context) ([]entities.ParkingRequest, error) {
	parkingRequests, err := s.repo.GetAllParkingRequests(ctx)
	if err != nil {
		return []entities.ParkingRequest{}, err
	}

	return parkingRequests, nil
}
