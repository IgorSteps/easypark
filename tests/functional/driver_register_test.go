package functional

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/IgorSteps/easypark/internal/adapters/rest/models"
	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/tests/functional/client"
	"github.com/stretchr/testify/suite"
)

type TestCreateUserSuite struct {
	client.RestClientSuite
}

// TestCreateUser tests user creation using our REST API.
func (s *TestCreateUserSuite) TestCreateUser_HappyPath() {
	// --------
	// ASSEMBLE
	// --------
	createUserReq := &models.UserCreationRequest{
		Username: "testuser",
		Email:    "testuser@example.com",
		Password: "password12",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// ---
	// ACT
	// ---
	responseBody, responseCode, err := s.CreateUser(ctx, createUserReq)

	// ------
	// ASSERT
	// ------
	s.Require().NoError(err, "Creating user should not return an error")

	var targetModel entities.User
	err = s.UnmarshalHTTPResponse(responseBody, &targetModel)
	if err != nil {
		s.T().Fail()
	}
	s.Require().Equal(createUserReq.Username, targetModel.Username, "Response body username is wrong")
	s.Require().Equal(createUserReq.Email, targetModel.Email, "Response body username is wrong")
	s.Require().Equal(createUserReq.Password, targetModel.Password, "Response body username is wrong")
	s.Require().Equal(http.StatusCreated, responseCode, "Response code is wrong")
}

// TestCreateUser_UnhappyPath_UserAlreadyExists tests user cannot be created using our REST API if another user with the same email or username is present.
func (s *TestCreateUserSuite) TestCreateUser_UnhappyPath_UserAlreadyExists() {
	// --------
	// ASSEMBLE
	// --------
	createUserReq1 := &models.UserCreationRequest{
		Username: "boom",
		Email:    "boom@example.com",
		Password: "password123",
	}

	// request with already present email
	createUserReq2 := &models.UserCreationRequest{
		Username: "gloom",
		Email:    "boom@example.com",
		Password: "password123",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// ---
	// ACT
	// ---
	responseBody, responseCode, err := s.CreateUser(ctx, createUserReq1)
	responseBody2, responseCode2, err2 := s.CreateUser(ctx, createUserReq2)

	// ------
	// ASSERT
	// ------
	s.Require().NoError(err, "Creating user 1 should not return an error")

	var targetModel entities.User
	err = s.UnmarshalHTTPResponse(responseBody, &targetModel)
	if err != nil {
		s.T().Fail()
	}
	s.Require().Equal(createUserReq1.Username, targetModel.Username, "Response body username is wrong")
	s.Require().Equal(createUserReq1.Email, targetModel.Email, "Response body username is wrong")
	s.Require().Equal(createUserReq1.Password, targetModel.Password, "Response body username is wrong")

	s.Require().Equal(http.StatusCreated, responseCode, "Response code 1 is wrong")

	s.Require().NoError(err2, "Creating user 2 should not return an error")
	s.Require().Equal("Resource 'gloom' already exists\n", string(responseBody2), "Response body 2 message is wrong")
	s.Require().Equal(http.StatusBadRequest, responseCode2, "Response code 2 is wrong")
}

func TestInit(t *testing.T) {
	suite.Run(t, new(TestCreateUserSuite))
}
