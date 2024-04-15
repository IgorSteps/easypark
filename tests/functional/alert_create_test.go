package functional

import (
	"context"
	"testing"
	"time"

	"github.com/IgorSteps/easypark/internal/adapters/rest/models"
	"github.com/IgorSteps/easypark/internal/domain/entities"
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
	driver, driverToken := utils.CreateAndLoginDriver(ctx, &s.RestClientSuite, nil)

	// Create a parking lot
	createParkingLot := &models.CreateParkingLotRequest{
		Name:     "test-lot",
		Capacity: 5,
	}
	parkingLot := utils.CreateParkingLot(ctx, adminToken, createParkingLot, &s.RestClientSuite)
	parkingSpaceID := parkingLot.ParkingSpaces[0].ID

	// Create parking request
	parkingRequest := utils.CreateParkingRequest(ctx, driverToken, driver.ID, parkingLot.ID, nil, &s.RestClientSuite)
	// Assign that parking request a space we chose above.
	utils.AssignParkingSpace(ctx, parkingSpaceID, parkingRequest.ID, adminToken, &s.RestClientSuite)

	arrivalNotification := &models.CreateNotificationRequest{
		ParkingRequestID: parkingRequest.ID,
		ParkingSpaceID:   parkingSpaceID,
		Location:         "wrong location",
		NotificationType: 0, // arrival
	}

	// --------
	// ACT
	// --------
	utils.CreateNotification(ctx, driverToken, adminToken, driver.ID, arrivalNotification, &s.RestClientSuite)
	space := utils.GetSingleParkingSpace(ctx, parkingSpaceID.String(), adminToken, &s.RestClientSuite)

	// --------
	// ASSERT
	// --------
	// Cannot be asserted that the alert was created, because it is created during the notification creation process,
	// hence the client only gets a new notification in the HTTP response - no alert data is fed back to the client.
	// We can only assert that parking space status hasn't been changed to 'occupied', because the alert has been sent.
	s.Require().Equal(entities.ParkingSpaceStatusAvailable, space.Status)
}

func TestAlertCreateTestSuiteInit(t *testing.T) {
	suite.Run(t, new(AlertCreateTestSuite))
}
