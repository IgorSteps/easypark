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

	// Create admin and driver
	adminToken := utils.CreateAndLoginAdmin(ctx, &s.RestClientSuite)
	driver, driverToken := utils.CreateAndLoginDriver(ctx, &s.RestClientSuite, nil)

	// Create a parking lot
	createParkingLot := &models.CreateParkingLotRequest{
		Name:     "test-lot",
		Capacity: 5,
	}
	parkingLot := utils.CreateParkingLot(ctx, adminToken, createParkingLot, &s.RestClientSuite)
	parkingSpaceID := parkingLot.ParkingSpaces[0].ID

	// Create parking request that will trigger the alert
	createParkingRequest := &models.CreateParkingRequestRequest{
		DestinationParkingLotID: parkingLot.ID,
		// Times are a little bit in the future, because we don't allow creation of parking request with times in the past.
		StartTime: time.Now().Add(3 * time.Second), // 3 seconds in the future
		EndTime:   time.Now().Add(5 * time.Second), // 5 seconds in the future
	}
	parkingRequest := utils.CreateParkingRequest(ctx, driverToken, driver.ID, parkingLot.ID, createParkingRequest, &s.RestClientSuite)
	// Assign that parking request a space we chose above.
	utils.AssignParkingSpace(ctx, parkingSpaceID, parkingRequest.ID, adminToken, &s.RestClientSuite)

	// Create a request in the distant future that won't trigger the alert.
	createParkingRequest2 := &models.CreateParkingRequestRequest{
		DestinationParkingLotID: parkingLot.ID,
		// Times are a little bit in the future, because we don't allow creation of parking request with times in the past.
		StartTime: time.Now(),                    //
		EndTime:   time.Now().Add(1 * time.Hour), //
	}
	parkingRequest2 := utils.CreateParkingRequest(ctx, driverToken, driver.ID, parkingLot.ID, createParkingRequest2, &s.RestClientSuite)
	// Assign that parking request a space we chose above.
	utils.AssignParkingSpace(ctx, parkingSpaceID, parkingRequest2.ID, adminToken, &s.RestClientSuite)

	// Sleep for 10 secs.
	time.Sleep(10 * time.Second)

	// Don't send the 'completed' notification.
	// Call the endpoint:
	checksReq := &models.CheckForOverStaysRequest{
		Threshold: 2 * time.Second, // 2 seconds
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
	// Should return 1 alert about 1 park request
	s.Require().NoError(err)
	s.Require().Len(alerts, 1)
	s.Require().Equal(parkingRequest.UserID, alerts[0].UserID)
	s.Require().Equal(parkingSpaceID, alerts[0].ParkingSpaceID)
}

func TestCheckOverStaysTestSuiteInit(t *testing.T) {
	suite.Run(t, new(CheckOverStaysTestSuite))
}
