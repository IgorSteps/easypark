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

type TestDeleteParkingLot struct {
	client.RestClientSuite
}

func (s *TestDeleteParkingLot) TestDeleteParkingLot_HappyPath() {
	// --------
	// ASSEMBLE
	// --------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	adminToken := utils.CreateAndLoginAdmin(ctx, &s.RestClientSuite)
	req := &models.CreateParkingLotRequest{
		Name:     "main campus",
		Capacity: 10,
	}

	parkLot := utils.CreateParkingLot(ctx, adminToken, req, &s.RestClientSuite)

	// --------
	// ACT
	// --------
	respBody, respCode, err := s.DeleteParkingLot(ctx, adminToken, parkLot.ID)

	// --------
	// ASSERT
	// --------
	s.Require().NoError(err, "Deleting parking lot should not return an error")
	s.Require().Equal(http.StatusOK, respCode, "Respone code should be 200")
	s.Require().Equal("\"successfully deleted parking lot\"\n", string(respBody), "Response body is wrong")
}

func TestDeleteParkingLotInit(t *testing.T) {
	suite.Run(t, new(TestDeleteParkingLot))
}
