package datastore_test

import (
	"context"
	"errors"
	"testing"

	"github.com/IgorSteps/easypark/internal/adapters/datastore"
	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	mocks "github.com/IgorSteps/easypark/mocks/adapters/datastore"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

	mockDatastore.EXPECT().WithContext(ctx).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Create(testUser).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Error().Return(nil).Once()

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

	mockDatastore.EXPECT().WithContext(ctx).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Create(testUser).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Error().Return(testError).Once()

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
	mockDatastore.AssertExpectations(t)
}

func Test_UserRepository_CheckUserExists_HappyPath_UserFound(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockDatastore := &mocks.Datastore{}
	repository := datastore.NewUserPostgresRepository(mockDatastore, testLogger)
	ctx := context.Background()

	query := "email = ? OR username = ?"
	testEmail := "email"
	testUsername := "username"

	mockDatastore.EXPECT().WithContext(ctx).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Where(query, testEmail, testUsername).Return(mockDatastore).Once()
	mockDatastore.EXPECT().First(&entities.User{}).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Error().Return(nil).Once()

	// --------
	// ACT
	// --------
	exist, err := repository.CheckUserExists(ctx, testEmail, testUsername)

	// --------
	// ASSERT
	// --------
	assert.Nil(t, err, "Error must be nil")
	assert.True(t, exist, "Must exist")
	mockDatastore.AssertExpectations(t)
}

func Test_UserRepository_CheckUserExists_HappyPath_UserNotFound(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockDatastore := &mocks.Datastore{}
	repository := datastore.NewUserPostgresRepository(mockDatastore, testLogger)
	ctx := context.Background()

	query := "email = ? OR username = ?"
	testEmail := "email"
	testUsername := "username"

	mockDatastore.EXPECT().WithContext(ctx).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Where(query, testEmail, testUsername).Return(mockDatastore).Once()
	mockDatastore.EXPECT().First(&entities.User{}).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Error().Return(gorm.ErrRecordNotFound).Once()

	// --------
	// ACT
	// --------
	exist, err := repository.CheckUserExists(ctx, testEmail, testUsername)

	// --------
	// ASSERT
	// --------
	assert.Nil(t, err, "Error must be nil")
	assert.False(t, exist, "Must not exist")
	mockDatastore.AssertExpectations(t)
}

func Test_UserRepository_CheckUserExists_UnhappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockDatastore := &mocks.Datastore{}
	repository := datastore.NewUserPostgresRepository(mockDatastore, testLogger)
	ctx := context.Background()

	testError := repositories.NewInternalError("failed to query for user in the database")
	query := "email = ? OR username = ?"
	testEmail := "email"
	testUsername := "username"

	mockDatastore.EXPECT().WithContext(ctx).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Where(query, testEmail, testUsername).Return(mockDatastore).Once()
	mockDatastore.EXPECT().First(&entities.User{}).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Error().Return(testError).Once()

	// --------
	// ACT
	// --------
	exist, err := repository.CheckUserExists(ctx, testEmail, testUsername)

	// --------
	// ASSERT
	// --------
	assert.NotNil(t, err, "Error must not be nil")
	assert.False(t, exist, "Must not exist")
	assert.Equal(t, testError, err, "Errors don't match")

	// Assert logger.
	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "failed to query for user in the database", hook.LastEntry().Message, "Messages are not equal")
	assert.Equal(t, "username", hook.LastEntry().Data["username"])
	assert.Equal(t, "email", hook.LastEntry().Data["email"])
	hook.Reset()
	assert.Nil(t, hook.LastEntry())

	mockDatastore.AssertExpectations(t)
}

func Test_UserRepository_FindByUsername_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockDatastore := &mocks.Datastore{}
	repository := datastore.NewUserPostgresRepository(mockDatastore, testLogger)
	ctx := context.Background()

	query := "username = ?"
	testUsername := "username"

	mockDatastore.EXPECT().WithContext(ctx).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Where(query, testUsername).Return(mockDatastore).Once()
	mockDatastore.EXPECT().First(&entities.User{}).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Error().Return(nil).Once()

	// --------
	// ACT
	// --------
	user, err := repository.GetDriverByUsername(ctx, testUsername)

	// --------
	// ASSERT
	// --------
	assert.Nil(t, err, "Error must be nil")
	assert.NotNil(t, user, "User must not be nil")
	mockDatastore.AssertExpectations(t)
}

