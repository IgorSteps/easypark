package postgresql_test

import (
	"context"
	"errors"
	"testing"

	"github.com/IgorSteps/easypark/internal/adapters/postgresql"
	"github.com/IgorSteps/easypark/internal/domain/entities"
	mocks "github.com/IgorSteps/easypark/mocks/adapters/postgresql"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestUserRepository_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockDBHandler := &mocks.DBHandler{}
	repository := postgresql.NewPostgreSQLUserRepository(mockDBHandler, testLogger)
	ctx := context.Background()

	testUser := entities.User{
		ID:        uuid.New(),
		Username:  "boo",
		Email:     "email",
		Password:  "foo",
		FirstName: "john",
		LastName:  "smith",
		Role:      entities.Driver,
	}

	mockDBHandler.EXPECT().Create(ctx, &testUser).Return(&gorm.DB{Error: nil}).Once()

	// --------
	// ACT
	// --------
	user, err := repository.CreateUser(ctx, testUser)

	// --------
	// ASSERT
	// --------
	assert.Nil(t, err, "Error must be nil")
	assert.NotEmpty(t, user, "User must be returned and not empty")

	// Assert logger.
	assert.Equal(t, 0, len(hook.Entries))
	hook.Reset()
	assert.Nil(t, hook.LastEntry())
}

func TestUserRepository_UnhappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockDBHandler := &mocks.DBHandler{}
	repository := postgresql.NewPostgreSQLUserRepository(mockDBHandler, testLogger)
	ctx := context.Background()
	testError := errors.New("boom")
	testUser := entities.User{
		ID:        uuid.New(),
		Username:  "boo",
		Email:     "email",
		Password:  "foo",
		FirstName: "john",
		LastName:  "smith",
		Role:      entities.Driver,
	}

	mockDBHandler.EXPECT().Create(ctx, &testUser).Return(&gorm.DB{Error: testError}).Once()

	// --------
	// ACT
	// --------
	user, err := repository.CreateUser(ctx, testUser)

	// --------
	// ASSERT
	// --------
	assert.NotNil(t, err, "Error must not be nil")
	assert.Empty(t, user, "User must be empty")

	// Assert logger.
	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "failed to insert user into the database", hook.LastEntry().Message, "Messages are not equal")

	hook.Reset()
	assert.Nil(t, hook.LastEntry())
}
