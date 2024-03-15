package usecases_test

import (
	"context"
	"errors"
	"testing"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/IgorSteps/easypark/internal/usecases"
	repositoriesMocks "github.com/IgorSteps/easypark/mocks/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticateUser_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	ctx := context.Background()
	logger, _ := test.NewNullLogger()
	mockRepo := &repositoriesMocks.UserRepository{}
	mockTokenService := &repositoriesMocks.TokenRepository{}
	usecase := usecases.NewAuthenticateUser(logger, mockRepo, mockTokenService)
	testUsername := "uname"
	testPassword := "pwd"
	testToken := "token"
	user := &entities.User{
		ID:        uuid.New(),
		Username:  testUsername,
		Password:  testPassword,
		Email:     "email",
		FirstName: "john",
		LastName:  "smith",
		Role:      entities.RoleDriver,
	}

	mockRepo.EXPECT().FindByUsername(ctx, testUsername).Return(user, nil).Once()
	mockTokenService.EXPECT().GenerateToken(user).Return(testToken, nil).Once()

	// --------
	// ACT
	// --------
	token, err := usecase.Execute(ctx, testUsername, testPassword)

	// --------
	// ASSERT
	// --------
	assert.Nil(t, err, "Error must be nil")
	assert.Equal(t, testToken, token, "Token's don't match")
	mockRepo.AssertExpectations(t)
	mockTokenService.AssertExpectations(t)
}

func TestAuthenticateUser_UnhappyPath_FindByUsername(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	ctx := context.Background()
	logger, _ := test.NewNullLogger()
	mockRepo := &repositoriesMocks.UserRepository{}
	mockTokenService := &repositoriesMocks.TokenRepository{}
	usecase := usecases.NewAuthenticateUser(logger, mockRepo, mockTokenService)
	testUsername := "uname"
	testPassword := "pwd"
	user := &entities.User{}
	testError := errors.New("boom")

	mockRepo.EXPECT().FindByUsername(ctx, testUsername).Return(user, testError).Once()

	// --------
	// ACT
	// --------
	token, err := usecase.Execute(ctx, testUsername, testPassword)

	// --------
	// ASSERT
	// --------
	assert.NotNil(t, err, "Error must not be nil")
	assert.Empty(t, token, "Token must be empty")
	assert.Equal(t, err, testError, "Error's don't match")
	mockRepo.AssertExpectations(t)
}

func TestAuthenticateUser_UnhappyPath_Credentials(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	ctx := context.Background()
	logger, hook := test.NewNullLogger()
	mockRepo := &repositoriesMocks.UserRepository{}
	mockTokenService := &repositoriesMocks.TokenRepository{}
	usecase := usecases.NewAuthenticateUser(logger, mockRepo, mockTokenService)
	testUsername := "uname"
	testPassword := "pwd"
	user := &entities.User{
		ID:        uuid.New(),
		Username:  testUsername,
		Password:  "differentPassword",
		Email:     "email",
		FirstName: "john",
		LastName:  "smith",
		Role:      entities.RoleDriver,
	}
	testError := errors.New("Invalid password provided")

	mockRepo.EXPECT().FindByUsername(ctx, testUsername).Return(user, nil).Once()

	// --------
	// ACT
	// --------
	token, err := usecase.Execute(ctx, testUsername, testPassword)

	// --------
	// ASSERT
	// --------
	assert.NotNil(t, err, "Error must not be nil")
	assert.Empty(t, token, "Token must be empty")
	assert.IsType(t, &repositories.InvalidCredentialsError{}, err, "Error returned is of wrong type")
	assert.Equal(t, err.Error(), testError.Error(), "Error's don't match")

	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.WarnLevel, hook.LastEntry().Level)
	assert.Equal(t, "provided invalid credentials", hook.LastEntry().Message)
	assert.Equal(t, hook.LastEntry().Data["username"], testUsername)
	hook.Reset()
	assert.Nil(t, hook.LastEntry())

	mockRepo.AssertExpectations(t)
}

func TestAuthenticateUser_UnhappyPath_Token(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	ctx := context.Background()
	logger, hook := test.NewNullLogger()
	mockRepo := &repositoriesMocks.UserRepository{}
	mockTokenService := &repositoriesMocks.TokenRepository{}
	usecase := usecases.NewAuthenticateUser(logger, mockRepo, mockTokenService)
	testUsername := "uname"
	testPassword := "pwd"
	user := &entities.User{
		ID:        uuid.New(),
		Username:  testUsername,
		Password:  testPassword,
		Email:     "email",
		FirstName: "john",
		LastName:  "smith",
		Role:      entities.RoleDriver,
	}
	testError := errors.New("Internal error: failed to generate auth token")

	mockRepo.EXPECT().FindByUsername(ctx, testUsername).Return(user, nil).Once()
	mockTokenService.EXPECT().GenerateToken(user).Return("", testError).Once()

	// --------
	// ACT
	// --------
	token, err := usecase.Execute(ctx, testUsername, testPassword)

	// --------
	// ASSERT
	// --------
	assert.NotNil(t, err, "Error must not be nil")
	assert.Empty(t, token, "Token must be empty")
	assert.IsType(t, &repositories.InternalError{}, err, "Error returned is of wrong type")
	assert.Equal(t, err.Error(), testError.Error(), "Error's don't match")

	// Assert logger.
	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "failed to generate auth token", hook.LastEntry().Message)
	hook.Reset()
	assert.Nil(t, hook.LastEntry())

	mockTokenService.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}
