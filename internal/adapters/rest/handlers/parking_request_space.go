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

// ParkingRequestSpaceHandler provides a REST Handler implementation to update parking spaces in parking requests and
// implements http.Handler interface.
type ParkingRequestSpaceHandler struct {
	logger *logrus.Logger
	facade ParkingRequestFacade
}

// NewParkingRequestSpaceHandler creates new instance of ParkingRequestSpaceHandler.
func NewParkingRequestSpaceHandler(f ParkingRequestFacade, l *logrus.Logger) *ParkingRequestSpaceHandler {
	return &ParkingRequestSpaceHandler{
		logger: l,
		facade: f,
	}
}

// ServeHTTP handles incoming HTTP request to assign/deassign a parking space to a parking request.
// Method name matches the http.Handler interface.
func (s *ParkingRequestSpaceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var updateParkingRequest models.ParkingRequestSpaceUpdateRequest

	parkingRequestID := chi.URLParam(r, "id")

	err := json.NewDecoder(r.Body).Decode(&updateParkingRequest)
	if err != nil {
		s.logger.WithError(err).Error("failed to decode request")
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	parsedID, err := uuid.Parse(parkingRequestID)
	if err != nil {
		s.logger.WithError(err).Error("failed to parse parking request id")
		http.Error(w, "invalid parking request id", http.StatusBadRequest)
		return
	}

	// Assign a parking spacec
	if updateParkingRequest.ParkingSpaceID != uuid.Nil {
		err = s.facade.AssignParkingSpace(r.Context(), parsedID, updateParkingRequest.ParkingSpaceID)

		if err != nil {
			s.logger.WithError(err).Error("failed to assign parking request a space")

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
		response := models.ParkingRequestSpaceUpdateResponse{
			Message: "successfully assigned a space to a parking request",
		}
		json.NewEncoder(w).Encode(response)
	} else if updateParkingRequest.ParkingSpaceID == uuid.Nil {
		err := s.facade.DeassignParkingSpace(r.Context(), parsedID)
		if err != nil {
			s.logger.WithError(err).Error("failed to deassign parking request from a space")

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
	} else {
		s.logger.WithField("space id", updateParkingRequest.ParkingSpaceID).Warn("unknown parking space ID state")
		w.WriteHeader(http.StatusOK)
		return
	}
}
