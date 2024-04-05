package datastore

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// ParkingSpacePostgresRepository implements ParkingSpaceRepository interface to provide database operation on ParkingSpaces.
type ParkingSpacePostgresRepository struct {
	Logger *logrus.Logger
	DB     Datastore
}

// NewParkingSpacePostgresRepository returns a new instance of ParkingSpacePostgresRepository.
func NewParkingSpacePostgresRepository(l *logrus.Logger, db Datastore) *ParkingSpacePostgresRepository {
	return &ParkingSpacePostgresRepository{
		Logger: l,
		DB:     db,
	}
}

// GetParkingSpaceByID gets a parking space with a given ID.
func (s *ParkingSpacePostgresRepository) GetParkingSpaceByID(ctx context.Context, id uuid.UUID) (entities.ParkingSpace, error) {
	var space entities.ParkingSpace

	result := s.DB.WithContext(ctx).First(&space, "id = ?", id)
	err := result.Error()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			s.Logger.WithField("id", id).Error("failed to find the parking space with given id in the database")
			return entities.ParkingSpace{}, repositories.NewNotFoundError(id.String())
		}

		s.Logger.WithError(err).Error("failed to query for parking space in the database")
		return entities.ParkingSpace{}, repositories.NewInternalError("failed to query for parking space in the database")
	}

	return space, nil
}

// Saves saves an updated parking space into the DB.
func (s *ParkingSpacePostgresRepository) Save(ctx context.Context, space *entities.ParkingSpace) error {
	result := s.DB.WithContext(ctx).Save(space)
	err := result.Error()
	if err != nil {
		s.Logger.WithError(err).Error("failed to save updated parking space in the database")
		return repositories.NewInternalError("failed to save updated parking space in the database")
	}

	return nil
}
