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
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

const (
	email    = "mail"
	username = "what"
)

func TestRegisterUser_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockUserRepository := &mocks.UserRepository{}
	usecase := usecases.NewRegisterDriver(testLogger, mockUserRepository)
	ctx := context.Background()

	testUser := CreateTestUser()

	mockUserRepository.EXPECT().CheckUserExists(ctx, email, username).Return(false, nil)
	mockUserRepository.EXPECT().CreateUser(ctx, testUser).Return(nil)

	// --------
	// ACT
	// --------
	err := usecase.Execute(ctx, testUser)

	// --------
	// ASSERT
	// --------
	assert.Nil(t, err, "Error must be nil")

	// Assert logger.
	assert.Equal(t, 0, len(hook.Entries))
	hook.Reset()
	assert.Nil(t, hook.LastEntry())
	mockUserRepository.AssertExpectations(t)
}

func TestRegisterUser_UnhappyPath_UserExists(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockUserRepository := &mocks.UserRepository{}
	usecase := usecases.NewRegisterDriver(testLogger, mockUserRepository)
	ctx := context.Background()

	testUser := CreateTestUser()

	mockUserRepository.EXPECT().CheckUserExists(ctx, email, username).Return(true, nil)

	// --------
	// ACT
	// --------
	err := usecase.Execute(ctx, testUser)

	// --------
	// ASSERT
	// --------
	assert.NotNil(t, err, "Error must not be nil")
	assert.Equal(t, err.Error(), "User 'what'/'mail' already exists")

	// Assert logger.
	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.WarnLevel, hook.LastEntry().Level)
	assert.Equal(t, "user already exists", hook.LastEntry().Message)
	assert.Equal(t, email, hook.LastEntry().Data["email"])

	hook.Reset()
	assert.Nil(t, hook.LastEntry())
	mockUserRepository.AssertExpectations(t)
}

func TestRegisterUser_UnhappyPath_CreateUser_Fails(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockUserRepository := &mocks.UserRepository{}
	usecase := usecases.NewRegisterDriver(testLogger, mockUserRepository)
	ctx := context.Background()

	testError := errors.New("boom")
	testUser := CreateTestUser()

	mockUserRepository.EXPECT().CheckUserExists(ctx, email, username).Return(false, nil)
	mockUserRepository.EXPECT().CreateUser(ctx, testUser).Return(testError)

	// --------
	// ACT
	// --------
	err := usecase.Execute(ctx, testUser)

	// --------
	// ASSERT
	// --------
	assert.NotNil(t, err, "Error must not be nil")
	assert.Equal(t, err, testError)

	// Assert logger.
	assert.Equal(t, 0, len(hook.Entries))
	hook.Reset()
	assert.Nil(t, hook.LastEntry())
	mockUserRepository.AssertExpectations(t)
}

func TestRegisterUser_UnhappyPath_CheckUserExistsByEmail_Fails(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockUserRepository := &mocks.UserRepository{}
	usecase := usecases.NewRegisterDriver(testLogger, mockUserRepository)
	ctx := context.Background()

	testError := errors.New("boom")
	testUser := CreateTestUser()

	mockUserRepository.EXPECT().CheckUserExists(ctx, email, username).Return(false, testError)

	// --------
	// ACT
	// --------
	err := usecase.Execute(ctx, testUser)

	// --------
	// ASSERT
	// --------
	assert.NotNil(t, err, "Error must not be nil")
	assert.Equal(t, err, testError, "Errors are not equal")

	// Assert logger.
	assert.Equal(t, 0, len(hook.Entries))

	hook.Reset()
	assert.Nil(t, hook.LastEntry())
	mockUserRepository.AssertExpectations(t)
}

func CreateTestUser() *entities.User {
	return &entities.User{
		ID:        uuid.New(),
		Username:  username,
		Email:     email,
		Password:  "foo",
		FirstName: "john",
		LastName:  "smith",
		Role:      entities.RoleDriver,
	}

}
