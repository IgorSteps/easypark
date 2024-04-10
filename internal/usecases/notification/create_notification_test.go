package usecases_test

import (
	"context"
	"testing"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	usecases "github.com/IgorSteps/easypark/internal/usecases/notification"
	mocks "github.com/IgorSteps/easypark/mocks/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateNotification_Execute_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockRepo := &mocks.NotificationRepository{}
	usecase := usecases.NewCreateNotification(testLogger, mockRepo)

	testCtx := context.Background()
	testDriverID := uuid.New()
	testSpaceID := uuid.New()
	testLocation := "boom"
	testNotificationType := 0

	mockRepo.EXPECT().Create(testCtx, mock.Anything).Return(nil).Once()

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
	mockRepo.AssertExpectations(t)
}

func TestCreateNotification_Execute_UnhappyPath_ParsingFailed(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockRepo := &mocks.NotificationRepository{}
	usecase := usecases.NewCreateNotification(testLogger, mockRepo)

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
