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

type TestAssignParkingSpace struct {
	client.RestClientSuite
}

// Testing that a parking request can be assigned an available parking space with 0 existing parking requests.
func (s *TestAssignParkingSpace) TestAssignParkingSpace_HappyPath_SingleRequest() {
	// --------
	// ASSEMBLE
	// --------
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

	// --------
	// ACT
	// --------
	respBody, respCode, err := s.UpdateParkingRequestSpace(ctx, adminToken, parkingRequest.ID, testUpdateRequestSpace)
	s.Require().NoError(err, "Updating a parking request shouldn't return an error")

	// --------
	// ASSERT
	// --------
	s.Require().Equal(http.StatusOK, respCode, "Updating parking request should return 200 code")

	var updateParkingRequestResp models.ParkingRequestSpaceUpdateResponse
	err = s.UnmarshalHTTPResponse(respBody, &updateParkingRequestResp)
	s.Require().NoError(err, "Must not return error")
	s.Require().Equal("successfully assigned a space to a parking request", updateParkingRequestResp.Message, "Response body is wrong.")
}

// Testing that a parking request can be assigned an available parking space with existing parking request when there is no time slot overlap.
func (s *TestAssignParkingSpace) TestAssignParkingSpace_HappyPath_MultipleRequests() {
	// --------
	// ASSEMBLE
	// --------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	driver, driverToken := utils.CreateAndLoginDriver(ctx, &s.RestClientSuite, nil)
	adminToken := utils.CreateAndLoginAdmin(ctx, &s.RestClientSuite)

	// Creating a parking lot.
	// All parking spaces statuses are automatically 'available'.
	parkingLot := utils.CreateParkingLot(ctx, adminToken, nil, &s.RestClientSuite)

	// Creating a parking request 1.
	createRequest1 := &models.CreateParkingRequestRequest{
		DestinationParkingLotID: parkingLot.ID,
		StartTime:               time.Now().Add(10 * time.Minute),
		EndTime:                 time.Now().Add(20 * time.Minute),
	}
	parkingRequest1 := utils.CreateParkingRequest(ctx, driverToken, driver.ID, parkingLot.ID, createRequest1, &s.RestClientSuite)

	// Creating a parking request 2, with times after the first one that shouldn't cause an overlap.
	createRequest2 := &models.CreateParkingRequestRequest{
		DestinationParkingLotID: parkingLot.ID,
		StartTime:               time.Now().Add(20 * time.Minute),
		EndTime:                 time.Now().Add(30 * time.Minute),
	}
	parkingRequest2 := utils.CreateParkingRequest(ctx, driverToken, driver.ID, parkingLot.ID, createRequest2, &s.RestClientSuite)

	testUpdateRequestSpace := &models.ParkingRequestSpaceUpdateRequest{
		ParkingSpaceID: parkingLot.ParkingSpaces[0].ID, // same parking space
	}

	// --------
	// ACT
	// --------
	// Assign this space to the first parking request
	respBody1, respCode1, err := s.UpdateParkingRequestSpace(ctx, adminToken, parkingRequest1.ID, testUpdateRequestSpace) // assign space to the parking request 1
	s.Require().NoError(err, "Assigning a space to a parking request shouldn't return an error")

	// Assign the same space to the second parking request
	respBody2, respCode2, err := s.UpdateParkingRequestSpace(ctx, adminToken, parkingRequest2.ID, testUpdateRequestSpace) // assign space to the parking request 2
	s.Require().NoError(err, "Assigning a space to a parking request shouldn't return an error")

	// --------
	// ASSERT
	// --------
	s.Require().Equal(http.StatusOK, respCode1, "Assigning a space to a parking request 1 should return 200 code")
	s.Require().Equal(http.StatusOK, respCode2, "Assigning a space to a parking request 2 should return 200 code")

	var updateParkingRequestResp1 models.ParkingRequestSpaceUpdateResponse
	err = s.UnmarshalHTTPResponse(respBody1, &updateParkingRequestResp1)
	s.Require().NoError(err, "Must not return error")
	s.Require().Equal("successfully assigned a space to a parking request", updateParkingRequestResp1.Message, "Response body of parking space 1 is wrong.")

	var updateParkingRequestResp2 models.ParkingRequestSpaceUpdateResponse
	err = s.UnmarshalHTTPResponse(respBody2, &updateParkingRequestResp2)
	s.Require().NoError(err, "Must not return error")
	s.Require().Equal("successfully assigned a space to a parking request", updateParkingRequestResp2.Message, "Response body of parking space 2 is wrong.")
}

