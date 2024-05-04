package usecases_test

import (
	"context"
	"testing"
	"time"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	usecases "github.com/IgorSteps/easypark/internal/usecases/alert"
	mocks "github.com/IgorSteps/easypark/mocks/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestCheckOverStays(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	testLogger.Level = logrus.DebugLevel
	mockAlertCreator := &mocks.AlertCreator{}
	mockReqRepo := &mocks.ParkingRequestRepository{}
	usecase := usecases.NewCheckOverStays(testLogger, mockReqRepo, mockAlertCreator)

	testCtx := context.Background()

	parkSpaceID := uuid.New()
	testReqs := []entities.ParkingRequest{
		{
			ID:                      uuid.New(),
			UserID:                  uuid.New(),
			DestinationParkingLotID: uuid.New(),
			ParkingSpaceID:          &parkSpaceID,
			StartTime:               time.Now().Add(-3 * time.Hour),
			EndTime:                 time.Now().Add(-2 * time.Hour), // 2 hours after reservation end
			Status:                  entities.RequestStatusActive,
		},
		{
			ID:                      uuid.New(),
			UserID:                  uuid.New(),
			DestinationParkingLotID: uuid.New(),
			ParkingSpaceID:          &parkSpaceID,
			StartTime:               time.Now().Add(-3 * time.Hour),
			EndTime:                 time.Now().Add(-2 * time.Hour),
			Status:                  entities.RequestStatusActive,
		},
		// this one should get filtered out
		{
			ID:                      uuid.New(),
			UserID:                  uuid.New(),
			DestinationParkingLotID: uuid.New(),
			ParkingSpaceID:          &parkSpaceID,
			StartTime:               time.Now().Add(-1 * time.Hour),
			EndTime:                 time.Now().Add(-25 * time.Minute), // 25 minutes after reservation end
			Status:                  entities.RequestStatusActive,
		},
	}
	query := map[string]interface{}{
		"status": entities.RequestStatusActive,
	}
	mockReqRepo.EXPECT().GetMany(testCtx, query).Return(testReqs, nil).Once()

	// Filter.
	filteredReqs := []entities.ParkingRequest{
		testReqs[0],
		testReqs[1],
	}

	for _, req := range filteredReqs {
		testAlert := &entities.Alert{
			ID:             uuid.New(),
			Type:           entities.OverStay,
			Message:        "boom",
			UserID:         req.UserID,
			ParkingSpaceID: *req.ParkingSpaceID,
		}
		mockAlertCreator.EXPECT().Execute(
			testCtx,
			entities.OverStay,
			"exit notification hasn't been received after 30 minutes from the parking request end time",
			req.UserID,
			*req.ParkingSpaceID,
		).Return(testAlert, nil).Once()
	}

	// --------
	// ACT
	// --------
	alerts, err := usecase.Execute(testCtx, time.Hour)

	// --------
	// ASSERT
	// --------
	assert.NoError(t, err)
	assert.Len(t, alerts, len(filteredReqs))
	assert.Equal(t, filteredReqs, hook.LastEntry().Data["requests"])
	for i, alert := range alerts {
		assert.Equal(t, filteredReqs[i].UserID, alert.UserID, "User IDs don't match")
		assert.Equal(t, filteredReqs[i].ParkingSpaceID, &alert.ParkingSpaceID, "Parking space IDs don't match")
	}
	mockReqRepo.AssertExpectations(t)
	mockAlertCreator.AssertExpectations(t)
}

func TestCheckSignificantlyOverdueStays(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger := logrus.New()             // Create a new logger
	testLogger.SetLevel(logrus.DebugLevel) // Set the desired log level
	mockAlertCreator := &mocks.AlertCreator{}
	mockReqRepo := &mocks.ParkingRequestRepository{}
	usecase := usecases.NewCheckOverStays(testLogger, mockReqRepo, mockAlertCreator)

	testCtx := context.Background()

	parkSpaceID := uuid.New()
	testReqs := []entities.ParkingRequest{
		{
			ID:                      uuid.New(),
			UserID:                  uuid.New(),
			DestinationParkingLotID: uuid.New(),
			ParkingSpaceID:          &parkSpaceID,
			StartTime:               time.Now().Add(-5 * time.Hour),
			EndTime:                 time.Now().Add(-4 * time.Hour), // 4 hours past
			Status:                  entities.RequestStatusActive,
		},
	}
	query := map[string]interface{}{
		"status": entities.RequestStatusActive,
	}
	mockReqRepo.EXPECT().GetMany(testCtx, query).Return(testReqs, nil).Once()

	expectedAlert := &entities.Alert{
		ID:             uuid.New(),
		Type:           entities.OverStay,
		Message:        "Significantly overdue",
		UserID:         testReqs[0].UserID,
		ParkingSpaceID: *testReqs[0].ParkingSpaceID,
	}
	mockAlertCreator.EXPECT().Execute(
		testCtx,
		entities.OverStay,
		"exit notification hasn't been received after 30 minutes from the parking request end time",
		testReqs[0].UserID,
		*testReqs[0].ParkingSpaceID,
	).Return(expectedAlert, nil).Once()

	// --------
	// ACT
	// --------
	alerts, err := usecase.Execute(testCtx, time.Hour)

	// --------
	// ASSERT
	// --------
	assert.NoError(t, err)
	assert.Len(t, alerts, 1)
	assert.Equal(t, expectedAlert, alerts[0], "Alerts should match")
	assert.Equal(t, expectedAlert.UserID, alerts[0].UserID, "User IDs don't match")
	assert.Equal(t, expectedAlert.ParkingSpaceID, alerts[0].ParkingSpaceID, "Parking space IDs don't match")
	mockReqRepo.AssertExpectations(t)
	mockAlertCreator.AssertExpectations(t)
}
