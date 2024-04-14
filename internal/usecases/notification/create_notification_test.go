package usecases_test

import (
	"context"
	"errors"
	"testing"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	usecases "github.com/IgorSteps/easypark/internal/usecases/notification"
	mocks "github.com/IgorSteps/easypark/mocks/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateNotification_Execute_HappyPath_Arrival(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockNotificationRepo := &mocks.NotificationRepository{}
	mockSpaceRepo := &mocks.ParkingSpaceRepository{}
	mockAlertCreator := &mocks.AlertCreator{}
	usecase := usecases.NewCreateNotification(testLogger, mockNotificationRepo, mockSpaceRepo, mockAlertCreator)

	testCtx := context.Background()
	testDriverID := uuid.New()
	testSpaceID := uuid.New()
	testLocation := "boom"
	testNotificationType := 0 // arrival

	testParkingSpace := entities.ParkingSpace{
		ID:     testSpaceID,
		Name:   testLocation,
		Status: entities.StatusAvailable,
	}
	mockSpaceRepo.EXPECT().GetParkingSpaceByID(testCtx, testSpaceID).Return(testParkingSpace, nil).Once()
	mockNotificationRepo.EXPECT().Create(testCtx, mock.Anything).Return(nil).Once()
	testParkingSpace.Status = entities.StatusOccupied
	mockSpaceRepo.EXPECT().Save(testCtx, &testParkingSpace).Return(nil).Once()

	// --------
	// ACT
	// --------
	notification, err := usecase.Execute(testCtx, testDriverID, testSpaceID, testLocation, testNotificationType)

	// --------
	// ASSERT
	// --------
	assert.Nil(t, err, "Error must be nil")
	assert.Equal(t, entities.ArrivalNotification, notification.Type, "Notification type is wrong")
	assert.Equal(t, testDriverID, notification.DriverID)
	assert.Equal(t, testLocation, notification.Location)
	assert.Equal(t, testSpaceID, notification.ParkingSpaceID)
	mockNotificationRepo.AssertExpectations(t)
	mockSpaceRepo.AssertExpectations(t)
}

func TestCreateNotification_Execute_HappyPath_Departure(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockNotificationRepo := &mocks.NotificationRepository{}
	mockSpaceRepo := &mocks.ParkingSpaceRepository{}
	mockAlertCreator := &mocks.AlertCreator{}
	usecase := usecases.NewCreateNotification(testLogger, mockNotificationRepo, mockSpaceRepo, mockAlertCreator)

	testCtx := context.Background()
	testDriverID := uuid.New()
	testSpaceID := uuid.New()
	testLocation := "boom"
	testNotificationType := 1 // departure

	testParkingSpace := entities.ParkingSpace{
		ID:     testSpaceID,
		Status: entities.StatusAvailable,
	}
	mockSpaceRepo.EXPECT().GetParkingSpaceByID(testCtx, testSpaceID).Return(testParkingSpace, nil).Once()
	mockNotificationRepo.EXPECT().Create(testCtx, mock.Anything).Return(nil).Once()
	testParkingSpace.Status = entities.StatusAvailable
	mockSpaceRepo.EXPECT().Save(testCtx, &testParkingSpace).Return(nil).Once()

	// --------
	// ACT
	// --------
	notification, err := usecase.Execute(testCtx, testDriverID, testSpaceID, testLocation, testNotificationType)

	// --------
	// ASSERT
	// --------
	assert.Nil(t, err, "Error must be nil")
	assert.Equal(t, entities.DepartureNotification, notification.Type, "Notification type is wrong")
	assert.Equal(t, testDriverID, notification.DriverID)
	assert.Equal(t, testLocation, notification.Location)
	assert.Equal(t, testSpaceID, notification.ParkingSpaceID)
	mockNotificationRepo.AssertExpectations(t)
	mockSpaceRepo.AssertExpectations(t)
}

func TestCreateNotification_Execute_UnhappyPath_ParsingFailed(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockRepo := &mocks.NotificationRepository{}
	mockSpaceRepo := &mocks.ParkingSpaceRepository{}
	mockAlertCreator := &mocks.AlertCreator{}
	usecase := usecases.NewCreateNotification(testLogger, mockRepo, mockSpaceRepo, mockAlertCreator)

	testCtx := context.Background()
	testDriverID := uuid.New()
	testSpaceID := uuid.New()
	testLocation := "boom"
	testNotificationType := 100

	// --------
	// ACT
	// --------
	notification, err := usecase.Execute(testCtx, testDriverID, testSpaceID, testLocation, testNotificationType)

	// --------
	// ASSERT
	// --------
	assert.EqualError(t, err, "invalid notification type", "Error message is wrong")
	assert.IsType(t, &repositories.InvalidInputError{}, err, "Error is of wrong type")
	assert.Equal(t, entities.ArrivalNotification, notification.Type, "Notification type is wrong")
	mockRepo.AssertExpectations(t)
}

