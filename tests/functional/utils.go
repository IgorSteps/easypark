package functional

import (
	"bytes"
	"context"
	"net/http"
	"os/exec"

	"github.com/IgorSteps/easypark/internal/adapters/rest/models"
	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/tests/functional/client"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// PopulateUsers creates multiple users in the database using our REST API.
func PopulateUsers(ctx context.Context, s *client.RestClientSuite) error {
	users := []models.UserCreationRequest{
		{
			Username:  "testuser1",
			Password:  "password1",
			Email:     "test1@example.com",
			Firstname: "Test",
			Lastname:  "User1",
		},
		{
			Username:  "testuser2",
			Password:  "password2",
			Email:     "test2@example.com",
			Firstname: "Test",
			Lastname:  "User2",
		},
	}

	for _, userReq := range users {
		_, statusCode, err := s.CreateUser(ctx, &userReq)
		if err != nil || statusCode != http.StatusCreated {
			s.T().Logf("Failed to create user: %v with status code: %d", err, statusCode)
			return err
		}
	}

	return nil
}

func CreateAdmin(ctx context.Context, s *client.RestClientSuite) string {
	cmd := exec.Command("sh", "../../build/createadmin.sh")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		s.T().Fatalf("Failed to create admin user: %v, Stderr %s", err, stderr.String())
	}

	// Must match what's in the above shell script.
	loginReq := &models.LoginUserRequest{
		Username: "adminUsername",
		Password: "securePassword",
	}
	respBody, _, err := s.LoginUser(ctx, loginReq)
	assert.NoError(s.T(), err, "Creating admin shouldn't return an error")

	// Unmarshal response to get the admin's token.
	var targetModel models.LoginUserResponse
	err = s.UnmarshalHTTPResponse(respBody, &targetModel)
	if err != nil {
		s.T().Fatalf("Failed to unmarshall login response %v", err)
	}

	s.T().Log("Got admin token")
	return targetModel.Token
}

func GetUserIDAndToken(ctx context.Context, s *client.RestClientSuite) (uuid.UUID, string) {
	// Create admin and get their auth token:
	token := CreateAdmin(ctx, s)

	// Get all drivers.
	respBody, respCode, err := s.GetAllDrivers(ctx, token)
	s.Require().NoError(err, "Getting all driver's must not return an error")
	s.Require().Equal(http.StatusOK, respCode, "Response codemust be 200")

	// Unmarshall response to get driver users.
	var users []entities.User
	err = s.UnmarshalHTTPResponse(respBody, &users)
	if err != nil {
		s.T().Fail()
	}

	// Extract one of the users ids.
	userID := users[0].ID

	// Extract user's uname and pword
	uname := users[0].Username
	pwod := users[0].Password

	// Get the auth token for them
	loginReq := models.LoginUserRequest{
		Username: uname,
		Password: pwod,
	}
	respBody, respCode, err = s.LoginUser(ctx, &loginReq)

	// Assert everything went okay.
	s.Require().NoError(err, "Login in user shouldn't error")
	s.Require().Equal(http.StatusOK, respCode, "Response code must be OK")

	var targetModel models.LoginUserResponse
	err = s.UnmarshalHTTPResponse(respBody, &targetModel)
	if err != nil {
		s.T().Fail()
	}

	userToken := targetModel.Token

	return userID, userToken
}
