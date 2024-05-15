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

type CheckOverStaysTestSuite struct {
	client.RestClientSuite
}

func (s *CheckOverStaysTestSuite) TestCheckOverStays() {
	// --------
	// ASSEMBLE
	// --------
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Create and login an admin
	adminToken := utils.CreateAndLoginAdmin(ctx, &s.RestClientSuite)

	// Create and login a driver
	driver, driverToken := utils.CreateAndLoginDriver(ctx, &s.RestClientSuite, nil)

	// Create a parking lot with a capacity of 5
	createParkingLot := &models.CreateParkingLotRequest{
		Name:     "test-lot",
		Capacity: 5,
	}
	parkingLot := utils.CreateParkingLot(ctx, adminToken, createParkingLot, &s.RestClientSuite)
	parkingSpaceID := parkingLot.ParkingSpaces[0].ID

	// Create a parking request that will end in 3 seconds
	createParkingRequest := &models.CreateParkingRequestRequest{
		DestinationParkingLotID: parkingLot.ID,
		StartTime:               time.Now(),
		EndTime:                 time.Now().Add(3 * time.Second),
	}
	parkingRequest := utils.CreateParkingRequest(ctx, driverToken, driver.ID, parkingLot.ID, createParkingRequest, &s.RestClientSuite)
	utils.AssignParkingSpace(ctx, parkingSpaceID, parkingRequest.ID, adminToken, &s.RestClientSuite)

	// Create another parking request that will end in 2 hours
	createParkingRequest2 := &models.CreateParkingRequestRequest{
		DestinationParkingLotID: parkingLot.ID,
		StartTime:               time.Now(),
		EndTime:                 time.Now().Add(2 * time.Hour),
	}
	parkingRequest2 := utils.CreateParkingRequest(ctx, driverToken, driver.ID, parkingLot.ID, createParkingRequest2, &s.RestClientSuite)
	utils.AssignParkingSpace(ctx, parkingSpaceID, parkingRequest2.ID, adminToken, &s.RestClientSuite)

	// Send the arrival notification for the first parking request
	testRequest := &models.CreateNotificationRequest{
		ParkingRequestID: parkingRequest.ID,
		ParkingSpaceID:   parkingSpaceID,
		Location:         "parkingSpaceName",
		NotificationType: 0, // arrival
	}
	s.CreateNotification(ctx, driverToken, driver.ID, testRequest)

	// Sleep for 10 seconds to simulate passage of time
	time.Sleep(10 * time.Second)

	// Check for overstays with a threshold of 2 seconds
	checksReq := &models.CheckForOverStaysRequest{
		Threshold: 1 * time.Second,
	}
	// --------
	// ACT
	// --------
	respBody, respCode, err := s.CheckForOverStays(ctx, adminToken, checksReq)
	// --------
	// ASSERT
	// --------
	s.Require().NoError(err)
	s.Require().Equal(http.StatusOK, respCode)

	var alerts []entities.Alert
	err = s.UnmarshalHTTPResponse(respBody, &alerts)
	s.Require().NoError(err)
	s.Require().Len(alerts, 1)
	s.Require().Equal(parkingRequest.UserID, alerts[0].UserID)
	s.Require().Equal(parkingSpaceID, alerts[0].ParkingSpaceID)
}

func TestCheckOverStaysTestSuiteInit(t *testing.T) {
	suite.Run(t, new(CheckOverStaysTestSuite))
}
