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

type TestUpdateParkingRequestSpaceSuite struct {
	client.RestClientSuite
}

func (s *TestUpdateParkingRequestSpaceSuite) TestUpdateParkingRequestSpace_HappyPath() {
	// --------
	// ASSEMBLE
	// --------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Populate db with mock data.
	err := PopulateUsers(ctx, &s.RestClientSuite)
	s.Require().NoError(err, "Populating system with mock user data shouldn't return an error")
	userID, adminToken, userToken := GetUserIDAndToken(ctx, &s.RestClientSuite)

	// Creating a parking lot.
	testCreateLotRequest := &models.CreateParkingLotRequest{
		Name:     "science",
		Capacity: 10,
	}
	respBody, respCode, err := s.CreateParkingLot(ctx, adminToken, testCreateLotRequest)
	s.Require().NoError(err, "Creating a parking lot shouldn't return an error")
	s.Require().Equal(http.StatusCreated, respCode, "Create parking request should return 201 code")

	var targetModel models.CreateParkingLotResponse
	err = s.UnmarshalHTTPResponse(respBody, &targetModel)
	s.Require().NoError(err, "Must not return error")

	// Creating a parking request.
	testRequest := &models.CreateParkingRequestRequest{
		DestinationParkingLotID: targetModel.ID,
		StartTime:               time.Now(),
		EndTime:                 time.Now().Add(555),
	}
	respBody, respCode, err = s.CreateParkingRequest(ctx, userToken, userID.String(), testRequest)
	s.Require().NoError(err, "Creating a parking request shouldn't return an error")
	s.Require().Equal(http.StatusCreated, respCode, "Create parking request should return 201 code")

	var createParkingRequestResponseModel models.CreateParkingRequestResponse
	err = s.UnmarshalHTTPResponse(respBody, &createParkingRequestResponseModel)
	s.Require().NoError(err, "Must not return error")

	testUpdateRequestSpace := &models.ParkingRequestSpaceUpdateRequest{
		ParkingSpaceID: targetModel.PakringSpaces[0].ID,
	}

	// --------
	// ACT
	// --------
	respBody, respCode, err = s.UpdateParkingRequestSpace(ctx, adminToken, createParkingRequestResponseModel.ID, testUpdateRequestSpace)
	s.Require().NoError(err, "Updating a parking request shouldn't return an error")

	// --------
	// ASSERT
	// --------
	s.Require().Equal(http.StatusOK, respCode, "Updating parking request should return 200 code")

	var updateParkingRequestResp models.UpdateParkingRequestStatusResponse
	err = s.UnmarshalHTTPResponse(respBody, &updateParkingRequestResp)
	s.Require().NoError(err, "Must not return error")
	s.Require().Equal("successfully assigned a space to a parking request", updateParkingRequestResp.Message, "Response body is wrong.")
}

func TestUpdateParkingRequestSpaceSuiteInit(t *testing.T) {
	suite.Run(t, new(TestUpdateParkingRequestSpaceSuite))
}
