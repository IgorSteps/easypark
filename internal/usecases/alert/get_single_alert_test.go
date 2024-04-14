package usecases_test

import (
	"context"
	"testing"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	usecases "github.com/IgorSteps/easypark/internal/usecases/alert"
	mocks "github.com/IgorSteps/easypark/mocks/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestGetSingleAlert_Execute(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockrepo := &mocks.AlertRepository{}
	usecase := usecases.NewGetSingleAlert(testLogger, mockrepo)

	testCtx := context.Background()

	testAlert := entities.Alert{
		ID:             uuid.New(),
		Type:           entities.LocationMismatch,
		UserID:         uuid.New(),
		ParkingSpaceID: uuid.New(),
	}
	mockrepo.EXPECT().GetSingle(testCtx, testAlert.ID).Return(testAlert, nil)

	// ---
	// ACT
	// ---
	alert, err := usecase.Execute(testCtx, testAlert.ID)

	// ------
	// ASSERT
	// ------
	assert.NoError(t, err)
	assert.Equal(t, testAlert, alert)
}
