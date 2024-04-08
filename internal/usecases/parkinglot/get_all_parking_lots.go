package usecases

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/sirupsen/logrus"
)

// GetAllParkingLots provides business logic to get all parking lots.
type GetAllParkingLots struct {
	Logger *logrus.Logger
	Repo   repositories.ParkingLotRepository
}

// NewGetAllParkingLots returns a new instance of GetAllParkingLots.
func NewGetAllParkingLots(l *logrus.Logger, r repositories.ParkingLotRepository) *GetAllParkingLots {
	return &GetAllParkingLots{
		Logger: l,
		Repo:   r,
	}
}

// Execute runs the business logic to get all parking lots.
func (s *GetAllParkingLots) Execute(ctx context.Context) ([]entities.ParkingLot, error) {
	lots, err := s.Repo.GetAllParkingLots(ctx)
	if err != nil {
		return nil, err
	}

	return lots, nil
}
