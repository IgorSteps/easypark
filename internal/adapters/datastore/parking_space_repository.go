package datastore

import (
	"context"
	"time"

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

// GetSingle gets a parking space with a given ID.
func (s *ParkingSpacePostgresRepository) GetSingle(ctx context.Context, id uuid.UUID) (entities.ParkingSpace, error) {
	var space entities.ParkingSpace

	result := s.DB.WithContext(ctx).Preload("ParkingRequests").First(&space, "id = ?", id)
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

// GetMany gets many parking spaces that match given query.
func (s *ParkingSpacePostgresRepository) GetMany(ctx context.Context, query map[string]interface{}) ([]entities.ParkingSpace, error) {
	var spaces []entities.ParkingSpace

	result := s.DB.WithContext(ctx).Where(query).Preload("ParkingRequests").FindAll(&spaces)

	err := result.Error()
	if err != nil {
		s.Logger.WithError(err).WithField("query", query).Error("failed to query for parking spaces")
		return nil, repositories.NewInternalError("failed to query for parking spaces")
	}

	return spaces, nil
}

// FindAvailableSpaces finds all available parking spaces.
func (s *ParkingSpacePostgresRepository) FindAvailableSpaces(
	ctx context.Context,
	lotID uuid.UUID,
	startTime,
	endTime time.Time,
) ([]entities.ParkingSpace, error) {
	var spaces []entities.ParkingSpace
	s.Logger.WithFields(logrus.Fields{
		"lotID":      lotID,
		"start time": startTime,
		"end time":   endTime,
	}).Debug("parameters")
	// Find all spaces and their parking requests that are available in the specified parking lot.
	result := s.DB.WithContext(ctx).
		Where("parking_lot_id = ? AND status = ?", lotID, entities.ParkingSpaceStatusAvailable).
		Preload("ParkingRequests").
		FindAll(&spaces)

	if result.Error() != nil {
		s.Logger.WithError(result.Error()).Error("failed to query for parking spaces")
		return nil, repositories.NewInternalError("failed to query for parking spaces")
	}
	s.Logger.WithField("spaces", spaces).Debug("got all available parking spaces in the desired parking lot")
	// Filter out spaces with overlapping parking requests.
	// TODO: can it be converted to a DB query?
	var availableSpaces []entities.ParkingSpace
	for _, space := range spaces {
		if !space.CheckForOverlap(startTime, endTime) {
			availableSpaces = append(availableSpaces, space)
		}
	}
	s.Logger.WithField("spaces", spaces).Debug("filtered parking spaces t")

	spaces = availableSpaces
	return spaces, nil
}
