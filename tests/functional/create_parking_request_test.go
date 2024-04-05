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

type TestCreateParkingRequest struct {
	client.RestClientSuite
}

func (s *TestCreateParkingRequest) TestCreateParkingRequest_HappyPath() {
	// --------
	// ASSEMBLE
	// --------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Populate db with mock data.
	err := PopulateUsers(ctx, &s.RestClientSuite)
	s.Require().NoError(err, "Populating system with mock user data shouldn't return an error")
	userID, _, userToken := GetUserIDAndToken(ctx, &s.RestClientSuite)

	testRequest := &models.CreateParkingRequestRequest{
		DestinationParkingLotID: uuid.New(),
		StartTime:               time.Now(),
		EndTime:                 time.Now().Add(555),
	}

	// --------
	// ACT
	// --------
	respBody, respCode, err := s.CreateParkingRequest(ctx, userToken, userID.String(), testRequest)

	// --------
	// ASSERT
	// --------
	s.Require().NoError(err, "Creating parking request should not return an error")
	var targetModel models.CreateParkingRequestResponse
	err = s.UnmarshalHTTPResponse(respBody, &targetModel)
	if err != nil {
		s.T().Fail()
	}

	//s.Require().Equal("science", targetModel.Destination, "Response body destination is wrong")
	s.Require().Equal(entities.RequestStatusPending, targetModel.Status, "Request status must be pending")
	s.Require().Equal(http.StatusCreated, respCode, "Response code is wrong")
}

func (s *TestCreateParkingRequest) TestCreateParkingRequest_UnhappyPath_InvalidInput() {
	// --------
	// ASSEMBLE
	// --------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Populate db with mock data.
	err := PopulateUsers(ctx, &s.RestClientSuite)
	s.Require().NoError(err, "Populating system with mock user data shouldn't return an error")
	userID, _, userToken := GetUserIDAndToken(ctx, &s.RestClientSuite)

	testRequest := &models.CreateParkingRequestRequest{
		DestinationParkingLotID: uuid.New(),
		StartTime:               time.Now().Add(500000),
		EndTime:                 time.Now(),
	}

	// --------
	// ACT
	// --------
	respBody, respCode, err := s.CreateParkingRequest(ctx, userToken, userID.String(), testRequest)

	// --------
	// ASSERT
	// --------
	s.Require().NoError(err, "Creating parking request should not return an error")
	s.Require().Equal("start time cannot be after the end time\n", string(respBody), "Response body is wrong")
	s.Require().Equal(http.StatusBadRequest, respCode, "Response code is wrong")
}

func TestCreateParkingRequestInit(t *testing.T) {
	suite.Run(t, new(TestCreateParkingRequest))
}
