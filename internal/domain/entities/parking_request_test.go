package entities_test

import (
	"testing"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestParkingRequest_OnCreate(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	parkRequest := entities.ParkingRequest{}

	// --------
	// ACT
	// --------
	parkRequest.OnCreate()

	// --------
	// ASSERT
	// --------
	assert.NotNil(t, parkRequest.ID, "Parking Request ID must be set.")
	assert.Equal(t, entities.RequestStatusPending, parkRequest.Status, "Status must be pending")
}

func TestParkingRequest_OnArrivalNotification(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	parkRequest := entities.ParkingRequest{}

	// --------
	// ACT
	// --------
	parkRequest.OnArrivalNotification()

	// --------
	// ASSERT
	// --------
	assert.Equal(t, entities.RequestStatusActive, parkRequest.Status, "Status must be active")
}

func TestParkingRequest_OnDeparture(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	parkRequest := entities.ParkingRequest{}

	// --------
	// ACT
	// --------
	parkRequest.OnDepartureNotification()

	// --------
	// ASSERT
	// --------
	assert.Equal(t, entities.RequestStatusCompleted, parkRequest.Status, "Status must be completed")
}

func TestParkingRequest_OnAssign(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	parkRequest := entities.ParkingRequest{}
	testSpaceID := uuid.New()
	// --------
	// ACT
	// --------
	parkRequest.OnSpaceAssign(testSpaceID)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, testSpaceID, *parkRequest.ParkingSpaceID)
	assert.Equal(t, entities.RequestStatusApproved, parkRequest.Status, "Status must be approved")
}
