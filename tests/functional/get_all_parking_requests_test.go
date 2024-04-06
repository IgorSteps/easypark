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

type TestGetAllParkingRequestsSuite struct {
	client.RestClientSuite
}

func (s *TestGetAllParkingRequestsSuite) TestGetAllParkingRequests_HappyPath() {
	// --------
	// ASSEMBLE
	// --------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	driver, driverToken := utils.CreateAndLoginDriver(ctx, &s.RestClientSuite, nil)
	adminToken := utils.CreateAndLoginAdmin(ctx, &s.RestClientSuite)

	// Create 10 parking requests:
	for i := 0; i < 10; i++ {
		testRequest := &models.CreateParkingRequestRequest{
			DestinationParkingLotID: uuid.New(),
			StartTime:               time.Now(),
			EndTime:                 time.Now().Add(555),
		}
		_, respCode, err := s.CreateParkingRequest(ctx, driverToken, driver.ID.String(), testRequest)
		s.Require().NoError(err, "Creating new parking request must not return an error")
		s.Require().Equal(http.StatusCreated, respCode, "Response code must be 201 CREATED")
	}

	// --------
	// ACT
	// --------
	respBody, respCode, err := s.GetAllParkingRequests(ctx, adminToken)

	// --------
	// ASSERT
	// --------
	s.Require().NoError(err, "Getting all parking requests shouldn't return an error")
	s.Require().Equal(http.StatusOK, respCode, "Response code must be 200")

	var targetModel []entities.ParkingRequest
	err = s.UnmarshalHTTPResponse(respBody, &targetModel)
	s.Require().NoError(err, "Unmarshalling response body must not error")

	s.Require().Equal(10, len(targetModel), "Must have 10 parking requests")
}

func TestGetAllParkingRequestsSuiteInit(t *testing.T) {
	suite.Run(t, new(TestGetAllParkingRequestsSuite))
}
