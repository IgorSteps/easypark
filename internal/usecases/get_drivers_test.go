package usecases_test

import (
	"context"
	"errors"
	"testing"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/usecases"
	mocks "github.com/IgorSteps/easypark/mocks/domain/repositories"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestGetDriverUser_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger := *logrus.New()
	mockRepo := &mocks.UserRepository{}
	usecase := usecases.NewGetDrivers(&testLogger, mockRepo)
	ctx := context.Background()
	expectedUsers := []entities.User{
		{Username: "user1", Email: "user1@example.com"},
		{Username: "user2", Email: "user2@example.com"},
	}
	mockRepo.EXPECT().GetAllDriverUsers(ctx).Return(expectedUsers, nil).Once()

	// --------
	// ACT
	// --------
	users, err := usecase.Execute(ctx)

	// --------
	// ASSERT
	// --------
	assert.NoError(t, err, "Must not have an error")
	assert.Equal(t, expectedUsers, users, "User slices must match")
}

func TestGetDriverUser_UnhappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger := *logrus.New()
	mockRepo := &mocks.UserRepository{}
	usecase := usecases.NewGetDrivers(&testLogger, mockRepo)
	ctx := context.Background()
	testError := errors.New("boom")
	mockRepo.EXPECT().GetAllDriverUsers(ctx).Return(nil, testError).Once()

	// --------
	// ACT
	// --------
	users, err := usecase.Execute(ctx)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, testError, err, "Must have same errors")
	assert.Nil(t, users, "User slice must be nil")
}
