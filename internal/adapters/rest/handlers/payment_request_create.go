package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/IgorSteps/easypark/internal/adapters/rest/models"
	"github.com/sirupsen/logrus"
)

// PaymentRequestCreateHandler provides a REST Handler implementation to create payment requests and
// implements http.Handler interface.
type PaymentRequestCreateHandler struct {
	logger *logrus.Logger
}

// NewPaymentRequestCreateHandler creates new instance of PaymentRequestCreateHandler.
func NewPaymentRequestCreateHandler(l *logrus.Logger) *PaymentRequestCreateHandler {
	return &PaymentRequestCreateHandler{
		logger: l,
	}
}

// ServeHTTP handles incoming HTTP request to create payment requests.
func (s *PaymentRequestCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request models.CreatePaymentRequestRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		s.logger.Error("failed to decode payment request creation request: ", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	resp := models.LoginUserResponse{Message: "Payment request sent successfully"}
	json.NewEncoder(w).Encode(resp)
}
