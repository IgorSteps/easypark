package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/IgorSteps/easypark/internal/adapters/rest/models"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/sirupsen/logrus"
)

// CheckLateArrivalHandler provides a REST Handler implementation to check for late arrivals and
// implements http.Handler interface.
type CheckLateArrivalHandler struct {
	logger *logrus.Logger
	facade AlertFacade
}

// NewCheckLateArrivalHandler returns a new instance of CheckLateArrivalHandler.
func NewCheckLateArrivalHandler(l *logrus.Logger, f AlertFacade) *CheckLateArrivalHandler {
	return &CheckLateArrivalHandler{
		logger: l,
		facade: f,
	}
}

// ServeHTTP handles incoming HTTP request to check for late arrivals.
func (s *CheckLateArrivalHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var modelRequest models.CheckForLateArrivalsRequest

	err := json.NewDecoder(r.Body).Decode(&modelRequest)
	if err != nil {
		s.logger.WithError(err).Error("failed to decode request")
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	alerts, err := s.facade.CheckForLateArrivals(r.Context(), modelRequest.Threshold)
	if err != nil {
		s.logger.WithError(err).Error("failed to check for late arrivals")

		switch err.(type) {
		case *repositories.InvalidInputError:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		case *repositories.InternalError:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		default:
			s.logger.WithError(err).Warn("unknown error type")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(alerts)
}
