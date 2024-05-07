package usecasefacades_test

import (
	"context"
	"testing"

	"github.com/IgorSteps/easypark/internal/adapters/usecasefacades"
	"github.com/IgorSteps/easypark/internal/domain/entities"
	mocks "github.com/IgorSteps/easypark/mocks/adapters/usecasefacades"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNotificationFacade_Create(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	notificationCreator := &mocks.NotificationCreator{}
	notificationGetter := &mocks.NotificationGetter{}
	facade := usecasefacades.NewNotificationFacade(notificationCreator, notificationGetter)
	testDriverID := uuid.New()
	testParkReqID := uuid.New()
	testSpaceID := uuid.New()
	location := "cmp"
	notType := 1
	testCtx := context.Background()
	notif := entities.Notification{
		ID: uuid.New(),
	}
	notificationCreator.EXPECT().Execute(testCtx, testDriverID, testParkReqID, testSpaceID, location, notType).Return(notif, nil).Once()

	// --------
	// ACT
	// --------
	notification, err := facade.CreateNotification(testCtx, testDriverID, testParkReqID, testSpaceID, location, notType)

	// --------
	// ASSERT
	// --------
	assert.NoError(t, err)
	assert.Equal(t, notif, notification)
}

func TestNotificationFacade_GetAll(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	notificationCreator := &mocks.NotificationCreator{}
	notificationGetter := &mocks.NotificationGetter{}
	facade := usecasefacades.NewNotificationFacade(notificationCreator, notificationGetter)
	testCtx := context.Background()
	notifs := []entities.Notification{
		{ID: uuid.New()},
	}
	notificationGetter.EXPECT().Execute(testCtx).Return(notifs, nil).Once()

	// --------
	// ACT
	// --------
	notifications, err := facade.GetAllNotifications(testCtx)

	// --------
	// ASSERT
	// --------
	assert.NoError(t, err)
	assert.Equal(t, len(notifications), 1)
}
