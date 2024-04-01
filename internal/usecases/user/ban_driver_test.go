package usecases_test

import (
	"context"
	"errors"
	"testing"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	usecases "github.com/IgorSteps/easypark/internal/usecases/user"
	mocks "github.com/IgorSteps/easypark/mocks/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestBanDriver_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger := *logrus.New()
	mockRepo := &mocks.UserRepository{}
	usecase := usecases.NewBanDriver(&testLogger, mockRepo)
	ctx := context.Background()
	testID := uuid.New()

	var testUser entities.User
	mockRepo.EXPECT().GetDriverByID(ctx, testID, &testUser).Return(nil).Once()

	var testUserBanned entities.User
	testUserBanned.Ban()
	mockRepo.EXPECT().Save(ctx, &testUserBanned).Return(nil).Once()

	// --------
	// ACT
	// --------
	err := usecase.Execute(ctx, testID)

	// --------
	// ASSERT
	// --------
	assert.NoError(t, err, "Must not have an error")
	mockRepo.AssertExpectations(t)
}

func TestBanDriver_UnhappyPath_GetDriverByID(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger := *logrus.New()
	mockRepo := &mocks.UserRepository{}
	usecase := usecases.NewBanDriver(&testLogger, mockRepo)
	ctx := context.Background()
	testID := uuid.New()
	testErr := errors.New("boom")

	var testUser entities.User
	mockRepo.EXPECT().GetDriverByID(ctx, testID, &testUser).Return(testErr).Once()

	// --------
	// ACT
	// --------
	err := usecase.Execute(ctx, testID)

	// --------
	// ASSERT
	// --------
	assert.Error(t, err, "Must have an error")
	assert.EqualError(t, err, testErr.Error(), "Error are not equal")
	mockRepo.AssertExpectations(t)
}

func TestBanDriver_UnhappyPath_Save(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger := *logrus.New()
	mockRepo := &mocks.UserRepository{}
	usecase := usecases.NewBanDriver(&testLogger, mockRepo)
	ctx := context.Background()
	testID := uuid.New()
	testErr := errors.New("boom")

	var testUser entities.User
	mockRepo.EXPECT().GetDriverByID(ctx, testID, &testUser).Return(nil).Once()

	var testUserBanned entities.User
	testUserBanned.Ban()
	mockRepo.EXPECT().Save(ctx, &testUserBanned).Return(testErr).Once()

	// --------
	// ACT
	// --------
	err := usecase.Execute(ctx, testID)

	// --------
	// ASSERT
	// --------
	assert.Error(t, err, "Must have an error")
	assert.EqualError(t, err, testErr.Error(), "Error are not equal")
	mockRepo.AssertExpectations(t)
}
