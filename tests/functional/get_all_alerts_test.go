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

type GetAllAlertsTestSuite struct {
	client.RestClientSuite
}

func (s *GetAllAlertsTestSuite) TestGetAllAlerts_HappyPath_HTTPEndpoint() {
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
		StartTime: time.Now().Add(3 * time.Second), // 5 seconds in the future
		EndTime:   time.Now().Add(5 * time.Second), // 7 seconds in the future
	}
	parkingRequest := utils.CreateParkingRequest(ctx, driverToken, driver.ID, parkingLot.ID, createParkingRequest, &s.RestClientSuite)
	// Assign that parking request a space we chose above.
	utils.AssignParkingSpace(ctx, parkingSpaceID, parkingRequest.ID, adminToken, &s.RestClientSuite)

	// Create a request in the distant future that won't trigger the alert.
	createParkingRequest2 := &models.CreateParkingRequestRequest{
		DestinationParkingLotID: parkingLot.ID,
		// Times are a little bit in the future, because we don't allow creation of parking request with times in the past.
		StartTime: time.Now().Add(5 * time.Hour), // 5 hours in the future
		EndTime:   time.Now().Add(7 * time.Hour), // 7 hours in the future
	}
	parkingRequest2 := utils.CreateParkingRequest(ctx, driverToken, driver.ID, parkingLot.ID, createParkingRequest2, &s.RestClientSuite)
	// Assign that parking request a space we chose above.
	utils.AssignParkingSpace(ctx, parkingSpaceID, parkingRequest2.ID, adminToken, &s.RestClientSuite)

	// Create another parking request that will trigger the alert
	createParkingRequest3 := &models.CreateParkingRequestRequest{
		DestinationParkingLotID: parkingLot.ID,
		// Times are a little bit in the future, because we don't allow creation of parking request with times in the past.
		StartTime: time.Now().Add(6 * time.Second), // 5 seconds in the future
		EndTime:   time.Now().Add(8 * time.Second), // 7 seconds in the future
	}
	parkingRequest3 := utils.CreateParkingRequest(ctx, driverToken, driver.ID, parkingLot.ID, createParkingRequest3, &s.RestClientSuite)
	// Assign that parking request a space we chose above.
	utils.AssignParkingSpace(ctx, parkingSpaceID, parkingRequest3.ID, adminToken, &s.RestClientSuite)

	// Sleep for 10 secs.
	time.Sleep(10 * time.Second)

	// Don't send the arrival notification.
	// Call the endpoint:
	checksReq := &models.CheckForLateArrivalsRequest{
		Threshold: 1 * time.Second, // 1 seconds
	}

	// --------
	// ACT
	// --------
	respBody, respCode, err := s.CheckForLateArrivals(ctx, adminToken, checksReq)
	s.Require().NoError(err)
	s.Require().Equal(http.StatusOK, respCode)

	respBody, respCode, err = s.GetAllAlerts(ctx, adminToken)

	// --------
	// ASSERT
	// --------
	s.Require().NoError(err)
	s.Require().Equal(http.StatusOK, respCode)

	var alerts []entities.Alert
	err = s.UnmarshalHTTPResponse(respBody, &alerts)
	// Should return 2 alerts about 2 park request
	s.Require().NoError(err)
	s.Require().Len(alerts, 2)
	s.Require().Equal(parkingRequest.UserID, alerts[0].UserID)
	s.Require().Equal(parkingSpaceID, alerts[0].ParkingSpaceID)
	s.Require().Equal(parkingRequest3.UserID, alerts[1].UserID)
	s.Require().Equal(parkingSpaceID, alerts[1].ParkingSpaceID)
}

// Can't test it.
func (s *GetAllAlertsTestSuite) TestGetAllAlerts_HappyPath_Scheduler() {
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
		StartTime: time.Now().Add(5 * time.Second), // 5 seconds in the future
		EndTime:   time.Now().Add(7 * time.Second), // 7 seconds in the future
	}
	parkingRequest := utils.CreateParkingRequest(ctx, driverToken, driver.ID, parkingLot.ID, createParkingRequest, &s.RestClientSuite)
	// Assign that parking request a space we chose above.
	utils.AssignParkingSpace(ctx, parkingSpaceID, parkingRequest.ID, adminToken, &s.RestClientSuite)

	// Create a request in the distant future that won't trigger the alert.
	createParkingRequest2 := &models.CreateParkingRequestRequest{
		DestinationParkingLotID: parkingLot.ID,
		// Times are a little bit in the future, because we don't allow creation of parking request with times in the past.
		StartTime: time.Now().Add(5 * time.Hour), // 5 hours in the future
		EndTime:   time.Now().Add(7 * time.Hour), // 7 hours in the future
	}
	parkingRequest2 := utils.CreateParkingRequest(ctx, driverToken, driver.ID, parkingLot.ID, createParkingRequest2, &s.RestClientSuite)
	// Assign that parking request a space we chose above.
	utils.AssignParkingSpace(ctx, parkingSpaceID, parkingRequest2.ID, adminToken, &s.RestClientSuite)

	// Create another parking request that will trigger the alert
	createParkingRequest3 := &models.CreateParkingRequestRequest{
		DestinationParkingLotID: parkingLot.ID,
		// Times are a little bit in the future, because we don't allow creation of parking request with times in the past.
		StartTime: time.Now().Add(6 * time.Second), // 5 seconds in the future
		EndTime:   time.Now().Add(8 * time.Second), // 7 seconds in the future
	}
	parkingRequest3 := utils.CreateParkingRequest(ctx, driverToken, driver.ID, parkingLot.ID, createParkingRequest3, &s.RestClientSuite)
	// Assign that parking request a space we chose above.
	utils.AssignParkingSpace(ctx, parkingSpaceID, parkingRequest3.ID, adminToken, &s.RestClientSuite)

	// Sleep for 10 secs.
	time.Sleep(10 * time.Second)

	// Don't send the arrival notification.

	// --------
	// ACT
	// --------

	// --------
	// ASSERT
	// --------

}

func TestGetAllAlertsTestSuiteInit(t *testing.T) {
	suite.Run(t, new(GetAllAlertsTestSuite))
}
