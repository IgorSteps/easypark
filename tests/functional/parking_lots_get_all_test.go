package functional

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/IgorSteps/easypark/internal/adapters/rest/models"
	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/tests/functional/client"
	"github.com/IgorSteps/easypark/tests/functional/utils"
	"github.com/stretchr/testify/suite"
)

type TestGetAllParkingLotsSuite struct {
	client.RestClientSuite
}

func (s *TestGetAllParkingLotsSuite) TestGetAllParkingLots_HappyPath() {
	// --------
	// ASSEMBLE
	// --------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	adminToken := utils.CreateAndLoginAdmin(ctx, &s.RestClientSuite)
	var expectedLots []entities.ParkingLot

	// Create 5 lots:
	for i := 0; i < 5; i++ {
		request := &models.CreateParkingLotRequest{
			Name:     fmt.Sprintf("Lot-%d", i),
			Capacity: 5,
		}
		lot := utils.CreateParkingLot(ctx, adminToken, request, &s.RestClientSuite)
		expectedLots = append(expectedLots, lot)
	}

	// --------
	// ACT
	// --------
	respBody, respCode, err := s.GetAllParkingLots(ctx, adminToken)

	// --------
	// ASSERT
	// --------
	s.Require().NoError(err, "Getting all parking requests shouldn't return an error")
	s.Require().Equal(http.StatusOK, respCode, "Response code must be 200")

	// Unmarshal response.
	var targetModel []entities.ParkingLot
	err = s.UnmarshalHTTPResponse(respBody, &targetModel)
	s.Require().NoError(err, "Unmarshalling response body must not error")
	for i, lot := range expectedLots {
		s.Require().Equal(lot.ID, targetModel[i].ID, "IDs must be equal")
		s.Require().Equal(lot.Name, targetModel[i].Name, "Names must be equal")
		s.Require().Equal(lot.Capacity, targetModel[i].Capacity, "Capacities must be equal")
		s.Require().NotEmpty(targetModel[i].ParkingSpaces, "Parking spaces must be populated")
		s.Require().Equal(lot.ID, targetModel[i].ParkingSpaces[i].ParkingLotID, "Parking space must have the parent parking lot ID set")
	}
}

func TestGetAllParkingLotsSuiteInit(t *testing.T) {
	suite.Run(t, new(TestGetAllParkingLotsSuite))
}
