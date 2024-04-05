package usecases

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/sirupsen/logrus"
)

// CreateParkingLot provides business logic to create a parking lot.
type CreateParkingLot struct {
	logger *logrus.Logger
	repo   repositories.ParkingLotRepository
}

// NewCreateParkingLot returns a new instance of CreateParkingLot.
func NewCreateParkingLot(l *logrus.Logger, r repositories.ParkingLotRepository) *CreateParkingLot {
	return &CreateParkingLot{
		logger: l,
		repo:   r,
	}
}

// Execute runs the business logic to create a parking lot.
func (s *CreateParkingLot) Execute(ctx context.Context, name string, capacity int) (entities.ParkingLot, error) {
	var lot entities.ParkingLot
	lot.OnCreate(name, capacity)

	err := s.repo.CreateParkingLot(ctx, &lot)
	if err != nil {
		return entities.ParkingLot{}, err
	}

	return lot, nil
}
