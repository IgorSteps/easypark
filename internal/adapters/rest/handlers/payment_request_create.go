package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/IgorSteps/easypark/internal/adapters/rest/models"
	"github.com/sirupsen/logrus"
)

// PaymentCreateHandler provides a REST Handler implementation to create payments and
// implements http.Handler interface.
type PaymentCreateHandler struct {
	logger *logrus.Logger
}

// NewPaymentCreateHandler creates new instance of PaymentCreateHandler.
func NewPaymentCreateHandler(l *logrus.Logger) *PaymentCreateHandler {
	return &PaymentCreateHandler{
		logger: l,
	}
}

// ServeHTTP handles incoming HTTP request to create payments.
func (s *PaymentCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request models.CreatePaymentRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		s.logger.Error("failed to decode payment creation request: ", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	resp := models.PaymentResponse{Message: "Payment sent successfully"}
	json.NewEncoder(w).Encode(resp)
}
