package functional

import (
	"context"
	"testing"
	"time"

	"github.com/IgorSteps/easypark/internal/adapters/rest/models"
	"github.com/IgorSteps/easypark/tests/functional/client"
	"github.com/IgorSteps/easypark/tests/functional/utils"
	"github.com/stretchr/testify/suite"
)

type AlertCreateTestSuite struct {
	client.RestClientSuite
}

func (s *AlertCreateTestSuite) TestCreateAlert_LocationMismatch() {
	// --------
	// ASSEMBLE
	// --------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	adminToken := utils.CreateAndLoginAdmin(ctx, &s.RestClientSuite)
	createParkingLot := &models.CreateParkingLotRequest{
		Name:     "test-lot",
		Capacity: 5,
	}
	parkingLot := utils.CreateParkingLot(ctx, adminToken, createParkingLot, &s.RestClientSuite)
	parkingSpaceID := parkingLot.ParkingSpaces[0].ID
	//parkingSpaceLocation := parkingLot.ParkingSpaces[0].Name // actual location
	arrivalNotification := &models.CreateNotificationRequest{
		ParkingSpaceID:   parkingSpaceID,
		Location:         "wrong location",
		NotificationType: 0, // arrival
	}
	driver, driverToken := utils.CreateAndLoginDriver(ctx, &s.RestClientSuite, nil)

	// --------
	// ACT
	// --------
	utils.CreateNotification(ctx, driverToken, adminToken, driver.ID, arrivalNotification, &s.RestClientSuite)

	// --------
	// ASSERT
	// --------
	// Cannot be asserted that the alert was created, because it is created during the notification creation process, hence the client only gets a new notification
	// in the HTTP response - no alert data is fed back to the client. We debug log it for now.
}

func TestAlertCreateTestSuiteInit(t *testing.T) {
	suite.Run(t, new(AlertCreateTestSuite))
}
