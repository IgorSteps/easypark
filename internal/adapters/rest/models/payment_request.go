package models

import (
	"time"
)

// CreatePaymentRequest represent the data in an incoming HTTP request to create a payment request.
type CreatePaymentRequest struct {
	Name           string
	BillingAddress string
	CardNumber     int
	ExpiryDate     time.Time
	CVC            int
}
