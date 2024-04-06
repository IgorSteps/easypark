package utils

import (
	"bytes"
	"context"
	"net/http"
	"os/exec"
	"time"

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

func CreateAndLoginAdmin(ctx context.Context, s *client.RestClientSuite) string {
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

func CreateAndLoginDriver(ctx context.Context, s *client.RestClientSuite, driver *models.UserCreationRequest) (entities.User, string) {
	if driver == nil {
		driver = &models.UserCreationRequest{
			Username:  "johnDoe",
			Password:  "password",
			Email:     "johnDoe@example.com",
			Firstname: "John",
			Lastname:  "Doe",
		}
	}

	respBody, respCode, err := s.CreateUser(ctx, driver)
	s.Require().NoError(err, "Failed to create driver")
	s.Require().Equal(http.StatusCreated, respCode, "Response code must be 201")
	var targetUserModel entities.User
	err = s.UnmarshalHTTPResponse(respBody, &targetUserModel)
	s.Require().NoError(err, "Failed to unmarshall register response")

	login := &models.LoginUserRequest{
		Username: "johnDoe",
		Password: "password",
	}

	respBody, respCode, err = s.LoginUser(ctx, login)
	s.Require().NoError(err, "Failed to login driver")
	s.Require().Equal(http.StatusOK, respCode, "Response code must be 200")

	var targetLoginModel models.LoginUserResponse
	err = s.UnmarshalHTTPResponse(respBody, &targetLoginModel)
	s.Require().NoError(err, "Failed to unmarshall login response")

	driverJWT := targetLoginModel.Token

	return targetUserModel, driverJWT
}

func CreateParkingRequest(
	ctx context.Context,
	driverToken string,
	driverID uuid.UUID,
	destinationLotID uuid.UUID,
	s *client.RestClientSuite,
) models.CreateParkingRequestResponse {
	createRequest := &models.CreateParkingRequestRequest{
		DestinationParkingLotID: destinationLotID,
		StartTime:               time.Now(),
		EndTime:                 time.Now().Add(555),
	}

	respBody, respCode, err := s.CreateParkingRequest(ctx, driverToken, driverID.String(), createRequest)

	s.Require().NoError(err, "Creating parking request should not return an error")
	var targetModel models.CreateParkingRequestResponse
	err = s.UnmarshalHTTPResponse(respBody, &targetModel)
	if err != nil {
		s.T().Fail()
	}
	s.Require().Equal(http.StatusCreated, respCode, "Response code must be 201")

	return targetModel
}

func CreateParkingLot(ctx context.Context, adminToken string, s *client.RestClientSuite) entities.ParkingLot {
	createRequest := &models.CreateParkingLotRequest{
		Name:     "science",
		Capacity: 10,
	}

	respBody, respCode, err := s.CreateParkingLot(ctx, adminToken, createRequest)
	s.Require().NoError(err, "Failed to create parking lot")
	s.Require().Equal(http.StatusCreated, respCode, "Response code must be 201")

	var targetModel entities.ParkingLot
	err = s.UnmarshalHTTPResponse(respBody, &targetModel)
	s.Require().NoError(err, "Failed to unmarshall create parking lot response")

	return targetModel
}
