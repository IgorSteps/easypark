package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// AlertGetSingleHandler provides a REST Handler implementation to get a single alert and
// implements http.Handler interface.
type AlertGetSingleHandler struct {
	logger *logrus.Logger
	facade AlertFacade
}

// NewAlertGetSingleHandler returns a new instance of AlertGetSingleHandler.
func NewAlertGetSingleHandler(l *logrus.Logger, f AlertFacade) *AlertGetSingleHandler {
	return &AlertGetSingleHandler{
		logger: l,
		facade: f,
	}
}

// ServeHTTP handles incoming requests to get a single alert. Method name matches the http.Handler interface.
func (s *AlertGetSingleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	alertID := chi.URLParam(r, "id")

	parsedID, err := uuid.Parse(alertID)
	if err != nil {
		s.logger.WithError(err).Error("failed to parse alert id")
		http.Error(w, "invalid alert id", http.StatusBadRequest)
		return
	}

	alert, err := s.facade.GetAlert(r.Context(), parsedID)
	if err != nil {
		s.logger.WithError(err).Error("failed to get single alert")

		switch err.(type) {
		case *repositories.NotFoundError:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
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
	json.NewEncoder(w).Encode(alert)
}
