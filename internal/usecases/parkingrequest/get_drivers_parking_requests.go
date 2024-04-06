package usecases

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// GetDriversParkingRequests provides business logic to get all parking requests for a particular driver.
// TODO: Ideally this would be refactored to get filtered parking requests.
type GetDriversParkingRequests struct {
	logger *logrus.Logger
	repo   repositories.ParkingRequestRepository
}

// NewGetDriversParkingRequests returns new instance of GetDriversParkingRequests.
func NewGetDriversParkingRequests(l *logrus.Logger, r repositories.ParkingRequestRepository) *GetDriversParkingRequests {
	return &GetDriversParkingRequests{
		logger: l,
		repo:   r,
	}
}

// Execute runs the business logic to get parking requests for a particular driver.
func (s *GetDriversParkingRequests) Execute(ctx context.Context, driverID uuid.UUID) ([]entities.ParkingRequest, error) {
	parkingRequests, err := s.repo.GetAllParkingRequestsForUser(ctx, driverID)
	if err != nil {
		return []entities.ParkingRequest{}, err
	}

	return parkingRequests, nil
}
