package datastore_test

import (
	"context"
	"errors"
	"testing"

	"github.com/IgorSteps/easypark/internal/adapters/datastore"
	"github.com/IgorSteps/easypark/internal/domain/entities"
	mocks "github.com/IgorSteps/easypark/mocks/adapters/datastore"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestUserRepository_CreateUser_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockDatastore := &mocks.Datastore{}
	repository := datastore.NewUserPostgresRepository(mockDatastore, testLogger)
	ctx := context.Background()

	testUser := CreateTestUser()

	mockDatastore.EXPECT().Create(ctx, &testUser).Return(&gorm.DB{Error: nil}).Once()

	// --------
	// ACT
	// --------
	err := repository.CreateUser(ctx, testUser)

	// --------
	// ASSERT
	// --------
	assert.Nil(t, err, "Error must be nil")

	// Assert logger.
	assert.Equal(t, 0, len(hook.Entries))
	hook.Reset()
	assert.Nil(t, hook.LastEntry())
}

func Test_UserRepository_CreateUser_UnhappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockDatastore := &mocks.Datastore{}
	repository := datastore.NewUserPostgresRepository(mockDatastore, testLogger)
	ctx := context.Background()
	testError := errors.New("boom")
	testUser := CreateTestUser()

	mockDatastore.EXPECT().Create(ctx, &testUser).Return(&gorm.DB{Error: testError}).Once()

	// --------
	// ACT
	// --------
	err := repository.CreateUser(ctx, testUser)

	// --------
	// ASSERT
	// --------
	assert.NotNil(t, err, "Error must not be nil")

	// Assert logger.
	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "failed to insert user into the database", hook.LastEntry().Message, "Messages are not equal")

	hook.Reset()
	assert.Nil(t, hook.LastEntry())
}

func CreateTestUser() *entities.User {
	return &entities.User{
		ID:        uuid.New(),
		Username:  "boo",
		Email:     "email",
		Password:  "foo",
		FirstName: "john",
		LastName:  "smith",
		Role:      entities.RoleDriver,
	}

}
