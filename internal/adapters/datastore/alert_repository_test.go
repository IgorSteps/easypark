package datastore_test

import (
	"context"
	"testing"

	"github.com/IgorSteps/easypark/internal/adapters/datastore"
	"github.com/IgorSteps/easypark/internal/domain/entities"
	mocks "github.com/IgorSteps/easypark/mocks/adapters/datastore"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAlertRepository_Create(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockDB := &mocks.Datastore{}
	repo := datastore.NewAlertPostgresRepository(testLogger, mockDB)

	testCtx := context.Background()
	testAlert := &entities.Alert{
		ID:             uuid.New(),
		Type:           entities.LocationMismatch,
		UserID:         uuid.New(),
		ParkingSpaceID: uuid.New(),
	}

	mockDB.EXPECT().WithContext(testCtx).Return(mockDB).Once()
	mockDB.EXPECT().Create(testAlert).Return(mockDB).Once()
	mockDB.EXPECT().Error().Return(nil).Once()

	// ---
	// ACT
	// ---
	err := repo.Create(testCtx, testAlert)

	// ------
	// ASSERT
	// ------
	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestAlertRepository_GetSingle(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockDB := &mocks.Datastore{}
	repo := datastore.NewAlertPostgresRepository(testLogger, mockDB)

	testAlertID := uuid.New()
	testCtx := context.Background()
	testAlert := entities.Alert{
		ID:             testAlertID,
		Type:           entities.LocationMismatch,
		UserID:         uuid.New(),
		ParkingSpaceID: uuid.New(),
	}

	mockDB.EXPECT().WithContext(testCtx).Return(mockDB).Once()
	var alert entities.Alert
	mockDB.EXPECT().First(&alert, "id = ?", testAlertID).Return(mockDB).Once().Run(func(args mock.Arguments) {
		arg := args.Get(0).(*entities.Alert) // Get the alert argument passed to Create
		*arg = testAlert                     // Set it to the expected alert
	})
	mockDB.EXPECT().Error().Return(nil).Once()

	// ---
	// ACT
	// ---
	actualAlert, err := repo.GetSingle(testCtx, testAlertID)

	// ------
	// ASSERT
	// ------
	assert.NoError(t, err)
	assert.Equal(t, testAlert, actualAlert)
	mockDB.AssertExpectations(t)
}
