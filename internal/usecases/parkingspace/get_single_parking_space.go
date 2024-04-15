package usecases

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// GetSingleParkingSpace provides business logic to get a single parking space.
type GetSingleParkingSpace struct {
	logger *logrus.Logger
	repo   repositories.ParkingSpaceRepository
}

// NewGetSingleParkingSpace returns a new instance of GetSingleParkingSpace.
func NewGetSingleParkingSpace(l *logrus.Logger, r repositories.ParkingSpaceRepository) *GetSingleParkingSpace {
	return &GetSingleParkingSpace{
		logger: l,
		repo:   r,
	}
}

// Execute runs the business logic.
func (s *GetSingleParkingSpace) Execute(ctx context.Context, parkingSpaceID uuid.UUID) (entities.ParkingSpace, error) {
	return s.repo.GetSingle(ctx, parkingSpaceID)
}
