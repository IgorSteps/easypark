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

func TestGetAllAlerts_Execute(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------

	testLogger, _ := test.NewNullLogger()
	mockrepo := &mocks.AlertRepository{}
	usecase := usecases.NewGetAllAlerts(testLogger, mockrepo)

	testCtx := context.Background()
	testAlerts := []entities.Alert{
		{
			ID: uuid.New(),
		},
	}
	mockrepo.EXPECT().GetAll(testCtx).Return(testAlerts, nil)

	// --------
	// ACT
	// --------

	alerts, err := usecase.Execute(testCtx)

	// ------
	// ASSERT
	// ------
	assert.NoError(t, err)
	assert.Equal(t, testAlerts, alerts)
	mockrepo.AssertExpectations(t)
}
