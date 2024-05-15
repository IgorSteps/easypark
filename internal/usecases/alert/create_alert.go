package usecases

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type CreateAlert struct {
	logger    *logrus.Logger
	alertRepo repositories.AlertRepository
}

func NewCreateAlert(l *logrus.Logger, r repositories.AlertRepository) *CreateAlert {
	return &CreateAlert{
		logger:    l,
		alertRepo: r,
	}
}

func (s *CreateAlert) Execute(ctx context.Context, alertType entities.AlertType, msg string, driverID, spaceID uuid.UUID) (*entities.Alert, error) {
	var alert *entities.Alert

	// Create correct alert depending on the type.
	switch alertType {
	case entities.LocationMismatch:
		alert = createLocationMismatchAlert(msg, driverID, spaceID)
	case entities.LateArrival:
		alert = createLateArrivalAlert(msg, driverID, spaceID)
	case entities.OverStay:
		alert = createOverStayAlert(msg, driverID, spaceID)
	default:
		s.logger.Warn("unknown alert type")
		return nil, repositories.NewInvalidInputError("unknown alert type")
	}

	err := s.alertRepo.Create(ctx, alert)
	if err != nil {
		return nil, err
	}

	return alert, nil
}

func createLateArrivalAlert(msg string, driverID, spaceID uuid.UUID) *entities.Alert {
	lateArrivalAlert := &entities.Alert{}
	lateArrivalAlert.CreateLateArrivalAlert(msg, driverID, spaceID)

	return lateArrivalAlert
}

func createOverStayAlert(msg string, driverID, spaceID uuid.UUID) *entities.Alert {
	overStayAlert := &entities.Alert{}
	overStayAlert.CreateOverStayAlert(msg, driverID, spaceID)

	return overStayAlert
}
func createLocationMismatchAlert(msg string, driverID, spaceID uuid.UUID) *entities.Alert {
	locationMisMatchAlert := &entities.Alert{}
	locationMisMatchAlert.CreateLocationMismatchAlert(msg, driverID, spaceID)

	return locationMisMatchAlert
}
