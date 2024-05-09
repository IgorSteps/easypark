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

// ParkingRequestCreateHandler provides a REST Handler implementation to create parking requests and
// implements http.Handler interface.
type ParkingRequestCreateHandler struct {
	logger *logrus.Logger
	facade ParkingRequestFacade
}

// NewParkingRequestCreateHandler creates new instance of ParkingRequestCreateHandler.
func NewParkingRequestCreateHandler(f ParkingRequestFacade, l *logrus.Logger) *ParkingRequestCreateHandler {
	return &ParkingRequestCreateHandler{
		logger: l,
		facade: f,
	}
}

// ServeHTTP handles incoming HTTP request to create parking requests.
func (s *ParkingRequestCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request models.CreateParkingRequestRequest
	driverID := chi.URLParam(r, "id")

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		s.logger.Error("failed to decode parking request creation request: ", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	parsedID, err := uuid.Parse(driverID)
	if err != nil {
		s.logger.WithError(err).Error("failed to parse driver id")
		http.Error(w, "invalid driver id", http.StatusBadRequest)
		return
	}

	domainParkingRequest := request.ToDomain()
	// TODO: Move to the usecase?
	domainParkingRequest.UserID = parsedID
	createdParkingRequest, err := s.facade.CreateParkingRequest(r.Context(), domainParkingRequest)
	if err != nil {
		s.logger.WithError(err).Error("failed to create a parking request")

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

	w.WriteHeader(http.StatusCreated)
	response := models.CreateParkingRequestResponse{
		ID:                 createdParkingRequest.ID,
		UserID:             createdParkingRequest.UserID,
		DestinationLotID:   createdParkingRequest.DestinationParkingLotID,
		DestinationLotName: createdParkingRequest.DestinationParkingLotName,
		StartTime:          createdParkingRequest.StartTime,
		EndTime:            createdParkingRequest.EndTime,
		Status:             createdParkingRequest.Status,
		Cost:               createdParkingRequest.Cost,
		CreatedAt:          createdParkingRequest.CreatedAt,
		UpdatedAt:          createdParkingRequest.UpdatedAt,
	}
	json.NewEncoder(w).Encode(response)
}