func (s *TestAssignParkingSpace) TestAssignParkingSpace_UnhappyPath_Overlap() {
	// --------
	// ASSEMBLE
	// --------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	driver, driverToken := utils.CreateAndLoginDriver(ctx, &s.RestClientSuite, nil)
	adminToken := utils.CreateAndLoginAdmin(ctx, &s.RestClientSuite)

	// Creating a parking lot.
	// All parking spaces statuses are automatically 'available'.
	parkingLot := utils.CreateParkingLot(ctx, adminToken, nil, &s.RestClientSuite)

	// Setup the times:
	startTime1 := time.Now().Add(10 * time.Minute)
	endTime1 := startTime1.Add(10 * time.Minute)

	// Creating a parking request 1.
	createRequest1 := &models.CreateParkingRequestRequest{
		DestinationParkingLotID: parkingLot.ID,
		StartTime:               startTime1,
		EndTime:                 endTime1,
	}
	parkingRequest1 := utils.CreateParkingRequest(ctx, driverToken, driver.ID, parkingLot.ID, createRequest1, &s.RestClientSuite)

	testUpdateRequestSpace := &models.ParkingRequestSpaceUpdateRequest{
		ParkingSpaceID: parkingLot.ParkingSpaces[0].ID, // same parking space
	}

	// Different scenarios for overlapping parking requests.
	scenarios := []struct {
		startTime time.Time
		endTime   time.Time
	}{
		{startTime: startTime1.Add(-5 * time.Minute), endTime: startTime1.Add(5 * time.Minute)}, // Starts before and ends during request 1
		{startTime: startTime1.Add(5 * time.Minute), endTime: endTime1.Add(5 * time.Minute)},    // Starts during and ends after request 1
		{startTime: startTime1.Add(-5 * time.Minute), endTime: endTime1.Add(5 * time.Minute)},   // Completely encompasses request 1
		{startTime: startTime1, endTime: endTime1},                                              // Same start and end time as request 1
	}

	// --------
	// ACT
	// --------
	// Assign this space to the first parking request
	respBody1, respCode1, err := s.UpdateParkingRequestSpace(ctx, adminToken, parkingRequest1.ID, testUpdateRequestSpace) // assign space to the parking request 1
	s.Require().NoError(err, "Assigning a space to a parking request shouldn't return an error")

	// --------
	// ASSERT
	// --------
	s.Require().Equal(http.StatusOK, respCode1, "Assigning a space to a parking request 1 should return 200 code")
	var updateParkingRequestResp1 models.ParkingRequestSpaceUpdateResponse
	err = s.UnmarshalHTTPResponse(respBody1, &updateParkingRequestResp1)
	s.Require().NoError(err, "Must not return error")
	s.Require().Equal("successfully assigned a space to a parking request", updateParkingRequestResp1.Message, "Response body of parking space 1 is wrong.")

	// Try assign to other parking requests with overlaps.
	for _, scenario := range scenarios {
		// Creating a new parking request with overlapping times
		req := &models.CreateParkingRequestRequest{
			DestinationParkingLotID: parkingLot.ID,
			StartTime:               scenario.startTime,
			EndTime:                 scenario.endTime,
		}
		parkingRequestOverlap := utils.CreateParkingRequest(ctx, driverToken, driver.ID, parkingLot.ID, req, &s.RestClientSuite)

		// Attempt to assign the same space to the overlapping parking request
		_, respCode, _ := s.UpdateParkingRequestSpace(ctx, adminToken, parkingRequestOverlap.ID, testUpdateRequestSpace)
		s.Require().Equal(http.StatusBadRequest, respCode, "Assigning a space to an overlapping parking request should return 400 code")
	}
}

