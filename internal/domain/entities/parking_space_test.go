package entities_test

import (
	"testing"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestParkingSpace_OnCreate(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	space := entities.ParkingSpace{}
	testName := "cool name"
	testParkingLotID := uuid.New()

	// --------
	// ACT
	// --------
	space.OnCreate(testName, testParkingLotID)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, testName, space.Name, "Parking space names must match")
	assert.Equal(t, testParkingLotID, space.ParkingLotID, "Parking space's parking lot ids must match")
	assert.Equal(t, entities.StatusAvailable, space.Status, "Parking space statuses mast match")
	assert.NotNil(t, space.ID, "Parking space must have ID set")
}
