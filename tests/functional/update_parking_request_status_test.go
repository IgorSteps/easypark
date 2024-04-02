package functional

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/IgorSteps/easypark/internal/adapters/rest/models"
	"github.com/IgorSteps/easypark/tests/functional/client"
	"github.com/stretchr/testify/suite"
)

type TestApproveParkingRequestSuite struct {
	client.RestClientSuite
}

func (s *TestApproveParkingRequestSuite) TestApproveParkingRequest_HappyPath() {
	// --------
	// ASSEMBLE
	// --------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Populate db with mock data.
	err := PopulateUsers(ctx, &s.RestClientSuite)
	s.Require().NoError(err, "Populating system with mock user data shouldn't return an error")
	userID, adminToken, userToken := GetUserIDAndToken(ctx, &s.RestClientSuite)

	testRequest := &models.CreateParkingRequestRequest{
		Destination: "science",
		StartTime:   time.Now(),
		EndTime:     time.Now().Add(555),
	}

	// Creating a parking request.
	respBody, respCode, err := s.CreateParkingRequest(ctx, userToken, userID.String(), testRequest)
	s.Require().NoError(err, "Creating a parking request shouldn't return an error")
	s.Require().Equal(http.StatusCreated, respCode, "Create parking request should return 201 code")

	// Unmarshall response body.
	var targetModel models.CreateParkingRequestResponse
	err = s.UnmarshalHTTPResponse(respBody, &targetModel)
	if err != nil {
		s.T().Fail()
	}
	parkingRequestID := targetModel.ID

	updateRequst := &models.UpdateParkingRequestStatusRequest{
		Status: "approved",
	}

	// --------
	// ACT
	// --------
	respBody, respCode, err = s.UpdateParkingRequestStatus(ctx, adminToken, parkingRequestID.String(), updateRequst)

	// --------
	// ASSERT
	// --------
	s.Require().NoError(err, "Updating status of the parking request shouldn't return error")
	s.Require().Equal(http.StatusOK, respCode, "Must return 200")

	// Unmarshall response body.
	var tModel models.UpdateParkingRequestStatusResponse
	err = s.UnmarshalHTTPResponse(respBody, &tModel)
	if err != nil {
		s.T().Fail()
	}
	s.Require().Equal("successfully updated parking request status", tModel.Message, "Resposne message is wrong")
}

func (s *TestApproveParkingRequestSuite) TestRejectParkingRequest_HappyPath() {
	// --------
	// ASSEMBLE
	// --------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Populate db with mock data.
	err := PopulateUsers(ctx, &s.RestClientSuite)
	s.Require().NoError(err, "Populating system with mock user data shouldn't return an error")
	userID, adminToken, userToken := GetUserIDAndToken(ctx, &s.RestClientSuite)

	testRequest := &models.CreateParkingRequestRequest{
		Destination: "science",
		StartTime:   time.Now(),
		EndTime:     time.Now().Add(555),
	}

	// Creating a parking request.
	respBody, respCode, err := s.CreateParkingRequest(ctx, userToken, userID.String(), testRequest)
	s.Require().NoError(err, "Creating a parking request shouldn't return an error")
	s.Require().Equal(http.StatusCreated, respCode, "Create parking request should return 201 code")

	// Unmarshall response body.
	var targetModel models.CreateParkingRequestResponse
	err = s.UnmarshalHTTPResponse(respBody, &targetModel)
	if err != nil {
		s.T().Fail()
	}
	parkingRequestID := targetModel.ID

	updateRequst := &models.UpdateParkingRequestStatusRequest{
		Status: "rejected",
	}

	// --------
	// ACT
	// --------
	respBody, respCode, err = s.UpdateParkingRequestStatus(ctx, adminToken, parkingRequestID.String(), updateRequst)

	// --------
	// ASSERT
	// --------
	s.Require().NoError(err, "Updating status of the parking request shouldn't return error")
	s.Require().Equal(http.StatusOK, respCode, "Must return 200")

	// Unmarshall response body.
	var tModel models.UpdateParkingRequestStatusResponse
	err = s.UnmarshalHTTPResponse(respBody, &tModel)
	if err != nil {
		s.T().Fail()
	}
	s.Require().Equal("successfully updated parking request status", tModel.Message, "Response message is wrong")
}

func (s *TestApproveParkingRequestSuite) TestUpdateParkingRequest_UnhappyPath_UnknownStatus() {
	// --------
	// ASSEMBLE
	// --------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Populate db with mock data.
	err := PopulateUsers(ctx, &s.RestClientSuite)
	s.Require().NoError(err, "Populating system with mock user data shouldn't return an error")
	userID, adminToken, userToken := GetUserIDAndToken(ctx, &s.RestClientSuite)

	testRequest := &models.CreateParkingRequestRequest{
		Destination: "science",
		StartTime:   time.Now(),
		EndTime:     time.Now().Add(555),
	}

	// Creating a parking request.
	respBody, respCode, err := s.CreateParkingRequest(ctx, userToken, userID.String(), testRequest)
	s.Require().NoError(err, "Creating a parking request shouldn't return an error")
	s.Require().Equal(http.StatusCreated, respCode, "Create parking request should return 201 code")

	// Unmarshall response body.
	var targetModel models.CreateParkingRequestResponse
	err = s.UnmarshalHTTPResponse(respBody, &targetModel)
	if err != nil {
		s.T().Fail()
	}
	parkingRequestID := targetModel.ID

	updateRequst := &models.UpdateParkingRequestStatusRequest{
		Status: "boom",
	}

	// --------
	// ACT
	// --------
	respBody, respCode, err = s.UpdateParkingRequestStatus(ctx, adminToken, parkingRequestID.String(), updateRequst)

	// --------
	// ASSERT
	// --------
	s.Require().NoError(err, "Updating status of the parking request shouldn't return error")
	s.Require().Equal(http.StatusBadRequest, respCode, "Must return 400")
	s.Require().Equal("unknown parking request status\n", string(respBody), "Response message is wrong")
}

func TestApproveParkingRequestSuiteInit(t *testing.T) {
	suite.Run(t, new(TestApproveParkingRequestSuite))
}