// Testing that a rejected parking request cannot be assigned a parking space.
func (s *TestAssignParkingSpace) TestAssignParkingSpace_UnhappyPath_RejectedRequest() {
	// --------
	// ASSEMBLE
	// --------
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

	// Reject this parking request
	utils.RejectParkingRequest(ctx, adminToken, parkingRequest.ID, &s.RestClientSuite)

	testUpdateRequestSpace := &models.ParkingRequestSpaceUpdateRequest{
		ParkingSpaceID: parkingLot.ParkingSpaces[0].ID,
	}

	// --------
	// ACT
	// --------
	respBody, respCode, err := s.UpdateParkingRequestSpace(ctx, adminToken, parkingRequest.ID, testUpdateRequestSpace)
	s.Require().NoError(err, "Updating a parking request shouldn't return an error")

	s.Require().Equal(http.StatusBadRequest, respCode, "Updating parking request should return 200 code")
	s.Require().Equal("not allowed to assign parking space to a 'rejected' parking request\n", string(respBody), "Response body is wrong.")
}

// Testing that a parking request with desired start time in the past cannot be assigned a space.
func (s *TestAssignParkingSpace) TestAssignParkingSpace_UnhappyPath_OutdatedRequest() {
	// --------
	// ASSEMBLE
	// --------
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
		StartTime:               time.Now().Add(-10 * time.Minute), // time is in the past
		EndTime:                 time.Now().Add(-5 * time.Minute),
	}
	parkingRequest := utils.CreateParkingRequest(ctx, driverToken, driver.ID, parkingLot.ID, createRequest, &s.RestClientSuite)

	testUpdateRequestSpace := &models.ParkingRequestSpaceUpdateRequest{
		ParkingSpaceID: parkingLot.ParkingSpaces[0].ID,
	}

	// --------
	// ACT
	// --------
	respBody, respCode, err := s.UpdateParkingRequestSpace(ctx, adminToken, parkingRequest.ID, testUpdateRequestSpace)
	s.Require().NoError(err, "Updating a parking request shouldn't return an error")

	s.Require().Equal(http.StatusBadRequest, respCode, "Updating parking request should return 200 code")
	s.Require().Equal("not allowed to assign a parking space to a parking request with the desired start time in the past\n", string(respBody), "Response body is wrong.")
}

// Testing that a parking request cannot be assigned a blocked space.
func (s *TestAssignParkingSpace) TestAssignParkingSpace_UnhappyPath_BlockedSpace() {
	// --------
	// ASSEMBLE
	// --------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	driver, driverToken := utils.CreateAndLoginDriver(ctx, &s.RestClientSuite, nil)
	adminToken := utils.CreateAndLoginAdmin(ctx, &s.RestClientSuite)

	// Creating a parking lot.
	// All parking spaces statuses are automatically 'available'.
	parkingLot := utils.CreateParkingLot(ctx, adminToken, nil, &s.RestClientSuite)
	// Block parking space.
	blockedPatrkingSpaceID := parkingLot.ParkingSpaces[0].ID
	utils.BlockParkingSpace(ctx, adminToken, blockedPatrkingSpaceID, &s.RestClientSuite)

	// Creating a parking request.
	createRequest := &models.CreateParkingRequestRequest{
		DestinationParkingLotID: parkingLot.ID,
		StartTime:               time.Now().Add(10 * time.Minute),
		EndTime:                 time.Now().Add(15 * time.Minute),
	}
	parkingRequest := utils.CreateParkingRequest(ctx, driverToken, driver.ID, parkingLot.ID, createRequest, &s.RestClientSuite)

	testUpdateRequestSpace := &models.ParkingRequestSpaceUpdateRequest{
		ParkingSpaceID: blockedPatrkingSpaceID,
	}

	// --------
	// ACT
	// --------
	respBody, respCode, err := s.UpdateParkingRequestSpace(ctx, adminToken, parkingRequest.ID, testUpdateRequestSpace)
	s.Require().NoError(err, "Updating a parking request shouldn't return an error")

	s.Require().Equal(http.StatusBadRequest, respCode, "Updating parking request should return 200 code")
	s.Require().Equal("not allowed to assign blocked parking space\n", string(respBody), "Response body is wrong.")
}

func TestAssignParkingSpaceInit(t *testing.T) {
	suite.Run(t, new(TestAssignParkingSpace))
}
