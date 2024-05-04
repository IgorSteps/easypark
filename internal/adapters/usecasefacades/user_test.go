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

type UsecaseFacadeTestSuite struct {
	mockDriverCreator  *mocks.DriverCreator
	mockUserAuthoriser *mocks.UserAuthenticator
	mockDriversGetter  *mocks.DriversGetter
	mockDriverBanner   *mocks.DriverBanner
}

func NewUsecaseFacadeTestSuite() *UsecaseFacadeTestSuite {
	return &UsecaseFacadeTestSuite{
		mockDriverCreator:  &mocks.DriverCreator{},
		mockUserAuthoriser: &mocks.UserAuthenticator{},
		mockDriversGetter:  &mocks.DriversGetter{},
		mockDriverBanner:   &mocks.DriverBanner{},
	}
}

func TestUsecaseFacade_CreateDriver_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	s := NewUsecaseFacadeTestSuite()
	facade := usecasefacades.NewUserFacade(
		s.mockDriverCreator,
		s.mockUserAuthoriser,
		s.mockDriversGetter,
		s.mockDriverBanner,
	)
	ctx := context.Background()
	testUser := &entities.User{
		ID:        uuid.New(),
		Username:  "boo",
		Email:     "email",
		Password:  "foo",
		FirstName: "john",
		LastName:  "smith",
	}

	s.mockDriverCreator.EXPECT().Execute(ctx, testUser).Return(CreateTestUser(), nil).Once()

	// --------
	// ACT
	// --------
	driver, err := facade.CreateDriver(ctx, testUser)

	// --------
	// ASSERT
	// --------
	assert.Nil(t, err, "Error must be nil")
	assert.NotNil(t, driver, "Driver cannot be nil")
	s.mockDriverCreator.AssertExpectations(t)
}

func TestUsecasefacade_CreateDriver_UnhappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	s := NewUsecaseFacadeTestSuite()
	facade := usecasefacades.NewUserFacade(
		s.mockDriverCreator,
		s.mockUserAuthoriser,
		s.mockDriversGetter,
		s.mockDriverBanner,
	)
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

	s.mockDriverCreator.EXPECT().Execute(ctx, testUser).Return(nil, testError).Once()

	// --------
	// ACT
	// --------
	driver, err := facade.CreateDriver(ctx, testUser)

	// --------
	// ASSERT
	// --------
	assert.NotNil(t, err, "Error must not be nil")
	assert.Nil(t, driver, "Driver must be nil")
	assert.Equal(t, testError, err, "Expected and actual errors don't match")
	s.mockDriverCreator.AssertExpectations(t)
}

func TestUsecasefacade_AuthoriseUser_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	s := NewUsecaseFacadeTestSuite()
	facade := usecasefacades.NewUserFacade(
		s.mockDriverCreator,
		s.mockUserAuthoriser,
		s.mockDriversGetter,
		s.mockDriverBanner,
	)
	ctx := context.Background()
	testEmail := "tmail"
	testPwd := "tpwd"
	token := "token"

	s.mockUserAuthoriser.EXPECT().Execute(ctx, testEmail, testPwd).Return(nil, token, nil).Once()

	// --------
	// ACT
	// --------
	_, actualToken, err := facade.AuthoriseUser(ctx, testEmail, testPwd)

	// --------
	// ASSERT
	// --------
	assert.Nil(t, err, "Error must be nil")
	assert.Equal(t, token, actualToken, "Expected and actual tokens don't match")
	s.mockUserAuthoriser.AssertExpectations(t)
}

func TestUsecasefacade_AuthoriseUser_UnhappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	s := NewUsecaseFacadeTestSuite()
	facade := usecasefacades.NewUserFacade(
		s.mockDriverCreator,
		s.mockUserAuthoriser,
		s.mockDriversGetter,
		s.mockDriverBanner,
	)
	ctx := context.Background()
	testEmail := "tmail"
	testPwd := "tpwd"
	emptyToken := ""
	testErr := errors.New("boom")
	s.mockUserAuthoriser.EXPECT().Execute(ctx, testEmail, testPwd).Return(nil, emptyToken, testErr).Once()

	// --------
	// ACT
	// --------
	_, actualToken, err := facade.AuthoriseUser(ctx, testEmail, testPwd)

	// --------
	// ASSERT
	// --------
	assert.NotNil(t, err, "Error must not be nil")
	assert.EqualError(t, err, "boom", "Error message is wrong")
	assert.Empty(t, actualToken, "Token must be empty")
	s.mockUserAuthoriser.AssertExpectations(t)
}

