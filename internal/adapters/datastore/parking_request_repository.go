package datastore

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// ParkingRequestPostgresRepository implements ParkingRequestPostgresRepository interface to provide database operation on Parking Lots.
type ParkingRequestPostgresRepository struct {
	Logger *logrus.Logger
	DB     Datastore
}

// NewParkingRequestPostgresRepository retuns new ParkingRequestPostgresRepository.
func NewParkingRequestPostgresRepository(db Datastore, lgr *logrus.Logger) *ParkingRequestPostgresRepository {
	return &ParkingRequestPostgresRepository{
		Logger: lgr,
		DB:     db,
	}
}

func (s *ParkingRequestPostgresRepository) CreateParkingRequest(ctx context.Context, parkReq *entities.ParkingRequest) error {
	result := s.DB.WithContext(ctx).Create(parkReq)
	err := result.Error()
	if err != nil {
		s.Logger.WithError(err).Error("failed to insert parking request into the database")
		return repositories.NewInternalError("failed to insert parking request into the database")
	}

	return nil
}

func (s *ParkingRequestPostgresRepository) GetAllParkingRequests(ctx context.Context) ([]entities.ParkingRequest, error) {
	var requests []entities.ParkingRequest

	result := s.DB.WithContext(ctx).FindAll(&requests)
	err := result.Error()
	if err != nil {
		s.Logger.WithError(err).Error("failed to query for all parking requests in the database")
		return []entities.ParkingRequest{}, repositories.NewInternalError("failed to query for all parking requests in the database")
	}

	return requests, nil
}

func (s *ParkingRequestPostgresRepository) GetAllParkingRequestsForUser(ctx context.Context, userID uuid.UUID) ([]entities.ParkingRequest, error) {
	var requests []entities.ParkingRequest

	result := s.DB.WithContext(ctx).Where("user_id = ?", userID).FindAll(&requests)
	err := result.Error()
	if err != nil {
		s.Logger.WithError(err).WithField("userID", userID).Error("failed to query for all parking requests in the database for particular user")
		return []entities.ParkingRequest{}, repositories.NewInternalError("failed to query for all parking requests in the database for particular user")
	}

	return requests, nil
}

func (s *ParkingRequestPostgresRepository) GetParkingRequestByID(ctx context.Context, id uuid.UUID) (entities.ParkingRequest, error) {
	var parkingRequest entities.ParkingRequest

	result := s.DB.WithContext(ctx).First(&parkingRequest, "id = ?", id)
	err := result.Error()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			s.Logger.WithField("parking request id", id).Error("failed to find parking request with given id in the database")
			return entities.ParkingRequest{}, repositories.NewNotFoundError(id.String())
		}

		s.Logger.WithError(err).Error("failed to query for parking request in the database")
		return entities.ParkingRequest{}, repositories.NewInternalError("failed to query for parking request in the database")
	}

	return parkingRequest, nil
}
