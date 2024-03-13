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
	facade := usecasefacades.NewUserFacade(mockUserCreator)
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
	facade := usecasefacades.NewUserFacade(mockUserCreator)
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
