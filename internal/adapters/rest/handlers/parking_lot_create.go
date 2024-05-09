package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/IgorSteps/easypark/internal/adapters/rest/models"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/sirupsen/logrus"
)

type ParkingLotCreateHandler struct {
	logger *logrus.Logger
	facade ParkingLotFacade
}

func NewParkingLotCreateHandler(l *logrus.Logger, f ParkingLotFacade) *ParkingLotCreateHandler {
	return &ParkingLotCreateHandler{
		logger: l,
		facade: f,
	}
}

// ServeHTTP handles incoming HTTP request to create a parking lot.
func (s *ParkingLotCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request models.CreateParkingLotRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		s.logger.Error("failed to decode parking lot creation request: ", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	lot, err := s.facade.CreateParkingLot(r.Context(), request.Name, request.Capacity)
	if err != nil {
		s.logger.WithError(err).Error("failed to create parking lot")

		switch err.(type) {
		case *repositories.ResourceAlreadyExistsError:
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
	resp := models.CreateParkingLotResponse{
		ID:            lot.ID,
		Name:          lot.Name,
		Capacity:      lot.Capacity,
		PakringSpaces: lot.ParkingSpaces,
		CreatedAt:     lot.CreatedAt,
		UpdatedAt:     lot.UpdatedAt,
	}
	json.NewEncoder(w).Encode(resp)
}
