package functional_test

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

type TestCreatePaymentRequest struct {
	client.RestClientSuite
}

func (s *TestCreatePaymentRequest) TestCreatePaymentRequest_HappyPath() {
	// --------
	// ASSEMBLE
	// --------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	driver, driverToken := utils.CreateAndLoginDriver(ctx, &s.RestClientSuite, nil)

	testRequest := &models.CreatePaymentRequest{
		Name:           "John Doe",
		BillingAddress: "123 Street Name, City, Postcode",
		CardNumber:     1111222233334444,
		ExpiryDate:     time.Date(2025, 01, 01, 01, 01, 01, 01, time.Local),
		CVC:            123,
	}

	// --------
	// ACT
	// --------
	_, respCode, err := s.CreatePayment(ctx, driverToken, driver.ID.String(), testRequest)

	// --------
	// ASSERT
	// --------
	s.Require().NoError(err, "Creating payment should not return an error")

	s.Require().Equal(http.StatusOK, respCode, "Response code is wrong")
}

func TestCreatePaymentRequestInit(t *testing.T) {
	suite.Run(t, new(TestCreatePaymentRequest))
}
