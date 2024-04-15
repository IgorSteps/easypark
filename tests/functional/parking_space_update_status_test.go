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
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type TestUpdateParkingSpaceStatusSuite struct {
	client.RestClientSuite
}

func (s *TestUpdateParkingSpaceStatusSuite) TestUpdateParkingSpaceStatusSuite_HappyPath() {
	// ---------
	// ASSEMBLE
	// ---------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	adminToken := utils.CreateAndLoginAdmin(ctx, &s.RestClientSuite)
	parkingLot := utils.CreateParkingLot(ctx, adminToken, nil, &s.RestClientSuite)
	parkingSpaceID := parkingLot.ParkingSpaces[0].ID

	tc := []struct {
		name           string
		status         string
		expectedStatus entities.ParkingSpaceStatus
	}{
		{
			name:           "Updating status to Blocked",
			status:         "blocked",
			expectedStatus: entities.ParkingSpaceStatusBlocked,
		},
		{
			name:           "Updating status to Reserved",
			status:         "reserved",
			expectedStatus: entities.ParkingSpaceStatusReserved,
		},
		{
			name:           "Updating status to Occupied",
			status:         "occupied",
			expectedStatus: entities.ParkingSpaceStatusOccupied,
		},
		{
			name:           "Updating status to Available",
			status:         "available",
			expectedStatus: entities.ParkingSpaceStatusAvailable,
		},
	}

	for _, test := range tc {
		s.T().Run(test.name, func(t *testing.T) {
			updateRequest := &models.UpdateParkingSpaceStatus{
				Status: test.status,
			}

			// ---------
			// ACT
			// ---------
			respBody, respCode, err := s.UpdateParkingSpaceStatus(ctx, adminToken, parkingSpaceID, updateRequest)

			// ---------
			// ASSERT
			// ---------
			s.Require().NoError(err, "No error must be returned when updating parking space status")
			s.Require().Equal(http.StatusOK, respCode, "Response code should be 200")

			// Unmarshall response.
			var targetSpaceModel entities.ParkingSpace
			err = s.UnmarshalHTTPResponse(respBody, &targetSpaceModel)
			s.Require().NoError(err, "failed to unmarshall response")

			s.Require().Equal(test.expectedStatus, targetSpaceModel.Status, "Wrong status")
		})
	}
}

func (s *TestUpdateParkingSpaceStatusSuite) TestUpdateParkingSpaceStatusSuite_UnhappyPath_InvalidStatus() {
	// ---------
	// ASSEMBLE
	// ---------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	adminToken := utils.CreateAndLoginAdmin(ctx, &s.RestClientSuite)

	updateRequest := &models.UpdateParkingSpaceStatus{
		Status: "boom",
	}

	// ---------
	// ACT
	// ---------
	respBody, respCode, err := s.UpdateParkingSpaceStatus(ctx, adminToken, uuid.New(), updateRequest)

	// ---------
	// ASSERT
	// ---------
	s.Require().NoError(err, "No error must be returned when updating parking space status")
	s.Require().Equal(http.StatusBadRequest, respCode, "Response code should be 400")
	s.Require().Equal("failed to parse given status\n", string(respBody))
}

func TestUpdateParkingSpaceStatusSuiteInit(t *testing.T) {
	suite.Run(t, new(TestUpdateParkingSpaceStatusSuite))
}
