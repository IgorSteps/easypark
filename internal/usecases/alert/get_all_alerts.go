package usecases

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/sirupsen/logrus"
)

// GetAllAlert provides business logic to get a all alerts.
type GetAllAlerts struct {
	logger    *logrus.Logger
	alertRepo repositories.AlertRepository
}

// NewGetAllAlert returns a new instance of GetSingleAlert
func NewGetAllAlerts(l *logrus.Logger, repo repositories.AlertRepository) *GetAllAlerts {
	return &GetAllAlerts{
		logger:    l,
		alertRepo: repo,
	}
}

// Execute runs the business logic.
func (s *GetAllAlerts) Execute(ctx context.Context) ([]entities.Alert, error) {
	return s.alertRepo.GetAll(ctx)
}
