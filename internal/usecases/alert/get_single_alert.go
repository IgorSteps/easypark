package usecases

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// GetSingleAlert provides business logic to get a single alert.
type GetSingleAlert struct {
	logger    *logrus.Logger
	alertRepo repositories.AlertRepository
}

// NewGetSingleAlert returns a new instance of GetSingleAlert
func NewGetSingleAlert(l *logrus.Logger, repo repositories.AlertRepository) *GetSingleAlert {
	return &GetSingleAlert{
		logger:    l,
		alertRepo: repo,
	}
}

// Execute runs the business logic.
func (s *GetSingleAlert) Execute(ctx context.Context, id uuid.UUID) (entities.Alert, error) {
	return s.alertRepo.GetSingle(ctx, id)
}
