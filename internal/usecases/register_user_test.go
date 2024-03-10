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

const email = "mail"

func TestRegisterUser_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockUserRepository := &mocks.UserRepository{}
	usecase := usecases.NewRegisterUser(testLogger, mockUserRepository)
	ctx := context.Background()

	testUser := entities.User{
		ID:        uuid.New(),
		Username:  "boo",
		Email:     email,
		Password:  "foo",
		FirstName: "john",
		LastName:  "smith",
		Role:      entities.Driver,
	}

	mockUserRepository.EXPECT().CheckUserExistsByEmail(ctx, email).Return(false, nil)
	mockUserRepository.EXPECT().CreateUser(ctx, testUser).Return(testUser, nil)

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
}

func TestRegisterUser_UnhappyPath_UserExists(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockUserRepository := &mocks.UserRepository{}
	usecase := usecases.NewRegisterUser(testLogger, mockUserRepository)
	ctx := context.Background()

	testUser := entities.User{
		ID:        uuid.New(),
		Username:  "boo",
		Email:     email,
		Password:  "foo",
		FirstName: "john",
		LastName:  "smith",
		Role:      entities.Driver,
	}

	mockUserRepository.EXPECT().CheckUserExistsByEmail(ctx, email).Return(true, nil)

	// --------
	// ACT
	// --------
	err := usecase.Execute(ctx, testUser)

	// --------
	// ASSERT
	// --------
	assert.NotNil(t, err, "Error must not be nil")
	assert.Equal(t, err.Error(), "user already exists")

	// Assert logger.
	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.WarnLevel, hook.LastEntry().Level)
	assert.Equal(t, "user already exists", hook.LastEntry().Message)
	assert.Equal(t, email, hook.LastEntry().Data["email"])

	hook.Reset()
	assert.Nil(t, hook.LastEntry())
}

func TestRegisterUser_UnhappyPath_CreateUser_Fails(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockUserRepository := &mocks.UserRepository{}
	usecase := usecases.NewRegisterUser(testLogger, mockUserRepository)
	ctx := context.Background()

	testError := errors.New("boom")
	testUser := entities.User{
		ID:        uuid.New(),
		Username:  "boo",
		Email:     email,
		Password:  "foo",
		FirstName: "john",
		LastName:  "smith",
		Role:      entities.Driver,
	}

	mockUserRepository.EXPECT().CheckUserExistsByEmail(ctx, email).Return(false, nil)
	mockUserRepository.EXPECT().CreateUser(ctx, testUser).Return(testUser, testError)

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
}

func TestRegisterUser_UnhappyPath_CheckUserExistsByEmail_Fails(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockUserRepository := &mocks.UserRepository{}
	usecase := usecases.NewRegisterUser(testLogger, mockUserRepository)
	ctx := context.Background()

	testError := errors.New("boom")
	testUser := entities.User{
		ID:        uuid.New(),
		Username:  "boo",
		Email:     email,
		Password:  "foo",
		FirstName: "john",
		LastName:  "smith",
		Role:      entities.Driver,
	}

	mockUserRepository.EXPECT().CheckUserExistsByEmail(ctx, email).Return(false, testError)

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
	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "failed to get user by email", hook.LastEntry().Message, "Messages are not equal")
	assert.Equal(t, email, hook.LastEntry().Data["email"])
	assert.Equal(t, testError.Error(), hook.LastEntry().Data["error"], "Error in the logger fields is not equal")

	hook.Reset()
	assert.Nil(t, hook.LastEntry())
}
