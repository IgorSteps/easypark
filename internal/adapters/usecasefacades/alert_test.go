package usecasefacades_test

import (
	"context"
	"testing"
	"time"

	"github.com/IgorSteps/easypark/internal/adapters/usecasefacades"
	"github.com/IgorSteps/easypark/internal/domain/entities"
	mocks "github.com/IgorSteps/easypark/mocks/adapters/usecasefacades"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAlertFacade_GetSingle(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	singleGetter := &mocks.AlertSingleGetter{}
	allGetter := &mocks.AlertAllGetter{}
	lateArrivalChecker := &mocks.AlertLateArrivalChecker{}
	overStayChecker := &mocks.AlertOverStayChecker{}

	facade := usecasefacades.NewAlertFacade(singleGetter, lateArrivalChecker, overStayChecker, allGetter)

	testID := uuid.New()
	testCtx := context.Background()
	alert := entities.Alert{
		ID: uuid.New(),
	}
	singleGetter.EXPECT().Execute(testCtx, testID).Return(alert, nil).Once()

	// --------
	// ACT
	// --------
	actualAlert, err := facade.GetAlert(testCtx, testID)

	// --------
	// ASSERT
	// --------
	assert.NoError(t, err)
	assert.Equal(t, alert, actualAlert)
}

func TestAlertFacade_GetAll(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	singleGetter := &mocks.AlertSingleGetter{}
	allGetter := &mocks.AlertAllGetter{}
	lateArrivalChecker := &mocks.AlertLateArrivalChecker{}
	overStayChecker := &mocks.AlertOverStayChecker{}

	facade := usecasefacades.NewAlertFacade(singleGetter, lateArrivalChecker, overStayChecker, allGetter)
	testCtx := context.Background()
	alerts := []entities.Alert{
		{ID: uuid.New()},
	}
	allGetter.EXPECT().Execute(testCtx).Return(alerts, nil).Once()

	// --------
	// ACT
	// --------
	actualAlerts, err := facade.GetAllAlerts(testCtx)

	// --------
	// ASSERT
	// --------
	assert.NoError(t, err)
	assert.Equal(t, len(actualAlerts), 1)
}

func TestAlertFacade_LateArrivalCheck(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	singleGetter := &mocks.AlertSingleGetter{}
	allGetter := &mocks.AlertAllGetter{}
	lateArrivalChecker := &mocks.AlertLateArrivalChecker{}
	overStayChecker := &mocks.AlertOverStayChecker{}

	facade := usecasefacades.NewAlertFacade(singleGetter, lateArrivalChecker, overStayChecker, allGetter)
	testCtx := context.Background()
	alerts := []entities.Alert{
		{ID: uuid.New()},
	}
	threshold := time.Duration(10)
	lateArrivalChecker.EXPECT().Execute(testCtx, threshold).Return(alerts, nil).Once()

	// --------
	// ACT
	// --------
	actualAlerts, err := facade.CheckForLateArrivals(testCtx, threshold)

	// --------
	// ASSERT
	// --------
	assert.NoError(t, err)
	assert.Equal(t, len(actualAlerts), 1)
}
