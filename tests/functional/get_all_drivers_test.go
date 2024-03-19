package functional

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/IgorSteps/easypark/tests/functional/client"
	"github.com/stretchr/testify/suite"
)

type TestGetAllDriversSuite struct {
	client.RestClientSuite
}

func (s *TestGetAllDriversSuite) TestGetAllDrivers_HappyPath() {
	// --------
	// ASSEMBLE
	// --------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := PopulateUsers(ctx, &s.RestClientSuite)
	s.Require().NoError(err, "Populating system with mock user data shouldn't return an error")
	token := CreateAdmin(ctx, &s.RestClientSuite)

	// --------
	// ACT
	// --------
	respBody, respCode, err := s.GetAllDrivers(ctx, token)

	// --------
	// ASSERT
	// --------
	s.Require().NoError(err, "Getting all drivers shouldn't return an error")
	s.Require().Equal(http.StatusOK, respCode, "Response codes don't match")
	s.Require().NotEmpty(respBody, "Response body can't be empty")
}

func TestGetAllDriversInit(t *testing.T) {
	suite.Run(t, new(TestGetAllDriversSuite))
}
