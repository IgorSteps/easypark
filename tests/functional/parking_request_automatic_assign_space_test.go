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

func TestAutomaticallyAssignParkingSpaceInit(t *testing.T) {
	suite.Run(t, new(TestAutomaticallyAssignParkingSpace))
}
