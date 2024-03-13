package functional

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/IgorSteps/easypark/internal/adapters/rest/models"
	"github.com/IgorSteps/easypark/tests/functional/client"
	"github.com/stretchr/testify/suite"
)

type TestLoginUserSuite struct {
	client.RestClientSuite
}

func (s *TestLoginUserSuite) TestLoginUser_HappyPath() {
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
	// Create user
	_, responseCode, err := s.CreateUser(ctx, createUserReq)
	s.Require().NoError(err, "Creating user should not return an error")
	s.Require().Equal(http.StatusCreated, responseCode, "Response code should be 201")

	// Login user
	responseBody, responseCode, err := s.LoginUser(ctx, loginReq)
	s.Require().NoError(err, "Loging user should not return an error")

	// --------
	// ASSERT
	// --------
	var targetModel models.LoginUserResponse
	err = s.UnmarshalHTTPResponse(responseBody, &targetModel)
	if err != nil {
		s.T().Fail()
	}
	s.Require().Equal(http.StatusOK, responseCode, "Login request should return 200 code")
	s.Require().Equal("User logged in successfully", targetModel.Message, "Response messages don't match")
	s.Require().NotEmpty(targetModel.Token, "Token must not be empty")
}

func (s *TestLoginUserSuite) TestLoginUser_UnhappyPath_InvalidCredentials() {
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
		Password: "differentPassword",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// --------
	// ACT
	// --------
	// Create user
	_, responseCode, err := s.CreateUser(ctx, createUserReq)
	s.Require().NoError(err, "Creating user should not return an error")
	s.Require().Equal(http.StatusCreated, responseCode, "Response code should be 201")

	// Login user
	responseBody, responseCode, err := s.LoginUser(ctx, loginReq)
	s.Require().NoError(err, "Loging user should not return an error")

	// --------
	// ASSERT
	// --------
	s.Require().Equal(http.StatusUnauthorized, responseCode, "Login request should return 400 code")
	s.Require().Equal("Invalid password provided\n", string(responseBody), "Response messages don't match")
}

func (s *TestLoginUserSuite) TestLoginUser_UnhappyPath_UserNotFound() {
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
	diffUsername := "differentUser"
	loginReq := &models.LoginUserRequest{
		Username: diffUsername,
		Password: password,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// --------
	// ACT
	// --------
	// Create user
	_, responseCode, err := s.CreateUser(ctx, createUserReq)
	s.Require().NoError(err, "Creating user should not return an error")
	s.Require().Equal(http.StatusCreated, responseCode, "Response code should be 201")

	// Login user
	responseBody, responseCode, err := s.LoginUser(ctx, loginReq)
	s.Require().NoError(err, "Loging user should not return an error")

	// --------
	// ASSERT
	// --------
	s.Require().Equal(http.StatusUnauthorized, responseCode, "Login request should return 400 code")
	s.Require().Equal(fmt.Sprintf("User '%s' not found\n", diffUsername), string(responseBody), "Response messages don't match")
}

func TestLoginInit(t *testing.T) {
	suite.Run(t, new(TestLoginUserSuite))
}
