package datastore

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// AlertPostgresRepository implements AlertRepository interface to provide database operations on Alerts.
type AlertPostgresRepository struct {
	logger *logrus.Logger
	db     Datastore
}

// NewAlertPostgresRepository returns a new instance of AlertPostgresRepository.
func NewAlertPostgresRepository(l *logrus.Logger, db Datastore) *AlertPostgresRepository {
	return &AlertPostgresRepository{
		logger: l,
		db:     db,
	}
}

// Create insert a new alert into the database.
func (s *AlertPostgresRepository) Create(ctx context.Context, alert *entities.Alert) error {
	result := s.db.WithContext(ctx).Create(alert)

	err := result.Error()
	if err != nil {
		s.logger.WithError(err).Error("failed to insert alert into the database")
		return repositories.NewInternalError("failed to insert alert into the database")
	}

	return nil
}

// GetSingle returns a single alert from the database using its id.
func (s *AlertPostgresRepository) GetSingle(ctx context.Context, id uuid.UUID) (entities.Alert, error) {
	var alert entities.Alert

	result := s.db.WithContext(ctx).First(&alert, "id = ?", id)
	err := result.Error()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.WithField("id", id).Error("failed to find the alert with given id in the database")
			return entities.Alert{}, repositories.NewNotFoundError(id.String())
		}

		s.logger.WithError(err).Error("failed to query for alert in the database")
		return entities.Alert{}, repositories.NewInternalError("failed to query for alert in the database")
	}

	return alert, nil
}

func (s *AlertPostgresRepository) GetAll(ctx context.Context) ([]entities.Alert, error) {
	var alerts []entities.Alert
	result := s.db.WithContext(ctx).FindAll(&alerts)

	err := result.Error()
	if err != nil {
		s.logger.WithError(err).Error("failed to get all alerts in the database")
		return nil, repositories.NewInternalError("failed to get all alerts in the database")
	}

	return alerts, nil
}
