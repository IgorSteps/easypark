package models

import (
	"time"
)

// CreatePaymentRequestRequest represent the data in an incoming HTTP request to create a payment request.
type CreatePaymentRequestRequest struct {
	Name           string
	BillingAddress string
	CardNumber     int
	ExpiryDate     time.Time
	CVC            int
}
