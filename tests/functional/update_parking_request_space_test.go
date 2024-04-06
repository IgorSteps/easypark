package functional

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/IgorSteps/easypark/internal/adapters/rest/models"
	"github.com/IgorSteps/easypark/tests/functional/client"
	"github.com/IgorSteps/easypark/tests/functional/utils"
	"github.com/stretchr/testify/suite"
)

type TestUpdateParkingRequestSpaceSuite struct {
	client.RestClientSuite
}

func (s *TestUpdateParkingRequestSpaceSuite) TestUpdateParkingRequestSpace_HappyPath() {
	// --------
	// ASSEMBLE
	// --------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	driver, driverToken := utils.CreateAndLoginDriver(ctx, &s.RestClientSuite, nil)
	adminToken := utils.CreateAndLoginAdmin(ctx, &s.RestClientSuite)

	// Creating a parking lot.
	parkingLot := utils.CreateParkingLot(ctx, adminToken, &s.RestClientSuite)

	// Creating a parking request.
	parkingRequest := utils.CreateParkingRequest(ctx, driverToken, driver.ID, parkingLot.ID, &s.RestClientSuite)

	testUpdateRequestSpace := &models.ParkingRequestSpaceUpdateRequest{
		ParkingSpaceID: parkingLot.ParkingSpaces[0].ID,
	}

	// --------
	// ACT
	// --------
	respBody, respCode, err := s.UpdateParkingRequestSpace(ctx, adminToken, parkingRequest.ID, testUpdateRequestSpace)
	s.Require().NoError(err, "Updating a parking request shouldn't return an error")

	// --------
	// ASSERT
	// --------
	s.Require().Equal(http.StatusOK, respCode, "Updating parking request should return 200 code")

	var updateParkingRequestResp models.UpdateParkingRequestStatusResponse
	err = s.UnmarshalHTTPResponse(respBody, &updateParkingRequestResp)
	s.Require().NoError(err, "Must not return error")
	s.Require().Equal("successfully assigned a space to a parking request", updateParkingRequestResp.Message, "Response body is wrong.")
}

func TestUpdateParkingRequestSpaceSuiteInit(t *testing.T) {
	suite.Run(t, new(TestUpdateParkingRequestSpaceSuite))
}
