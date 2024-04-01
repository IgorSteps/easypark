package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/IgorSteps/easypark/internal/adapters/rest/models"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// DriverStatusHandler provides a REST Handler implementation to change status of driver users and
// implements http.Handler interface.
type DriverStatusHandler struct {
	logger *logrus.Logger
	facade UserFacade
}

// NewDriverStatusHandler creates new instance of DriverStatusHandler.
func NewDriverStatusHandler(f UserFacade, l *logrus.Logger) *DriverStatusHandler {
	return &DriverStatusHandler{
		logger: l,
		facade: f,
	}
}

// ServeHTTP handles incoming HTTP request to change user status. Method name matches the http.Handler interface.
func (s *DriverStatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request models.UpdateStatusRequest
	driverID := chi.URLParam(r, "id")

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		s.logger.WithError(err).Error("failed to decode request")
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	parsedID, err := uuid.Parse(driverID)
	if err != nil {
		s.logger.WithError(err).Error("failed to parse driver id")
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	if request.Status == "ban" {
		err = s.facade.BanDriver(r.Context(), parsedID)
		if err != nil {
			s.logger.WithError(err).Error("failed to ban user")

			switch err.(type) {
			case *repositories.NotFoundError:
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
	} else {
		s.logger.WithField("status", request.Status).Warn("unimplemented status update workflow called")
		http.Error(w, "unimplemented status", http.StatusNotImplemented)
		return
	}

	w.WriteHeader(http.StatusOK)
	resp := models.UpdateStatusResponse{
		Message: "successfully updated user status",
	}
	json.NewEncoder(w).Encode(resp)
}
