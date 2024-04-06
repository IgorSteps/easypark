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

type TestGetDriversParkingRequestsSuite struct {
	client.RestClientSuite
}

func (s *TestGetDriversParkingRequestsSuite) TestGetDriversParkingRequests_HappyPath() {
	// --------
	// ASSEMBLE
	// --------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	driver, driverToken := utils.CreateAndLoginDriver(ctx, &s.RestClientSuite, nil)

	anotherDriver := &models.UserCreationRequest{
		Username:  "test",
		Password:  "test",
		Email:     "test@example.com",
		Firstname: "test",
		Lastname:  "test",
	}
	driver2, driver2Token := utils.CreateAndLoginDriver(ctx, &s.RestClientSuite, anotherDriver)

	for i := 0; i < 5; i++ {
		_ = utils.CreateParkingRequest(ctx, driverToken, driver.ID, uuid.New(), &s.RestClientSuite)
	}

	for i := 0; i < 5; i++ {
		_ = utils.CreateParkingRequest(ctx, driver2Token, driver2.ID, uuid.New(), &s.RestClientSuite)
	}

	// --------
	// ACT
	// --------
	respBody, respCode, err := s.GetDriversParkingRequests(ctx, driverToken, driver.ID)

	// --------
	// ASSERT
	// --------
	s.Require().NoError(err, "Getting driver's parking requests shouldn't return an error")
	s.Require().Equal(http.StatusOK, respCode, "Response code must be 200")

	var targetModel []entities.ParkingRequest
	err = s.UnmarshalHTTPResponse(respBody, &targetModel)
	s.Require().NoError(err, "Unmarshalling response body must not error")

	s.Require().Equal(5, len(targetModel), "Must have 5 parking requests")
}

func TestGetDriversParkingRequestsSuiteInit(t *testing.T) {
	suite.Run(t, new(TestGetDriversParkingRequestsSuite))
}
