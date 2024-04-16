package entities_test

import (
	"testing"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAlert_OnCreate(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	alert := &entities.Alert{}
	msg := "boom"
	driverID := uuid.New()
	spaceID := uuid.New()

	// ---
	// ACT
	// ---
	alert.CreateLocationMismatchAlert(msg, driverID, spaceID)

	// ------
	// ASSERT
	// ------
	assert.NotEmpty(t, alert.ID)
	assert.Equal(t, entities.LocationMismatch, alert.Type)
	assert.Equal(t, msg, alert.Message)
	assert.Equal(t, driverID, alert.UserID)
	assert.Equal(t, spaceID, alert.ParkingSpaceID)
}