func Test_UserRepository_FindByUsername_UnhappyPath_NotFound(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockDatastore := &mocks.Datastore{}
	repository := datastore.NewUserPostgresRepository(mockDatastore, testLogger)
	ctx := context.Background()

	query := "username = ?"
	testUsername := "username"

	mockDatastore.EXPECT().WithContext(ctx).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Where(query, testUsername).Return(mockDatastore).Once()
	mockDatastore.EXPECT().First(&entities.User{}).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Error().Return(gorm.ErrRecordNotFound).Once()

	// --------
	// ACT
	// --------
	user, err := repository.GetDriverByUsername(ctx, testUsername)

	// --------
	// ASSERT
	// --------
	assert.NotNil(t, err, "Error must not be nil")
	assert.EqualError(t, err, "Resource 'username' not found")
	assert.Empty(t, user, "User must be empty")
	mockDatastore.AssertExpectations(t)
}

func Test_UserRepository_FindByUsername_UnhappyPath_FailedToQuery(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockDatastore := &mocks.Datastore{}
	repository := datastore.NewUserPostgresRepository(mockDatastore, testLogger)
	ctx := context.Background()

	query := "username = ?"
	testUsername := "username"
	testError := repositories.NewInternalError("boom")

	mockDatastore.EXPECT().WithContext(ctx).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Where(query, testUsername).Return(mockDatastore).Once()
	mockDatastore.EXPECT().First(&entities.User{}).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Error().Return(testError).Once()

	// --------
	// ACT
	// --------
	user, err := repository.GetDriverByUsername(ctx, testUsername)

	// --------
	// ASSERT
	// --------
	assert.NotNil(t, err, "Error must not be nil")
	assert.EqualError(t, err, "Internal error: failed to query for user in the database")
	assert.Empty(t, user, "User must be empty")
	mockDatastore.AssertExpectations(t)
}

func Test_UserRepository_GetAllDriverUsers_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockDatastore := &mocks.Datastore{}
	repository := datastore.NewUserPostgresRepository(mockDatastore, testLogger)
	ctx := context.Background()

	var usersParam []entities.User
	expectedUsers := []entities.User{
		{Username: "user1", Email: "user1@example.com"},
		{Username: "user2", Email: "user2@example.com"},
	}
	mockDatastore.EXPECT().WithContext(ctx).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Where("role <> ?", entities.RoleAdmin).Return(mockDatastore).Once()
	mockDatastore.EXPECT().FindAll(&usersParam).Return(mockDatastore).Once().Run(func(args mock.Arguments) {
		arg := args.Get(0).(*[]entities.User) // Get the argument passed to FindAll
		*arg = expectedUsers                  // Set it to the expected users
	})
	mockDatastore.EXPECT().Error().Return(nil).Once()

	// --------
	// ACT
	// --------
	users, err := repository.GetAllDriverUsers(ctx)

	// --------
	// ASSERT
	// --------
	assert.Nil(t, err, "Error must not be nil")
	assert.NotEmpty(t, users, "Users slice must not be empty")
	assert.Equal(t, expectedUsers, users, "Users slices must be equal")
	mockDatastore.AssertExpectations(t)
}

func Test_UserRepository_GetAllDriverUsers_UnhappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockDatastore := &mocks.Datastore{}
	repository := datastore.NewUserPostgresRepository(mockDatastore, testLogger)
	ctx := context.Background()

	testErrror := errors.New("boom")
	var usersParam []entities.User
	expectedUsers := []entities.User{}
	mockDatastore.EXPECT().WithContext(ctx).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Where("role <> ?", entities.RoleAdmin).Return(mockDatastore).Once()
	mockDatastore.EXPECT().FindAll(&usersParam).Return(mockDatastore).Once().Run(func(args mock.Arguments) {
		arg := args.Get(0).(*[]entities.User) // Get the argument passed to FindAll
		*arg = expectedUsers                  // Set it to the expected users
	})
	mockDatastore.EXPECT().Error().Return(testErrror).Once()

	// --------
	// ACT
	// --------
	users, err := repository.GetAllDriverUsers(ctx)

	// --------
	// ASSERT
	// --------
	assert.NotNil(t, err, "Error must not be nil")
	assert.IsType(t, &repositories.InternalError{}, err, "Error returned is of wrong type")
	assert.Equal(t, "Internal error: failed to query for all drivers in the database", err.Error(), "Errors must be equal")
	assert.Empty(t, users, "Users slice must be empty")

	// Assert logger
	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "failed to query for all drivers in the database", hook.LastEntry().Message, "Messages are not equal")
	hook.Reset()
	assert.Nil(t, hook.LastEntry())

	mockDatastore.AssertExpectations(t)
}

