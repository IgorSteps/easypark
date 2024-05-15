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

func TestCreateAlert_Execute_OverStayAlert(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockRepo := &mocks.AlertRepository{}
	usecase := usecases.NewCreateAlert(testLogger, mockRepo)

	testCtx := context.Background()
	alertType := entities.OverStay
	msg := "overstay detected"
	driverID := uuid.New()
	spaceID := uuid.New()
	overStayAlert := &entities.Alert{}
	overStayAlert.CreateOverStayAlert(msg, driverID, spaceID)

	mockRepo.EXPECT().Create(testCtx, mock.Anything).Return(nil).Once()

	// ---
	// ACT
	// ---
	alert, err := usecase.Execute(testCtx, alertType, msg, driverID, spaceID)

	// ------
	// ASSERT
	// ------
	assert.NoError(t, err, "No error")
	assert.Equal(t, overStayAlert.Message, alert.Message)
	assert.Equal(t, overStayAlert.Type, alert.Type)
	assert.Equal(t, overStayAlert.UserID, alert.UserID)
	assert.Equal(t, overStayAlert.ParkingSpaceID, alert.ParkingSpaceID)
}

func TestCreateAlert_Execute_LateArrivalAlert(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockRepo := &mocks.AlertRepository{}
	usecase := usecases.NewCreateAlert(testLogger, mockRepo)

	testCtx := context.Background()
	alertType := entities.LateArrival
	msg := "late arrival noted"
	driverID := uuid.New()
	spaceID := uuid.New()
	lateArrivalAlert := &entities.Alert{}
	lateArrivalAlert.CreateLateArrivalAlert(msg, driverID, spaceID)

	mockRepo.EXPECT().Create(testCtx, mock.Anything).Return(nil).Once()

	// ---
	// ACT
	// ---
	alert, err := usecase.Execute(testCtx, alertType, msg, driverID, spaceID)

	// ------
	// ASSERT
	// ------
	assert.NoError(t, err, "No error")
	assert.Equal(t, lateArrivalAlert.Message, alert.Message)
	assert.Equal(t, lateArrivalAlert.Type, alert.Type)
	assert.Equal(t, lateArrivalAlert.UserID, alert.UserID)
	assert.Equal(t, lateArrivalAlert.ParkingSpaceID, alert.ParkingSpaceID)
}
