package functional

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/IgorSteps/easypark/internal/adapters/rest/models"
	"github.com/IgorSteps/easypark/tests/functional/client"
	"github.com/IgorSteps/easypark/tests/functional/utils"
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

	driver, driverToken := utils.CreateAndLoginDriver(ctx, &s.RestClientSuite, nil)
	adminToken := utils.CreateAndLoginAdmin(ctx, &s.RestClientSuite)

	// Creating a parking lot.
	parkingLot := utils.CreateParkingLot(ctx, adminToken, nil, &s.RestClientSuite)

	// Creating a parking request.
	parkingRequest := utils.CreateParkingRequest(ctx, driverToken, driver.ID, parkingLot.ID, nil, &s.RestClientSuite)

	testUpdateRequestSpace := &models.ParkingRequestSpaceUpdateRequest{
		ParkingSpaceID: parkingLot.ParkingSpaces[0].ID,
	}

	// --------
	// ACT
	// --------
	respBody, respCode, err := s.UpdateParkingRequestSpace(ctx, adminToken, parkingRequest.ID, testUpdateRequestSpace)
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

func (s *TestUpdateParkingRequestSpaceSuite) TestUpdateParkingRequestSpace_HappyPath_MultipleParkingRequests() {
	// --------
	// ASSEMBLE
	// --------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	driver, driverToken := utils.CreateAndLoginDriver(ctx, &s.RestClientSuite, nil)
	adminToken := utils.CreateAndLoginAdmin(ctx, &s.RestClientSuite)

	// Creating a parking lot.
	parkingLot := utils.CreateParkingLot(ctx, adminToken, nil, &s.RestClientSuite)

	// First parking request (will be approved)
	startTime1 := time.Now()                  // now
	endTime1 := startTime1.Add(3 * time.Hour) // three hours from now
	createParkingRequestReq1 := &models.CreateParkingRequestRequest{
		DestinationParkingLotID: parkingLot.ID,
		StartTime:               startTime1,
		EndTime:                 endTime1,
	}
	parkingRequest1 := utils.CreateParkingRequest(ctx, driverToken, driver.ID, parkingLot.ID, createParkingRequestReq1, &s.RestClientSuite)

	// Assign a space to parkingRequest1.
	assignSpaceReq := &models.ParkingRequestSpaceUpdateRequest{
		ParkingSpaceID: parkingLot.ParkingSpaces[0].ID,
	}
	utils.UpdateParkingRequestSpace(ctx, adminToken, parkingRequest1.ID, assignSpaceReq, &s.RestClientSuite)

	// Approve parkingRequest1.
	approveRequest := &models.UpdateParkingRequestStatusRequest{
		Status: "approved",
	}
	utils.UpdateParkingRequestStatus(ctx, approveRequest, adminToken, parkingRequest1.ID, &s.RestClientSuite)

	// Second parking request(must not cause an overlap issue)
	// First request is a 3 hour window from now.
	startTime2 := startTime1.Add(3 * time.Hour) // Right after the end time of 1st park request
	endTime2 := startTime2.Add(2 * time.Hour)   // Will not overlap
	createParkingRequestReq2 := &models.CreateParkingRequestRequest{
		DestinationParkingLotID: parkingLot.ID,
		StartTime:               startTime2,
		EndTime:                 endTime2,
	}
	parkingRequest2 := utils.CreateParkingRequest(ctx, driverToken, driver.ID, parkingLot.ID, createParkingRequestReq2, &s.RestClientSuite)

	// Try to assign the same space to parkingRequest2.
	assignTheSameSpaceReq := &models.ParkingRequestSpaceUpdateRequest{
		ParkingSpaceID: parkingLot.ParkingSpaces[0].ID,
	}

	// --------
	// ACT
	// --------
	respBody, respCode, err := s.UpdateParkingRequestSpace(ctx, adminToken, parkingRequest2.ID, assignTheSameSpaceReq)
	s.Require().NoError(err, "Updating a parking request shouldn't return an error")

	// --------
	// ASSERT
	// --------
	s.Require().Equal(http.StatusOK, respCode, "Must return 200 code")
	var updateParkingRequestResp models.ParkingRequestSpaceUpdateResponse
	err = s.UnmarshalHTTPResponse(respBody, &updateParkingRequestResp)
	s.Require().NoError(err, "Must not return error")
	s.Require().Equal("successfully assigned a space to a parking request", updateParkingRequestResp.Message, "Response body is wrong.")
}

func (s *TestUpdateParkingRequestSpaceSuite) TestUpdateParkingRequestSpace_UnhappyPath_MultipleParkingRequestsOverlap() {
	// --------
	// ASSEMBLE
	// --------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	driver, driverToken := utils.CreateAndLoginDriver(ctx, &s.RestClientSuite, nil)
	adminToken := utils.CreateAndLoginAdmin(ctx, &s.RestClientSuite)

	// Creating a parking lot.
	parkingLot := utils.CreateParkingLot(ctx, adminToken, nil, &s.RestClientSuite)

	// First parking request (will be approved)
	startTime1 := time.Now()                  // now
	endTime1 := startTime1.Add(3 * time.Hour) // three hours from now
	createParkingRequestReq1 := &models.CreateParkingRequestRequest{
		DestinationParkingLotID: parkingLot.ID,
		StartTime:               startTime1,
		EndTime:                 endTime1,
	}
	parkingRequest1 := utils.CreateParkingRequest(ctx, driverToken, driver.ID, parkingLot.ID, createParkingRequestReq1, &s.RestClientSuite)

	// Assign a space to parkingRequest1.
	assignSpaceReq := &models.ParkingRequestSpaceUpdateRequest{
		ParkingSpaceID: parkingLot.ParkingSpaces[0].ID,
	}
	utils.UpdateParkingRequestSpace(ctx, adminToken, parkingRequest1.ID, assignSpaceReq, &s.RestClientSuite)

	// Approve parkingRequest1.
	approveRequest := &models.UpdateParkingRequestStatusRequest{
		Status: "approved",
	}
	utils.UpdateParkingRequestStatus(ctx, approveRequest, adminToken, parkingRequest1.ID, &s.RestClientSuite)

	// Second parking request(will cause an overlap issue)
	// First reques is a 3 hour window from now.
	startTime2 := startTime1.Add(1 * time.Hour) // Overlaps the first request
	endTime2 := endTime1.Add(2 * time.Hour)
	createParkingRequestReq2 := &models.CreateParkingRequestRequest{
		DestinationParkingLotID: parkingLot.ID,
		StartTime:               startTime2,
		EndTime:                 endTime2,
	}
	parkingRequest2 := utils.CreateParkingRequest(ctx, driverToken, driver.ID, parkingLot.ID, createParkingRequestReq2, &s.RestClientSuite)

	// Try to assign the same space to parkingRequest2.
	assignTheSameSpaceReq := &models.ParkingRequestSpaceUpdateRequest{
		ParkingSpaceID: parkingLot.ParkingSpaces[0].ID,
	}

	// --------
	// ACT
	// --------
	respBody, respCode, err := s.UpdateParkingRequestSpace(ctx, adminToken, parkingRequest2.ID, assignTheSameSpaceReq)
	s.Require().NoError(err, "Updating a parking request shouldn't return an error")

	// --------
	// ASSERT
	// --------
	s.Require().Equal(http.StatusBadRequest, respCode, "Must return 400 code")
	s.Require().Equal("parking space isn't available at the requested time\n", string(respBody), "Response body is wrong.")
}

func TestUpdateParkingRequestSpaceSuiteInit(t *testing.T) {
	suite.Run(t, new(TestUpdateParkingRequestSpaceSuite))
}