func TestCreateNotification_Execute_LocationMismatchAlert(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	testLogger.Level = logrus.DebugLevel
	mockRepo := &mocks.NotificationRepository{}
	mockSpaceRepo := &mocks.ParkingSpaceRepository{}
	mockAlertCreator := &mocks.AlertCreator{}
	usecase := usecases.NewCreateNotification(testLogger, mockRepo, mockSpaceRepo, mockAlertCreator)

	testCtx := context.Background()
	testDriverID := uuid.New()
	testSpaceID := uuid.New()
	testLocation := "boom"
	testNotificationType := 0

	testParkingSpace := entities.ParkingSpace{
		ID:     testSpaceID,
		Name:   "different location", // different location
		Status: entities.StatusAvailable,
	}

	mockSpaceRepo.EXPECT().GetParkingSpaceByID(testCtx, testSpaceID).Return(testParkingSpace, nil).Once()
	mockRepo.EXPECT().Create(testCtx, mock.Anything).Return(nil).Once()
	testAlert := &entities.Alert{}
	testAlertMsg := "driver arrived at wrong parking space"
	testAlert.OnLocationMismatchAlertCreate(testAlertMsg, testDriverID, testSpaceID)
	mockAlertCreator.EXPECT().Execute(testCtx, entities.LocationMismatch, testAlertMsg, testDriverID, testSpaceID).Return(testAlert, nil).Once()
	testParkingSpace.Status = entities.StatusOccupied
	mockSpaceRepo.EXPECT().Save(testCtx, &testParkingSpace).Return(nil).Once()

	// --------
	// ACT
	// --------
	notification, err := usecase.Execute(testCtx, testDriverID, testSpaceID, testLocation, testNotificationType)

	// --------
	// ASSERT
	// --------
	assert.NoError(t, err, "Must not return an error")
	assert.Equal(t, entities.ArrivalNotification, notification.Type, "Notification type is wrong")
	assert.Equal(t, testDriverID, notification.DriverID)
	assert.Equal(t, testLocation, notification.Location)
	assert.Equal(t, testSpaceID, notification.ParkingSpaceID)

	// assert logger
	assert.Equal(t, 2, len(hook.Entries), "Logger must log twice")
	assert.Equal(t, "created alert for location mismatch on arrival", hook.Entries[0].Message)
	assert.Equal(t, testAlert, hook.Entries[0].Data["alert"])
	assert.Equal(t, "location mismatch, not updating parking space status to occupied", hook.Entries[1].Message)
}

func TestCreateNotification_Execute_LocationMismatchAlert_Fail(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	testLogger.Level = logrus.DebugLevel
	mockRepo := &mocks.NotificationRepository{}
	mockSpaceRepo := &mocks.ParkingSpaceRepository{}
	mockAlertCreator := &mocks.AlertCreator{}
	usecase := usecases.NewCreateNotification(testLogger, mockRepo, mockSpaceRepo, mockAlertCreator)

	testCtx := context.Background()
	testDriverID := uuid.New()
	testSpaceID := uuid.New()
	testLocation := "boom"
	testNotificationType := 0

	testParkingSpace := entities.ParkingSpace{
		ID:     testSpaceID,
		Name:   "different location", // different location
		Status: entities.StatusAvailable,
	}
	mockSpaceRepo.EXPECT().GetParkingSpaceByID(testCtx, testSpaceID).Return(testParkingSpace, nil).Once()
	mockRepo.EXPECT().Create(testCtx, mock.Anything).Return(nil).Once()

	testAlertMsg := "driver arrived at wrong parking space"
	testError := errors.New("boom")
	// return error
	mockAlertCreator.EXPECT().Execute(testCtx, entities.LocationMismatch, testAlertMsg, testDriverID, testSpaceID).Return(nil, testError).Once()

	testParkingSpace.Status = entities.StatusOccupied
	mockSpaceRepo.EXPECT().Save(testCtx, &testParkingSpace).Return(nil).Once()

	// --------
	// ACT
	// --------
	notification, err := usecase.Execute(testCtx, testDriverID, testSpaceID, testLocation, testNotificationType)

	// --------
	// ASSERT
	// --------
	assert.Empty(t, notification, "Must be empty")
	assert.EqualError(t, err, testError.Error(), "Error is wrong")

	// assert logger
	assert.Equal(t, 1, len(hook.Entries), "Logger must log once")
	assert.Equal(t, "failed to create location mismatch alert", hook.LastEntry().Message)
}
