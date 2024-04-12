package usecases

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type GetSingleParkingSpace struct {
	logger *logrus.Logger
	repo   repositories.ParkingSpaceRepository
}

func NewGetSingleParkingSpace(l *logrus.Logger, r repositories.ParkingSpaceRepository) *GetSingleParkingSpace {
	return &GetSingleParkingSpace{
		logger: l,
		repo:   r,
	}
}

func (s *GetSingleParkingSpace) Execute(ctx context.Context, parkingSpaceID uuid.UUID) (entities.ParkingSpace, error) {
	return s.repo.GetParkingSpaceByID(ctx, parkingSpaceID)
}