func TestUsecasefacade_GetAllDriverUsers_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	s := NewUsecaseFacadeTestSuite()
	facade := usecasefacades.NewUserFacade(
		s.mockDriverCreator,
		s.mockUserAuthoriser,
		s.mockDriversGetter,
		s.mockDriverBanner,
	)
	ctx := context.Background()
	expectedUsers := []entities.User{
		{Username: "user1", Email: "user1@example.com"},
		{Username: "user2", Email: "user2@example.com"},
	}
	s.mockDriversGetter.EXPECT().Execute(ctx).Return(expectedUsers, nil).Once()

	// --------
	// ACT
	// --------
	users, err := facade.GetAllDriverUsers(ctx)

	// --------
	// ASSERT
	// --------
	assert.Nil(t, err, "Error must be nil")
	assert.Equal(t, expectedUsers, users, "User slices must be the same")
	s.mockDriversGetter.AssertExpectations(t)
}

func TestUsecasefacade_GetAllDriverUsers_UnhappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	s := NewUsecaseFacadeTestSuite()
	facade := usecasefacades.NewUserFacade(
		s.mockDriverCreator,
		s.mockUserAuthoriser,
		s.mockDriversGetter,
		s.mockDriverBanner,
	)
	ctx := context.Background()
	testError := errors.New("boom")

	s.mockDriversGetter.EXPECT().Execute(ctx).Return([]entities.User{}, testError).Once()

	// --------
	// ACT
	// --------
	users, err := facade.GetAllDriverUsers(ctx)

	// --------
	// ASSERT
	// --------
	assert.NotNil(t, err, "Error must not be nil")
	assert.Equal(t, testError, err, "Errors must equal")
	assert.Empty(t, users, "User slice must be empty")
	s.mockDriversGetter.AssertExpectations(t)
}

func TestUsecasefacade_BanDriver_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	s := NewUsecaseFacadeTestSuite()
	facade := usecasefacades.NewUserFacade(
		s.mockDriverCreator,
		s.mockUserAuthoriser,
		s.mockDriversGetter,
		s.mockDriverBanner,
	)
	ctx := context.Background()
	testID := uuid.New()

	s.mockDriverBanner.EXPECT().Execute(ctx, testID).Return(nil).Once()

	// --------
	// ACT
	// --------
	err := facade.BanDriver(ctx, testID)

	// --------
	// ASSERT
	// --------
	assert.Nil(t, err, "Error must be nil")
	s.mockDriverBanner.AssertExpectations(t)
}

func TestUsecasefacade_BanDriver_UnhappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	s := NewUsecaseFacadeTestSuite()
	facade := usecasefacades.NewUserFacade(
		s.mockDriverCreator,
		s.mockUserAuthoriser,
		s.mockDriversGetter,
		s.mockDriverBanner,
	)
	testError := errors.New("boom")
	ctx := context.Background()
	testID := uuid.New()

	s.mockDriverBanner.EXPECT().Execute(ctx, testID).Return(testError).Once()

	// --------
	// ACT
	// --------
	err := facade.BanDriver(ctx, testID)

	// --------
	// ASSERT
	// --------
	assert.NotNil(t, err, "Error must be nil")
	assert.EqualError(t, err, testError.Error())
	s.mockDriverBanner.AssertExpectations(t)
}

func CreateTestUser() *entities.User {
	return &entities.User{
		ID:        uuid.New(),
		Username:  "boom",
		Email:     "bom",
		Password:  "foo",
		FirstName: "john",
		LastName:  "smith",
		Role:      entities.RoleDriver,
	}
}
