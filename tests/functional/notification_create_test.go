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

type TestCreateNotificationSuite struct {
	client.RestClientSuite
}

func (s *TestCreateNotificationSuite) TestCreateNotification_HappyPath_Arrival() {
	// --------
	// ASSEMBLE
	// --------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Created admin and driver.
	adminToken := utils.CreateAndLoginAdmin(ctx, &s.RestClientSuite)
	driver, driverToken := utils.CreateAndLoginDriver(ctx, &s.RestClientSuite, nil)

	// Create parking lot.
	parkingLot := utils.CreateParkingLot(ctx, adminToken, nil, &s.RestClientSuite)
	parkingSpaceID := parkingLot.ParkingSpaces[0].ID
	parkingSpaceName := parkingLot.ParkingSpaces[0].Name

	// Create parking request
	parkingRequest := utils.CreateParkingRequest(ctx, driverToken, driver.ID, parkingLot.ID, nil, &s.RestClientSuite)
	// Assign that parking request a space we chose above.
	utils.AssignParkingSpace(ctx, parkingSpaceID, parkingRequest.ID, adminToken, &s.RestClientSuite)

	// Setup the request to create an arrival notification.
	testRequest := &models.CreateNotificationRequest{
		ParkingRequestID: parkingRequest.ID,
		ParkingSpaceID:   parkingSpaceID,
		Location:         parkingSpaceName,
		NotificationType: 0, // arrival
	}

	// --------
	// ACT
	// --------
	respBody, respCode, err := s.CreateNotification(ctx, driverToken, driver.ID, testRequest)
	// Get updated parking space to check the status has been updated.
	parkingSpace := utils.GetSingleParkingSpace(ctx, parkingSpaceID.String(), adminToken, &s.RestClientSuite)

	// --------
	// ASSERT
	// --------
	s.Require().NoError(err, "Must not return an error")
	s.Require().Equal(http.StatusCreated, respCode, "Response code must be 201")

	var targetModel entities.Notification
	err = s.UnmarshalHTTPResponse(respBody, &targetModel)
	s.Require().NoError(err, "Must not return an error")

	// Check statuses.
	s.Require().Equal(entities.ParkingSpaceStatusOccupied, parkingSpace.Status)
	for _, parkReq := range parkingSpace.ParkingRequests {
		// check the status of the parking request was changed as well
		if parkReq.ID == parkingRequest.ID {
			s.Require().Equal(entities.RequestStatusActive, parkReq.Status)
		}
	}
	s.Require().NotEmpty(targetModel.ID)
	s.Require().Equal(driver.ID, targetModel.DriverID)
	s.Require().Equal(testRequest.ParkingSpaceID, targetModel.ParkingSpaceID)
	s.Require().Equal(testRequest.Location, targetModel.Location)
	s.Require().Equal(testRequest.NotificationType, int(targetModel.Type))
}

func (s *TestCreateNotificationSuite) TestCreateNotification_HappyPath_Departure() {
	// --------
	// ASSEMBLE
	// --------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Created admin and driver.
	adminToken := utils.CreateAndLoginAdmin(ctx, &s.RestClientSuite)
	driver, driverToken := utils.CreateAndLoginDriver(ctx, &s.RestClientSuite, nil)

	// Create parking lot.
	parkingLot := utils.CreateParkingLot(ctx, adminToken, nil, &s.RestClientSuite)
	// Lets choose the first parking space in that lot.
	parkingSpaceID := parkingLot.ParkingSpaces[0].ID

	// Create parking request
	parkingRequest := utils.CreateParkingRequest(ctx, driverToken, driver.ID, parkingLot.ID, nil, &s.RestClientSuite)
	// Assign that parking request a space we chose above.
	utils.AssignParkingSpace(ctx, parkingSpaceID, parkingRequest.ID, adminToken, &s.RestClientSuite)

	// Setup the request to create an arrival notification.
	testArrivalRequest := &models.CreateNotificationRequest{
		ParkingRequestID: parkingRequest.ID,
		ParkingSpaceID:   parkingSpaceID,
		Location:         parkingLot.ParkingSpaces[0].Name,
		NotificationType: 0, // departure
	}
	// First a space must get an arrival notification in order to be able to receive departure notifs.
	respBody, respCode, err := s.CreateNotification(ctx, driverToken, driver.ID, testArrivalRequest)
	s.Require().NoError(err, "Must not return an error")
	s.Require().Equal(http.StatusCreated, respCode, "Response code must be 201")

	// Setup the request to create a departure notification.
	testDepartureRequest := &models.CreateNotificationRequest{
		ParkingRequestID: parkingRequest.ID,
		ParkingSpaceID:   parkingSpaceID,
		Location:         parkingLot.ParkingSpaces[0].Name,
		NotificationType: 1, // departure
	}

	// --------
	// ACT
	// --------
	respBody, respCode, err = s.CreateNotification(ctx, driverToken, driver.ID, testDepartureRequest)
	// Get updated parking space to check the status has been updated.
	parkingSpace := utils.GetSingleParkingSpace(ctx, parkingSpaceID.String(), adminToken, &s.RestClientSuite)

	// --------
	// ASSERT
	// --------
	s.Require().NoError(err, "Must not return an error")
	s.Require().Equal(http.StatusCreated, respCode, "Response code must be 201")

	var targetModel entities.Notification
	err = s.UnmarshalHTTPResponse(respBody, &targetModel)
	s.Require().NoError(err, "Must not return an error")

	// Check statuses.
	s.Require().Equal(entities.ParkingSpaceStatusAvailable, parkingSpace.Status)
	for _, parkReq := range parkingSpace.ParkingRequests {
		// check the status of the parking request was changed as well.
		if parkReq.ID == parkingRequest.ID {
			s.Require().Equal(entities.RequestStatusCompleted, parkReq.Status)
		}
	}

	s.Require().NotEmpty(targetModel.ID)
	s.Require().Equal(driver.ID, targetModel.DriverID)
	s.Require().Equal(testDepartureRequest.ParkingSpaceID, targetModel.ParkingSpaceID)
	s.Require().Equal(testDepartureRequest.Location, targetModel.Location)
	s.Require().Equal(testDepartureRequest.NotificationType, int(targetModel.Type))
}

func TestCreateNotificationSuiteInit(t *testing.T) {
	suite.Run(t, new(TestCreateNotificationSuite))
}
