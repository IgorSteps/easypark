package functional

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/IgorSteps/easypark/internal/adapters/rest/models"
	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/tests/functional/client"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type TestRBACSuite struct {
	client.RestClientSuite
}

func (s *TestRBACSuite) TestRBAC_HappyPath_DriverAccessesDriverRoutes() {
	// --------
	// ASSEMBLE
	// --------
	username := "testuser"
	password := "password"
	createUserReq := &models.UserCreationRequest{
		Username: username,
		Email:    "testuser@example.com",
		Password: password,
	}
	loginReq := &models.LoginUserRequest{
		Username: username,
		Password: password,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// --------
	// ACT
	// --------
	// Create driver
	respBody, responseCode, err := s.CreateUser(ctx, createUserReq)
	s.Require().NoError(err, "Creating user should not return an error")
	s.Require().Equal(http.StatusCreated, responseCode, "Response code should be 201")

	// Unmarshal response to get the user.
	var targetUserModel entities.User
	err = s.UnmarshalHTTPResponse(respBody, &targetUserModel)
	if err != nil {
		s.T().Fail()
	}

	// Login driver
	responseBody, responseCode, err := s.LoginUser(ctx, loginReq)
	s.Require().NoError(err, "Loging user should not return an error")

	// Unmarshal response to get the token.
	var targetModel models.LoginUserResponse
	err = s.UnmarshalHTTPResponse(responseBody, &targetModel)
	if err != nil {

		s.T().Fail()
	}
	req := &models.CreateParkingRequestRequest{
		DestinationParkingLotID: uuid.New(),
		StartTime:               time.Now().Add(5 * time.Minute),
		EndTime:                 time.Now().Add(15 * time.Minute),
	}
	// Make request to driver route
	_, respC, err := s.CreateParkingRequest(ctx, targetModel.Token, targetUserModel.ID.String(), req)
	s.Require().NoError(err, "Making request to driver route should not return an error")

	// --------
	// ASSERT
	// --------
	s.Require().Equal(http.StatusOK, responseCode, "Login request should return 200 code")
	s.Require().NotEmpty(targetModel.User, "User must not be empty")
	s.Require().NotEmpty(targetModel.Token, "Token must not be empty")

	s.Require().Equal(http.StatusCreated, respC, "request to create park request should return 201")
}

func (s *TestRBACSuite) TestRBAC_HappyPath_DriverCannotAccessAdminRoutes() {
	// --------
	// ASSEMBLE
	// --------
	username := "testuser"
	password := "password"
	createUserReq := &models.UserCreationRequest{
		Username: username,
		Email:    "testuser@example.com",
		Password: password,
	}
	loginReq := &models.LoginUserRequest{
		Username: username,
		Password: password,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// --------
	// ACT
	// --------
	// Create driver
	_, responseCode, err := s.CreateUser(ctx, createUserReq)
	s.Require().NoError(err, "Creating user should not return an error")
	s.Require().Equal(http.StatusCreated, responseCode, "Response code should be 201")

	// Login driver
	responseBody, responseCode, err := s.LoginUser(ctx, loginReq)
	s.Require().NoError(err, "Loging user should not return an error")

	// Unmarshal response to get the token.
	var targetModel models.LoginUserResponse
	err = s.UnmarshalHTTPResponse(responseBody, &targetModel)
	if err != nil {
		s.T().Fail()
	}
	s.T().Log(targetModel.Token)
	// Make request to driver route
	respB, respC, err := s.GetAllDrivers(ctx, targetModel.Token)
	s.Require().NoError(err, "Making request to driver route should not return an error")

	// --------
	// ASSERT
	// --------
	s.Require().Equal(http.StatusOK, responseCode, "Login request should return 200 code")
	s.Require().NotEmpty(targetModel.User, "User must not be empty")
	s.Require().NotEmpty(targetModel.Token, "Token must not be empty")

	s.Require().Equal(http.StatusForbidden, respC, "Request to admin route by driver hould return 403")
	s.Require().Equal("Forbidden\n", string(respB))
}

func TestRBACInit(t *testing.T) {
	suite.Run(t, new(TestRBACSuite))
}
