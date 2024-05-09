package functional

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/IgorSteps/easypark/internal/adapters/rest/models"
	"github.com/IgorSteps/easypark/tests/functional/client"
	"github.com/IgorSteps/easypark/tests/functional/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type TestDeassignParkingSpaceSuite struct {
	client.RestClientSuite
}

func (s *TestDeassignParkingSpaceSuite) TestDeassignParkingSpace_HappyPath() {
	// --------
	// ASSEMBLE
	// --------
	// Assign a parking space...
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	driver, driverToken := utils.CreateAndLoginDriver(ctx, &s.RestClientSuite, nil)
	adminToken := utils.CreateAndLoginAdmin(ctx, &s.RestClientSuite)

	// Creating a parking lot.
	// All parking spaces statuses are automatically 'available'.
	parkingLot := utils.CreateParkingLot(ctx, adminToken, nil, &s.RestClientSuite)

	// Creating a parking request.
	createRequest := &models.CreateParkingRequestRequest{
		DestinationParkingLotID: parkingLot.ID,
		StartTime:               time.Now().Add(10 * time.Minute),
		EndTime:                 time.Now().Add(20 * time.Minute),
	}
	parkingRequest := utils.CreateParkingRequest(ctx, driverToken, driver.ID, parkingLot.ID, createRequest, &s.RestClientSuite)
	testUpdateRequestSpace := &models.ParkingRequestSpaceUpdateRequest{
		ParkingSpaceID: parkingLot.ParkingSpaces[0].ID,
	}
	// Assign space.
	respBody, respCode, err := s.UpdateParkingRequestSpace(ctx, adminToken, parkingRequest.ID, testUpdateRequestSpace)
	s.Require().NoError(err, "Updating a parking request shouldn't return an error")

	// Assert all went okay.
	s.Require().Equal(http.StatusOK, respCode, "Updating parking request should return 200 code")
	var updateParkingRequestResp models.ParkingRequestSpaceUpdateResponse
	err = s.UnmarshalHTTPResponse(respBody, &updateParkingRequestResp)
	s.Require().NoError(err, "Must not return error")
	s.Require().Equal("successfully assigned a space to a parking request", updateParkingRequestResp.Message, "Response body is wrong.")

	// --------
	// ACT
	// --------
	deassignRequest := &models.ParkingRequestSpaceUpdateRequest{
		ParkingSpaceID: uuid.Nil, // nil
	}
	respBody, respCode, err = s.UpdateParkingRequestSpace(ctx, adminToken, parkingRequest.ID, deassignRequest)
	s.Require().NoError(err, "Updating a parking request shouldn't return an error")

	// --------
	// ASSERT
	// --------
	// Assert all went okay.
	s.Require().Equal(http.StatusOK, respCode, "Updating parking request should return 200 code")
	var deassignResp models.ParkingRequestSpaceUpdateResponse
	err = s.UnmarshalHTTPResponse(respBody, &deassignResp)
	s.Require().NoError(err, "Must not return error")
	s.Require().Equal("successfully deassigned the space from a parking request", deassignResp.Message, "Response body is wrong.")

}

func TestDeassignParkingSpaceSuiteInit(t *testing.T) {
	suite.Run(t, new(TestDeassignParkingSpaceSuite))
}
