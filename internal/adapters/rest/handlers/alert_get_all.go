package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/sirupsen/logrus"
)

// AlertGetAllHandler provides a REST Handler implementation to get all alerts and
// implements http.Handler interface.
type AlertGetAllHandler struct {
	logger *logrus.Logger
	facade AlertFacade
}

// NewAlertGetAllHandler returns a new instance of AlertGetAllHandler.
func NewAlertGetAllHandler(l *logrus.Logger, f AlertFacade) *AlertGetAllHandler {
	return &AlertGetAllHandler{
		logger: l,
		facade: f,
	}
}

// ServeHTTP handles incoming requests to get all alerts. Method name matches the http.Handler interface.
func (s *AlertGetAllHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	alerts, err := s.facade.GetAllAlerts(r.Context())
	if err != nil {
		s.logger.WithError(err).Error("failed to get all alerts")

		switch err.(type) {
		case *repositories.InternalError:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		default:
			s.logger.Warn("unknown error type, returning internal server error")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(alerts)
}
