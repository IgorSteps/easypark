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

type TestBanDriverSuite struct {
	client.RestClientSuite
}

func (s *TestBanDriverSuite) TestBanDriver_HappyPath() {
	// --------
	// ASSEMBLE
	// --------
	testReq := &models.UpdateStatusRequest{
		Status: "ban",
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	driver, _ := utils.CreateAndLoginDriver(ctx, &s.RestClientSuite, nil)
	adminToken := utils.CreateAndLoginAdmin(ctx, &s.RestClientSuite)

	// --------
	// ACT
	// --------
	respBody, respCode, err := s.BanDriver(ctx, adminToken, driver.ID.String(), testReq)

	// --------
	// ASSERT
	// --------
	s.Require().NoError(err, "Banning a driver shouldn't error out")
	s.Require().Equal(http.StatusOK, respCode, "Response codes don't match")

	var targetModal models.UpdateStatusResponse
	err = s.UnmarshalHTTPResponse(respBody, &targetModal)
	if err != nil {
		s.T().Fail()
	}
	s.Require().Equal("successfully updated user status", targetModal.Message, "Response body content is wrong")
}

func (s *TestBanDriverSuite) TestBannedDriver_CannotAccessAnyRoutes() {
	// --------
	// ASSEMBLE
	// --------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	driver, driverToken := utils.CreateAndLoginDriver(ctx, &s.RestClientSuite, nil)
	adminToken := utils.CreateAndLoginAdmin(ctx, &s.RestClientSuite)

	testBanReq := &models.UpdateStatusRequest{
		Status: "ban",
	}

	// Ban this driver:
	respBody, respCode, err := s.BanDriver(ctx, adminToken, driver.ID.String(), testBanReq)

	// Assert everything is okay.
	s.Require().NoError(err, "Banning a driver shouldn't error out")
	s.Require().Equal(http.StatusOK, respCode, "Response codes don't match")

	// ------
	// ACT
	// ------
	// Try accessing an endpoint.
	respBody, respCode, err = s.PlaceholderDriverRoute(ctx, driverToken)

	// ------
	// ASSERT
	// ------
	s.Require().NoError(err, "Shouldn't error")
	s.Require().Equal(http.StatusForbidden, respCode, "Response code must be FORBIDEN")
	s.Require().Equal("Account is banned.\n", string(respBody), "Response body is wrong")
}

func TestBanDriverSuiteInit(t *testing.T) {
	suite.Run(t, new(TestBanDriverSuite))
}
