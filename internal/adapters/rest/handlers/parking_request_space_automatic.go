package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/IgorSteps/easypark/internal/adapters/rest/models"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/sirupsen/logrus"
)

// ParkingRequestAutomaticSpaceHandler provides a REST Handler implementation to automatically assign parking spaces
// to parking requests and implements http.Handler interface.
type ParkingRequestAutomaticSpaceHandler struct {
	logger *logrus.Logger
	facade ParkingRequestFacade
}

// NewParkingRequestAutomaticSpaceHandler creates new instance of ParkingRequestAutomaticSpaceHandler.
func NewParkingRequestAutomaticSpaceHandler(f ParkingRequestFacade, l *logrus.Logger) *ParkingRequestAutomaticSpaceHandler {
	return &ParkingRequestAutomaticSpaceHandler{
		logger: l,
		facade: f,
	}
}

func (s *ParkingRequestAutomaticSpaceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var requestModel models.ParkingRequestAutomaticSpaceUpdateRequest

	err := json.NewDecoder(r.Body).Decode(&requestModel)
	if err != nil {
		s.logger.WithError(err).Error("failed to decode request")
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	space, err := s.facade.AutomaticallyAssignParkingSpace(r.Context(), requestModel.ParkingRequestID)
	if err != nil {
		s.logger.WithError(err).Error("failed to update parking request")

		switch err.(type) {
		case *repositories.InvalidInputError:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
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
	json.NewEncoder(w).Encode(space)
}
