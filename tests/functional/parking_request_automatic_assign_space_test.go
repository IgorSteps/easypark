package functional

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/IgorSteps/easypark/internal/adapters/rest/models"
	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/tests/functional/client"
	"github.com/IgorSteps/easypark/tests/functional/utils"
	"github.com/stretchr/testify/suite"
)

type TestAutomaticallyAssignParkingSpace struct {
	client.RestClientSuite
}

// Testing that a parking request can be assigned an available parking space with 0 existing parking requests.
func (s *TestAutomaticallyAssignParkingSpace) TestAutomaticallyAssignParkingSpace_HappyPath_SingleRequest() {
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

	updateRequest := &models.ParkingRequestAutomaticSpaceUpdateRequest{
		ParkingRequestID: parkingRequest.ID,
	}

	// --------
	// ACT
	// --------
	respBody, respCode, err := s.AutomaticallyUpdateParkingRequestSpace(ctx, adminToken, parkingRequest.ID, updateRequest)
	s.Require().NoError(err, "Updating a parking request shouldn't return an error")

	// --------
	// ASSERT
	// --------
	s.Require().Equal(http.StatusOK, respCode, "Updating parking request should return 200 code")

	var assignedParkingSpace entities.ParkingSpace
	err = s.UnmarshalHTTPResponse(respBody, &assignedParkingSpace)
	s.Require().NoError(err, "Must not return error")
	s.Require().NotNil(assignedParkingSpace)
}

// Testing that a parking request can be assigned an available parking space with existing parking request when there is no time slot overlap.
func (s *TestAutomaticallyAssignParkingSpace) TestAutomaticallyAssignParkingSpace_HappyPath_MultipleRequests() {
	// --------
	// ASSEMBLE
	// --------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	driver, driverToken := utils.CreateAndLoginDriver(ctx, &s.RestClientSuite, nil)
	adminToken := utils.CreateAndLoginAdmin(ctx, &s.RestClientSuite)

	// Creating a parking lot.
	// All parking spaces statuses are automatically 'available'.
	createParkingLotReq := &models.CreateParkingLotRequest{
		Name:     "test-lot",
		Capacity: 3,
	}
	parkingLot := utils.CreateParkingLot(ctx, adminToken, createParkingLotReq, &s.RestClientSuite)

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

	updateRequest := &models.ParkingRequestAutomaticSpaceUpdateRequest{
		ParkingRequestID: parkingRequest2.ID,
	}
	// Assign the same space to the second parking request
	respBody2, respCode2, err := s.AutomaticallyUpdateParkingRequestSpace(ctx, adminToken, parkingRequest2.ID, updateRequest) // assign space to the parking request 2
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

	var space entities.ParkingSpace
	err = s.UnmarshalHTTPResponse(respBody2, &space)
	s.Require().NoError(err, "Must not return error")
	s.Require().NotNil(space)
}

func (s *TestAutomaticallyAssignParkingSpace) TestAutomaticallyAssignParkingSpace_UnhappyPath_Overlap() {
	// --------
	// ASSEMBLE
	// --------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	driver, driverToken := utils.CreateAndLoginDriver(ctx, &s.RestClientSuite, nil)
	adminToken := utils.CreateAndLoginAdmin(ctx, &s.RestClientSuite)

	// Creating a parking lot with available spaces.
	createParkingLotReq := &models.CreateParkingLotRequest{
		Name:     "test-lot",
		Capacity: 1,
	}
	parkingLot := utils.CreateParkingLot(ctx, adminToken, createParkingLotReq, &s.RestClientSuite)

	// Setup the times:
	startTime1 := time.Now().Add(10 * time.Minute)
	endTime1 := startTime1.Add(10 * time.Minute)

	// Creating a parking request.
	createRequest1 := &models.CreateParkingRequestRequest{
		DestinationParkingLotID: parkingLot.ID,
		StartTime:               startTime1,
		EndTime:                 endTime1,
	}
	parkingRequest1 := utils.CreateParkingRequest(ctx, driverToken, driver.ID, parkingLot.ID, createRequest1, &s.RestClientSuite)
	updateRequest := &models.ParkingRequestAutomaticSpaceUpdateRequest{
		ParkingRequestID: parkingRequest1.ID,
	}

	// --------
	// ACT
	// --------
	// Assign spaces automatically to the first parking request
	_, respStatus, err := s.AutomaticallyUpdateParkingRequestSpace(ctx, adminToken, parkingRequest1.ID, updateRequest)
	s.Require().NoError(err, "Assigning a space to a parking request shouldn't return an error")
	s.Require().Equal(http.StatusOK, respStatus)

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
	// ASSERT
	// --------
	// Try to assign spaces automatically to other parking requests with overlaps.
	for _, scenario := range scenarios {
		// Creating a new parking request with overlapping times
		req := &models.CreateParkingRequestRequest{
			DestinationParkingLotID: parkingLot.ID,
			StartTime:               scenario.startTime,
			EndTime:                 scenario.endTime,
		}
		parkingRequestOverlap := utils.CreateParkingRequest(ctx, driverToken, driver.ID, parkingLot.ID, req, &s.RestClientSuite)

		// Attempt to assign spaces automatically to the overlapping parking request
		updateRequest := &models.ParkingRequestAutomaticSpaceUpdateRequest{
			ParkingRequestID: parkingRequestOverlap.ID,
		}
		respBody, respStatus, err := s.AutomaticallyUpdateParkingRequestSpace(ctx, adminToken, parkingRequestOverlap.ID, updateRequest)
		s.T().Log(string(respBody))
		s.Require().NoError(err, "Assigning a space to an overlapping parking request should return an error")
		s.Require().Equal(http.StatusBadRequest, respStatus)
		s.Require().Equal("no available parking spaces at the desired time\n", string(respBody))
	}
}

func TestAutomaticallyAssignParkingSpaceInit(t *testing.T) {
	suite.Run(t, new(TestAutomaticallyAssignParkingSpace))
}
