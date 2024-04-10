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

	// Created admin.
	adminToken := utils.CreateAndLoginAdmin(ctx, &s.RestClientSuite)

	// Create parking lot.
	parkingLot := utils.CreateParkingLot(ctx, adminToken, nil, &s.RestClientSuite)

	testRequest := &models.CreateNotificationRequest{
		ParkingSpaceID:   parkingLot.ParkingSpaces[0].ID,
		Location:         "cmp",
		NotificationType: 0,
	}

	driver, driverToken := utils.CreateAndLoginDriver(ctx, &s.RestClientSuite, nil)

	// --------
	// ACT
	// --------
	respBody, respCode, err := s.CreateNotification(ctx, driverToken, driver.ID, testRequest)

	// --------
	// ASSERT
	// --------
	s.Require().NoError(err, "Must not return an error")
	s.Require().Equal(http.StatusCreated, respCode, "Response code must be 201")

	var targetModel entities.Notification
	err = s.UnmarshalHTTPResponse(respBody, &targetModel)
	s.Require().NoError(err, "Must not return an error")

	s.Require().NotEmpty(targetModel.ID)
	s.Require().Equal(driver.ID, targetModel.DriverID)
	s.Require().Equal(testRequest.ParkingSpaceID, targetModel.ParkingSpaceID)
	s.Require().Equal(testRequest.Location, targetModel.Location)
	s.Require().Equal(testRequest.NotificationType, int(targetModel.Type))
}

func TestCreateNotificationSuiteInit(t *testing.T) {
	suite.Run(t, new(TestCreateNotificationSuite))
}
