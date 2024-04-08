package entities_test

import (
	"fmt"
	"testing"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestParkingLot_OnCreate(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	lot := entities.ParkingLot{}
	testName := "cool name"
	testCapacity := 10

	// --------
	// ACT
	// --------
	lot.OnCreate(testName, testCapacity)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, testName, lot.Name, "Parking lot names must match")
	assert.Equal(t, testCapacity, lot.Capacity, "Parking lot capacities must match")

	// Assert parking spaces were created correctly:
	for i, space := range lot.ParkingSpaces {
		assert.NotNil(t, space.ID, "Parking space ID must be set")
		assert.Equal(t, lot.ID, space.ParkingLotID, "Parking space must have a reference to the parking lot")
		assert.Equal(t, entities.StatusAvailable, space.Status, "Parking space status must be available")

		// Assert parking space name:
		assert.Contains(t, space.Name, testName, "Parking space number should include the parking lot name")
		assert.Contains(t, space.Name, "-", "Parking space number must include a dash")
		assert.Contains(t, space.Name, fmt.Sprint(i+1), "Parking space number must include a number")
	}
}

func TestParkingLot_OnGet(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	lot := entities.ParkingLot{}
	testName := "cool name"
	testCapacity := 10

	// Create a lot with 10 spaces.
	lot.OnCreate(testName, testCapacity)
	// Set parking space statuses
	lot.ParkingSpaces[0].Status = entities.StatusBlocked
	lot.ParkingSpaces[1].Status = entities.StatusOccupied
	lot.ParkingSpaces[2].Status = entities.StatusReserved

	// --------
	// ACT
	// --------
	lot.OnGet()

	// --------
	// ASSERT
	// --------
	assert.Equal(t, 7, lot.Available, "Must have 7 available spaces")
	assert.Equal(t, 1, lot.Blocked, "Must have 1 blocked space")
	assert.Equal(t, 1, lot.Reserved, "Must have 1 reserved space")
	assert.Equal(t, 1, lot.Occupied, "Must have 1 occupied space")
}
