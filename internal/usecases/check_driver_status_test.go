package usecases_test

import (
	"context"
	"errors"
	"testing"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/usecases"
	mocks "github.com/IgorSteps/easypark/mocks/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCheckDriverStatus_HappyPath_Active(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger := *logrus.New()
	mockRepo := &mocks.UserRepository{}
	usecase := usecases.NewCheckDriverStatus(&testLogger, mockRepo)
	ctx := context.Background()
	testID := uuid.New()

	var testUser entities.User
	mockRepo.EXPECT().GetDriverByID(ctx, testID, &testUser).Return(nil).Once().Run(func(args mock.Arguments) {
		arg := args.Get(2).(*entities.User)                 // Get the argument passed to FindAll
		*arg = entities.User{Status: entities.StatusActive} // Set it to the expected users
	})

	// --------
	// ACT
	// --------
	isBanned, err := usecase.Execute(ctx, testID)

	// --------
	// ASSERT
	// --------
	assert.NoError(t, err, "Must not have an error")
	assert.False(t, isBanned, "Must not be banned")
	mockRepo.AssertExpectations(t)
}

func TestCheckDriverStatus_HappyPath_Banned(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger := *logrus.New()
	mockRepo := &mocks.UserRepository{}
	usecase := usecases.NewCheckDriverStatus(&testLogger, mockRepo)
	ctx := context.Background()
	testID := uuid.New()

	var testUser entities.User
	mockRepo.EXPECT().GetDriverByID(ctx, testID, &testUser).Return(nil).Once().Run(func(args mock.Arguments) {
		arg := args.Get(2).(*entities.User)                 // Get the argument passed to FindAll
		*arg = entities.User{Status: entities.StatusBanned} // Set it to the expected users
	})

	// --------
	// ACT
	// --------
	isBanned, err := usecase.Execute(ctx, testID)

	// --------
	// ASSERT
	// --------
	assert.NoError(t, err, "Must not have an error")
	assert.True(t, isBanned, "Must be banned")
	mockRepo.AssertExpectations(t)
}

func TestCheckDriverStatus_UnhappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger := *logrus.New()
	mockRepo := &mocks.UserRepository{}
	usecase := usecases.NewCheckDriverStatus(&testLogger, mockRepo)
	ctx := context.Background()
	testID := uuid.New()
	testErr := errors.New("boom")
	var testUser entities.User
	mockRepo.EXPECT().GetDriverByID(ctx, testID, &testUser).Return(testErr).Once()

	// --------
	// ACT
	// --------
	isBanned, err := usecase.Execute(ctx, testID)

	// --------
	// ASSERT
	// --------
	assert.Error(t, err, "Must have an error")
	assert.EqualError(t, err, testErr.Error(), "Errors are not equal")
	assert.True(t, isBanned, "Must be true")
	mockRepo.AssertExpectations(t)
}
