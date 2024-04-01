package entities_test

import (
	"testing"

	"github.com/IgorSteps/easypark/internal/domain/entities"
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
