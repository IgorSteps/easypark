package datastore

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Error code to signify that a unique contraint has been violated.
const ErrCodeUniqueViolation = "23505"

// ParkingLotPostgresRepository implements PakringLotRepository interface to provide database operation on parking lots.
type ParkingLotPostgresRepository struct {
	Logger *logrus.Logger
	DB     Datastore
}

// NewParkingParkingLotPostgresRepository returns a new instance of ParkingLotPostgresRepository.
func NewParkingParkingLotPostgresRepository(l *logrus.Logger, db Datastore) *ParkingLotPostgresRepository {
	return &ParkingLotPostgresRepository{
		Logger: l,
		DB:     db,
	}
}

// CreateParkingLot creates a parking lot in the database.
func (s *ParkingLotPostgresRepository) CreateParkingLot(ctx context.Context, parkingLot *entities.ParkingLot) error {
	result := s.DB.WithContext(ctx).Create(parkingLot)
	err := result.Error()
	if err != nil {
		if pgError, ok := err.(*pgconn.PgError); ok {
			if pgError.Code == ErrCodeUniqueViolation {
				s.Logger.WithField("name", parkingLot.Name).WithError(err).Warn("parking lot with this name already exists")
				return repositories.NewResourceAlreadyExistsError(parkingLot.Name)
			}
		}

		s.Logger.WithError(err).Error("failed to insert parking lot into the database")
		return repositories.NewInternalError("failed to insert parking lot into the database")
	}

	return nil
}

func (s *ParkingLotPostgresRepository) GetAllParkingLots(ctx context.Context) ([]entities.ParkingLot, error) {
	var lots []entities.ParkingLot

	result := s.DB.WithContext(ctx).Preload("ParkingSpaces").FindAll(&lots)
	err := result.Error()
	if err != nil {
		s.Logger.WithError(err).Error("failed to query for all parking lots in the database")
		return []entities.ParkingLot{}, repositories.NewInternalError("failed to query for all parking lots in the database")
	}

	return lots, nil
}

// GetSingle gets a single parking lot using its ID.
func (s *ParkingLotPostgresRepository) GetSingle(ctx context.Context, id uuid.UUID) (*entities.ParkingLot, error) {
	var parkingLot entities.ParkingLot

	result := s.DB.WithContext(ctx).Preload("ParkingSpaces.ParkingRequests").First(&parkingLot, "id = ?", id)
	err := result.Error()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			s.Logger.WithField("parking lot id", id).Error("failed to find parking lot with given id in the database")
			return nil, repositories.NewNotFoundError(id.String())
		}

		s.Logger.WithError(err).Error("failed to query for parking lot in the database")
		return nil, repositories.NewInternalError("failed to query for parking lot in the database")
	}

	return &parkingLot, nil
}

func (s *ParkingLotPostgresRepository) DeleteParkingLot(ctx context.Context, id uuid.UUID) error {
	result := s.DB.WithContext(ctx).Delete(&entities.ParkingLot{}, id)

	err := result.Error()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			s.Logger.WithError(err).WithField("id", id).Error("parking lot not found")
			return repositories.NewInvalidInputError("parking lot not found")
		}

		s.Logger.WithError(err).WithField("id", id).Error("failed to delete a parking lot")
		return repositories.NewInternalError("failed to delete a parking lot")
	}
	return nil
}
