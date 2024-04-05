package functional

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/IgorSteps/easypark/internal/adapters/rest/models"
	"github.com/IgorSteps/easypark/tests/functional/client"
	"github.com/stretchr/testify/suite"
)

type TestCreateParkingLot struct {
	client.RestClientSuite
}

func (s *TestCreateParkingLot) TestCreateParkingLot_HappyPath() {
	// --------
	// ASSEMBLE
	// --------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Created admin
	adminToken := CreateAdmin(ctx, &s.RestClientSuite)
	testRequest := &models.CreateParkingLotRequest{
		Name:     "science",
		Capacity: 10,
	}

	// --------
	// ACT
	// --------
	respBody, respCode, err := s.CreateParkingLot(ctx, adminToken, testRequest)

	// --------
	// ASSERT
	// --------
	s.Require().NoError(err, "Creating parking lot should not return an error")

	var targetModel models.CreateParkingLotResponse
	err = s.UnmarshalHTTPResponse(respBody, &targetModel)
	if err != nil {
		s.T().Fail()
	}

	s.Require().Equal(http.StatusCreated, respCode, "Response code must be 201")
	s.Require().NotNil(targetModel.ID, "Created parking lot should have ID set")
	s.Require().Equal("science", targetModel.Name, "Parking lot name is wrong")
	s.Require().Equal(10, targetModel.Capacity, "Parking lot capacity is wrong")
	s.Require().NotEmpty(targetModel.PakringSpaces, "Parking spaces slice cannot be empty")
}

func (s *TestCreateParkingLot) TestCreateParkingLot_UnhappyPath_AlreadyExists() {
	// --------
	// ASSEMBLE
	// --------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Created admin
	adminToken := CreateAdmin(ctx, &s.RestClientSuite)
	testRequest := &models.CreateParkingLotRequest{
		Name:     "science",
		Capacity: 10,
	}
	// Create a parking lot with name "science"
	_, _, err := s.CreateParkingLot(ctx, adminToken, testRequest)
	s.Require().NoError(err, "Creating parking lot should not return an error")

	testRequest2 := &models.CreateParkingLotRequest{
		Name:     "science",
		Capacity: 10,
	}

	// --------
	// ACT
	// --------
	// Create a parking lot with the same name
	respBody, respCode, err := s.CreateParkingLot(ctx, adminToken, testRequest2)

	// --------
	// ASSERT
	// --------
	s.Require().NoError(err, "Creating parking lot should not return an error")
	s.Require().Equal(http.StatusBadRequest, respCode, "Response code must be 400")
	s.Require().Equal("Resource 'science' already exists\n", string(respBody), "Resonse body is wrong")
}

func TestCreateParkingLotInit(t *testing.T) {
	suite.Run(t, new(TestCreateParkingLot))
}