func Test_UserRepository_GetDriverByID_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockDatastore := &mocks.Datastore{}
	repository := datastore.NewUserPostgresRepository(mockDatastore, testLogger)
	ctx := context.Background()

	var user entities.User
	testID := uuid.New()
	mockDatastore.EXPECT().WithContext(ctx).Return(mockDatastore).Once()
	mockDatastore.EXPECT().First(&user, "id = ?", testID).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Error().Return(nil).Once()

	// --------
	// ACT
	// --------
	err := repository.GetSingle(ctx, testID, &user)

	// --------
	// ASSERT
	// --------
	assert.Nil(t, err, "Error must be nil")
	mockDatastore.AssertExpectations(t)

	// Assert logger
	assert.Equal(t, 0, len(hook.Entries))
}

func Test_UserRepository_GetDriverByID_UnhappyPath_NotFound(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockDatastore := &mocks.Datastore{}
	repository := datastore.NewUserPostgresRepository(mockDatastore, testLogger)
	ctx := context.Background()

	var user entities.User
	testID := uuid.New()
	mockDatastore.EXPECT().WithContext(ctx).Return(mockDatastore).Once()
	mockDatastore.EXPECT().First(&user, "id = ?", testID).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Error().Return(gorm.ErrRecordNotFound).Once()

	// --------
	// ACT
	// --------
	err := repository.GetSingle(ctx, testID, &user)

	// --------
	// ASSERT
	// --------
	assert.NotNil(t, err, "Error must not be nil")
	assert.IsType(t, &repositories.NotFoundError{}, err, "Wrong error type")
	mockDatastore.AssertExpectations(t)

	// Assert logger
	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "failed to find driver with given id in the database", hook.LastEntry().Message, "Messages are not equal")
	assert.Equal(t, testID, hook.LastEntry().Data["id"], "ID field is incorrect")
}

func Test_UserRepository_GetDriverByID_UnhappyPath_InternalError(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockDatastore := &mocks.Datastore{}
	repository := datastore.NewUserPostgresRepository(mockDatastore, testLogger)
	ctx := context.Background()
	testErr := errors.New("boom")

	var user entities.User
	testID := uuid.New()
	mockDatastore.EXPECT().WithContext(ctx).Return(mockDatastore).Once()
	mockDatastore.EXPECT().First(&user, "id = ?", testID).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Error().Return(testErr).Once()

	// --------
	// ACT
	// --------
	err := repository.GetSingle(ctx, testID, &user)

	// --------
	// ASSERT
	// --------
	assert.NotNil(t, err, "Error must not be nil")
	assert.IsType(t, &repositories.InternalError{}, err, "Wrong error type")
	mockDatastore.AssertExpectations(t)

	// Assert logger
	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "failed to query for user in the database", hook.LastEntry().Message, "Messages are not equal")
	assert.Equal(t, testErr, hook.LastEntry().Data["error"], "Error field is incorrect")
}

func Test_UserRepository_Save_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockDatastore := &mocks.Datastore{}
	repository := datastore.NewUserPostgresRepository(mockDatastore, testLogger)
	ctx := context.Background()

	var user entities.User
	mockDatastore.EXPECT().WithContext(ctx).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Save(&user).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Error().Return(nil).Once()

	// --------
	// ACT
	// --------
	err := repository.Save(ctx, &user)

	// --------
	// ASSERT
	// --------
	assert.Nil(t, err, "Error must be nil")
	mockDatastore.AssertExpectations(t)

	// Assert logger
	assert.Equal(t, 0, len(hook.Entries))
}

func Test_UserRepository_Save_UnhappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockDatastore := &mocks.Datastore{}
	repository := datastore.NewUserPostgresRepository(mockDatastore, testLogger)
	ctx := context.Background()

	testErr := errors.New("boom")
	var user entities.User
	mockDatastore.EXPECT().WithContext(ctx).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Save(&user).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Error().Return(testErr).Once()

	// --------
	// ACT
	// --------
	err := repository.Save(ctx, &user)

	// --------
	// ASSERT
	// --------
	assert.NotNil(t, err, "Error must not be nil")
	assert.IsType(t, &repositories.InternalError{}, err, "Wrong error type")
	mockDatastore.AssertExpectations(t)

	// Assert logger
	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "failed to save updated user in the database", hook.LastEntry().Message, "Messages are not equal")
	assert.Equal(t, testErr, hook.LastEntry().Data["error"], "Error field is incorrect")
}

// Utility function to create a test user instance.
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
