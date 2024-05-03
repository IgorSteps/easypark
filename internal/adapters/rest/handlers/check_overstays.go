package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/IgorSteps/easypark/internal/adapters/rest/models"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/sirupsen/logrus"
)

// CheckOverStaysHandler provides a REST Handler implementation to check for over staying users and
// implements http.Handler interface.
type CheckOverStaysHandler struct {
	logger *logrus.Logger
	facade AlertFacade
}

// NewCheckOverStaysHandler returns a new instance of CheckOverStaysHandler.
func NewCheckOverStaysHandler(l *logrus.Logger, f AlertFacade) *CheckOverStaysHandler {
	return &CheckOverStaysHandler{
		logger: l,
		facade: f,
	}
}

// ServeHTTP handles incoming HTTP request to check for over staying users.
func (s *CheckOverStaysHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var modelRequest models.CheckForOverStaysRequest

	err := json.NewDecoder(r.Body).Decode(&modelRequest)
	if err != nil {
		s.logger.WithError(err).Error("failed to decode request")
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	alerts, err := s.facade.CheckForOverStays(r.Context(), modelRequest.Threshold)
	if err != nil {
		s.logger.WithError(err).Error("failed to check for overdue stays")

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
