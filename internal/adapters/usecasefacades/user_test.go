package usecasefacades_test

import (
	"context"
	"errors"
	"testing"

	"github.com/IgorSteps/easypark/internal/adapters/usecasefacades"
	"github.com/IgorSteps/easypark/internal/domain/entities"
	mocks "github.com/IgorSteps/easypark/mocks/adapters/usecasefacades"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUsecaseFacade_CreateUser_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	mockUserCreator := &mocks.UserCreator{}
	mockUserAuthenticator := &mocks.UserAuthenticator{}
	facade := usecasefacades.NewUserFacade(mockUserCreator, mockUserAuthenticator)
	ctx := context.Background()
	testUser := &entities.User{
		ID:        uuid.New(),
		Username:  "boo",
		Email:     "email",
		Password:  "foo",
		FirstName: "john",
		LastName:  "smith",
	}

	mockUserCreator.EXPECT().Execute(ctx, testUser).Return(nil).Once()

	// --------
	// ACT
	// --------
	err := facade.CreateUser(ctx, testUser)

	// --------
	// ASSERT
	// --------
	assert.Nil(t, err, "Error must be nil")
	mockUserCreator.AssertExpectations(t)
}

func TestUsecasefacade_CreateUser_UnhappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	mockUserCreator := &mocks.UserCreator{}
	mockUserAuthenticator := &mocks.UserAuthenticator{}
	facade := usecasefacades.NewUserFacade(mockUserCreator, mockUserAuthenticator)
	ctx := context.Background()
	testUser := &entities.User{
		ID:        uuid.New(),
		Username:  "boo",
		Email:     "email",
		Password:  "foo",
		FirstName: "john",
		LastName:  "smith",
	}
	testError := errors.New("boom")

	mockUserCreator.EXPECT().Execute(ctx, testUser).Return(testError).Once()

	// --------
	// ACT
	// --------
	err := facade.CreateUser(ctx, testUser)

	// --------
	// ASSERT
	// --------
	assert.NotNil(t, err, "Error must not be nil")
	assert.Equal(t, testError, err, "Expected and actual errors don't match")
	mockUserCreator.AssertExpectations(t)
}

func TestUsecasefacade_AuthoriseUser_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	mockUserCreator := &mocks.UserCreator{}
	mockUserAuthenticator := &mocks.UserAuthenticator{}
	facade := usecasefacades.NewUserFacade(mockUserCreator, mockUserAuthenticator)
	ctx := context.Background()
	testEmail := "tmail"
	testPwd := "tpwd"
	token := "token"

	mockUserAuthenticator.EXPECT().Execute(ctx, testEmail, testPwd).Return(token, nil).Once()

	// --------
	// ACT
	// --------
	actualToken, err := facade.AuthoriseUser(ctx, testEmail, testPwd)

	// --------
	// ASSERT
	// --------
	assert.Nil(t, err, "Error must be nil")
	assert.Equal(t, token, actualToken, "Expected and actual tokens don't match")
	mockUserAuthenticator.AssertExpectations(t)
}

func TestUsecasefacade_AuthoriseUser_UnhappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	mockUserCreator := &mocks.UserCreator{}
	mockUserAuthenticator := &mocks.UserAuthenticator{}
	facade := usecasefacades.NewUserFacade(mockUserCreator, mockUserAuthenticator)
	ctx := context.Background()
	testEmail := "tmail"
	testPwd := "tpwd"
	emptyToken := ""
	testErr := errors.New("boom")
	mockUserAuthenticator.EXPECT().Execute(ctx, testEmail, testPwd).Return(emptyToken, testErr).Once()

	// --------
	// ACT
	// --------
	actualToken, err := facade.AuthoriseUser(ctx, testEmail, testPwd)

	// --------
	// ASSERT
	// --------
	assert.NotNil(t, err, "Error must not be nil")
	assert.EqualError(t, err, "boom", "Error message is wrong")
	assert.Empty(t, actualToken, "Token must be empty")
	mockUserAuthenticator.AssertExpectations(t)
}
