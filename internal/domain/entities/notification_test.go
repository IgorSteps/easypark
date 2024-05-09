package entities_test

import (
	"testing"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNotifiation_OnCreate(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	notification := &entities.Notification{}
	testDriverID := uuid.New()
	testParkSpaceID := uuid.New()
	testLocation := "blob"
	testType := entities.ArrivalNotification

	// ----
	// ACT
	// ----
	notification.OnCreate(testDriverID, testParkSpaceID, testLocation, testType)

	// ------
	// ASSERT
	// ------
	assert.NotNil(t, notification.ID)
	assert.Equal(t, testDriverID, notification.DriverID)
	assert.Equal(t, testParkSpaceID, notification.ParkingSpaceID)
	assert.Equal(t, testLocation, notification.Location)
	assert.Equal(t, testType, notification.Type)
	assert.NotNil(t, notification.CreatedAt)
}
