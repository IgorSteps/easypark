package functional

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/IgorSteps/easypark/internal/adapters/rest/models"
	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/tests/functional/client"
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
	err := PopulateUsers(ctx, &s.RestClientSuite)
	s.Require().NoError(err, "Populating system with mock user data shouldn't return an error")
	token := CreateAdmin(ctx, &s.RestClientSuite)

	// Get all drivers.
	respBody, respCode, err := s.GetAllDrivers(ctx, token)
	s.Require().NoError(err, "Getting all driver's must not return an error")
	s.Require().Equal(http.StatusOK, respCode, "Response codemust be 200")

	// Unmarshall response
	var users []entities.User
	err = s.UnmarshalHTTPResponse(respBody, &users)
	if err != nil {
		s.T().Fail()
	}

	// Extract user's id to ban.
	idToBan := users[0].ID

	// --------
	// ACT
	// --------
	respBody, respCode, err = s.BanDriver(ctx, token, idToBan.String(), testReq)

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

func TestBanDriverSuiteInit(t *testing.T) {
	suite.Run(t, new(TestBanDriverSuite))
}
