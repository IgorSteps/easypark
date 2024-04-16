package usecases_test

import (
	"context"
	"testing"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	usecases "github.com/IgorSteps/easypark/internal/usecases/alert"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mocks "github.com/IgorSteps/easypark/mocks/domain/repositories"
	"github.com/sirupsen/logrus/hooks/test"
)

func TestCreateAlert_Execute_LocationMismatchAlert(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockRepo := &mocks.AlertRepository{}
	usecase := usecases.NewCreateAlert(testLogger, mockRepo)

	testCtx := context.Background()
	alertType := entities.LocationMismatch
	msg := "boom"
	driverID := uuid.New()
	spaceID := uuid.New()
	locationMisMatchAlert := &entities.Alert{}
	locationMisMatchAlert.CreateLocationMismatchAlert(msg, driverID, spaceID)

	mockRepo.EXPECT().Create(testCtx, mock.Anything).Return(nil).Once()

	// ---
	// ACT
	// ---
	alert, err := usecase.Execute(testCtx, alertType, msg, driverID, spaceID)

	// ------
	// ASSERT
	// ------
	assert.NoError(t, err, "Must not error")
	assert.Equal(t, locationMisMatchAlert.Message, alert.Message)
	assert.Equal(t, locationMisMatchAlert.Type, alert.Type)
	assert.Equal(t, locationMisMatchAlert.UserID, alert.UserID)
	assert.Equal(t, locationMisMatchAlert.ParkingSpaceID, alert.ParkingSpaceID)
}
