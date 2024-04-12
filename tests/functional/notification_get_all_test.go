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

type TestGetAllNotificationSuite struct {
	client.RestClientSuite
}

func (s *TestCreateNotificationSuite) TestGetAllNotifications_HappyPath() {
	// --------
	// ASSEMBLE
	// --------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Created admin.
	adminToken := utils.CreateAndLoginAdmin(ctx, &s.RestClientSuite)

	// Create driver.
	driver, driverToken := utils.CreateAndLoginDriver(ctx, &s.RestClientSuite, nil)

	// Create park lot
	parkingLot := utils.CreateParkingLot(ctx, adminToken, nil, &s.RestClientSuite)

	var expectedNotfs []entities.Notification
	for i := 0; i < 10; i++ {
		parkingSpaceID := parkingLot.ParkingSpaces[i].ID
		testRequest := &models.CreateNotificationRequest{
			ParkingSpaceID:   parkingSpaceID,
			Location:         "cmp",
			NotificationType: 0, // arrival
		}
		notification := utils.CreateNotification(ctx, driverToken, adminToken, driver.ID, testRequest, &s.RestClientSuite)
		expectedNotfs = append(expectedNotfs, notification)
	}

	// --------
	// ACT
	// --------
	respBody, respStatus, err := s.GetAllNotifications(ctx, adminToken)

	// --------
	// ASSERT
	// --------
	s.Require().NoError(err)
	s.Require().Equal(http.StatusOK, respStatus)
	var targetModel []entities.Notification
	err = s.UnmarshalHTTPResponse(respBody, &targetModel)
	s.Require().NoError(err)
	s.Require().NoError(err, "Must not return error")
	s.Require().Equal(expectedNotfs, targetModel)
}

func TestGetAllNotificationSuiteInit(t *testing.T) {
	suite.Run(t, new(TestGetAllNotificationSuite))
}
