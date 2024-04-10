package datastore_test

import (
	"context"
	"errors"
	"testing"

	"github.com/IgorSteps/easypark/internal/adapters/datastore"
	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	mocks "github.com/IgorSteps/easypark/mocks/adapters/datastore"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestNotificationPostgresRepository_Create_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockDatastore := &mocks.Datastore{}
	repo := datastore.NewNotificationPostgresRepository(testLogger, mockDatastore)
	testCtx := context.Background()
	noification := entities.Notification{}

	mockDatastore.EXPECT().Create(&noification).Return(mockDatastore).Once()
	mockDatastore.EXPECT().WithContext(testCtx).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Error().Return(nil).Once()

	// ----
	// ACT
	// ----
	err := repo.Create(testCtx, &noification)

	// ------
	// ASSERT
	// ------
	assert.Nil(t, err, "Error must be nil")
	mockDatastore.AssertExpectations(t)
}

func TestNotificationPostgresRepository_Create_UnhappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockDatastore := &mocks.Datastore{}
	repo := datastore.NewNotificationPostgresRepository(testLogger, mockDatastore)
	testCtx := context.Background()
	noification := entities.Notification{}

	testError := errors.New("boom")
	mockDatastore.EXPECT().Create(&noification).Return(mockDatastore).Once()
	mockDatastore.EXPECT().WithContext(testCtx).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Error().Return(testError).Once()

	// ----
	// ACT
	// ----
	err := repo.Create(testCtx, &noification)

	// ------
	// ASSERT
	// ------
	assert.EqualError(t, err, "Internal error: failed to create a notification in the database", "Error is wrong")
	assert.IsType(t, &repositories.InternalError{}, err, "Error is of wrong type")
	mockDatastore.AssertExpectations(t)
}
