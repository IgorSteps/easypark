package functional

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/IgorSteps/easypark/internal/adapters/rest/models"
	"github.com/IgorSteps/easypark/tests/functional/client"
	"github.com/IgorSteps/easypark/tests/functional/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type TestApproveParkingRequestSuite struct {
	client.RestClientSuite
}

func (s *TestApproveParkingRequestSuite) TestUpdateParkingRequestStatus_UnhappyPath_StatusNotAllowed() {
	// --------
	// ASSEMBLE
	// --------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	driver, driverToken := utils.CreateAndLoginDriver(ctx, &s.RestClientSuite, nil)
	adminToken := utils.CreateAndLoginAdmin(ctx, &s.RestClientSuite)

	// Create parking request.
	parkingRequest := utils.CreateParkingRequest(ctx, driverToken, driver.ID, uuid.New(), nil, &s.RestClientSuite)

	updateRequst := &models.UpdateParkingRequestStatusRequest{
		Status: "approved",
	}

	// --------
	// ACT
	// --------
	respBody, respCode, err := s.UpdateParkingRequestStatus(ctx, adminToken, parkingRequest.ID.String(), updateRequst)

	// --------
	// ASSERT
	// --------
	s.Require().NoError(err, "Updating status of the parking request shouldn't return error")
	s.Require().Equal(http.StatusBadRequest, respCode, "Must return 400")

	s.Require().Equal("unknown or not allowed parking request status\n", string(respBody), "Response message is wrong")
}

func (s *TestApproveParkingRequestSuite) TestRejectParkingRequest_HappyPath() {
	// --------
	// ASSEMBLE
	// --------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	driver, driverToken := utils.CreateAndLoginDriver(ctx, &s.RestClientSuite, nil)
	adminToken := utils.CreateAndLoginAdmin(ctx, &s.RestClientSuite)

	// Create parking request.
	parkingRequest := utils.CreateParkingRequest(ctx, driverToken, driver.ID, uuid.New(), nil, &s.RestClientSuite)

	updateRequst := &models.UpdateParkingRequestStatusRequest{
		Status: "rejected",
	}

	// --------
	// ACT
	// --------
	respBody, respCode, err := s.UpdateParkingRequestStatus(ctx, adminToken, parkingRequest.ID.String(), updateRequst)

	// --------
	// ASSERT
	// --------
	s.Require().NoError(err, "Updating status of the parking request shouldn't return error")
	s.Require().Equal(http.StatusOK, respCode, "Must return 200")

	// Unmarshall response body.
	var tModel models.UpdateParkingRequestStatusResponse
	err = s.UnmarshalHTTPResponse(respBody, &tModel)
	if err != nil {
		s.T().Fail()
	}
	s.Require().Equal("successfully updated parking request status", tModel.Message, "Response message is wrong")
}

func (s *TestApproveParkingRequestSuite) TestUpdateParkingRequest_UnhappyPath_UnknownStatus() {
	// --------
	// ASSEMBLE
	// --------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	driver, driverToken := utils.CreateAndLoginDriver(ctx, &s.RestClientSuite, nil)
	adminToken := utils.CreateAndLoginAdmin(ctx, &s.RestClientSuite)

	// Create parking request.
	parkingRequest := utils.CreateParkingRequest(ctx, driverToken, driver.ID, uuid.New(), nil, &s.RestClientSuite)

	updateRequst := &models.UpdateParkingRequestStatusRequest{
		Status: "boom",
	}

	// --------
	// ACT
	// --------
	respBody, respCode, err := s.UpdateParkingRequestStatus(ctx, adminToken, parkingRequest.ID.String(), updateRequst)

	// --------
	// ASSERT
	// --------
	s.Require().NoError(err, "Updating status of the parking request shouldn't return error")
	s.Require().Equal(http.StatusBadRequest, respCode, "Must return 400")
	s.Require().Equal("unknown or not allowed parking request status\n", string(respBody), "Response message is wrong")
}

func TestApproveParkingRequestSuiteInit(t *testing.T) {
	suite.Run(t, new(TestApproveParkingRequestSuite))
}
