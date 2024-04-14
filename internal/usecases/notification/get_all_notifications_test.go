package usecases_test

import (
	"context"
	"testing"
	"time"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	usecases "github.com/IgorSteps/easypark/internal/usecases/notification"
	mocks "github.com/IgorSteps/easypark/mocks/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestGetAllNotifications_Execute(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockRepo := &mocks.NotificationRepository{}
	usecase := usecases.NewGetAllNotifications(testLogger, mockRepo)

	testCtx := context.Background()

	testNotifications := []entities.Notification{
		{
			ID:             uuid.New(),
			Type:           0,
			DriverID:       uuid.New(),
			ParkingSpaceID: uuid.New(),
			Location:       "bom",
			Timestamp:      time.Now(),
		},
		{
			ID:             uuid.New(),
			Type:           0,
			DriverID:       uuid.New(),
			ParkingSpaceID: uuid.New(),
			Location:       "bom",
			Timestamp:      time.Now(),
		},
		{
			ID:             uuid.New(),
			Type:           0,
			DriverID:       uuid.New(),
			ParkingSpaceID: uuid.New(),
			Location:       "bom",
			Timestamp:      time.Now(),
		},
	}
	mockRepo.EXPECT().GetAll(testCtx).Return(testNotifications, nil).Once()

	// ----
	// ACT
	// ----
	notifications, err := usecase.Execute(testCtx)

	// ------
	// ASSERT
	// ------
	assert.Nil(t, err, "Error must be nil")
	assert.Len(t, notifications, 3)
	assert.Equal(t, testNotifications, notifications)
}
