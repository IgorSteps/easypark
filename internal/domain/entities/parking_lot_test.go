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
	testLocation := "comp"
	testCapacity := 10

	// --------
	// ACT
	// --------
	lot.OnCreate(testName, testLocation, testCapacity)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, testName, lot.Name, "Parking lot names must match")
	assert.Equal(t, testLocation, lot.Location, "Parking lot locations must match")
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
