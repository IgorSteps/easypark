package usecases

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// UpdateParkingSpaceStatus provides business logic to update a parking space status.
type UpdateParkingSpaceStatus struct {
	logger      *logrus.Logger
	spaceRepo   repositories.ParkingSpaceRepository
	requestRepo repositories.ParkingRequestRepository
}

// NewUpdateParkingSpaceStatus returns a new instance of the UpdateParkingSpaceStatus.
func NewUpdateParkingSpaceStatus(
	l *logrus.Logger,
	spaceRepo repositories.ParkingSpaceRepository,
	reqRepo repositories.ParkingRequestRepository,
) *UpdateParkingSpaceStatus {
	return &UpdateParkingSpaceStatus{
		logger:      l,
		spaceRepo:   spaceRepo,
		requestRepo: reqRepo,
	}
}

// Execute runs the business logic.
func (s *UpdateParkingSpaceStatus) Execute(ctx context.Context, id uuid.UUID, status string) (entities.ParkingSpace, error) {
	domainStatus, err := parseSpaceStatus(status)
	if err != nil {
		s.logger.WithField("status", status).WithError(err).Error("failed to parse given status")
		return entities.ParkingSpace{}, err
	}

	parkSpace, err := s.spaceRepo.GetSingle(ctx, id)
	if err != nil {
		return entities.ParkingSpace{}, err
	}

	// If we are blocking/reserving parking spaces, we must check if they have been assigned to a parking request.
	if domainStatus == entities.ParkingSpaceStatusBlocked || domainStatus == entities.ParkingSpaceStatusReserved {
		// De-assign from this parking space.
		if parkSpace.ParkingRequests != nil {
			var ids []uuid.UUID
			for _, req := range parkSpace.ParkingRequests {
				ids = append(ids, req.ID)
			}

			// Create the query map with ids
			query := map[string]interface{}{
				"id": ids,
			}
			parkingRequests, err := s.requestRepo.GetMany(ctx, query)
			if err != nil {
				return entities.ParkingSpace{}, err
			}

			for _, req := range parkingRequests {
				req.OnSpaceDeassign()
				s.requestRepo.Save(ctx, &req)
			}

			// Remove reference to parking requests.
			parkSpace.ParkingRequests = nil
		}
	}

	// Update status.
	parkSpace.Status = domainStatus

	err = s.spaceRepo.Save(ctx, &parkSpace)
	if err != nil {
		return entities.ParkingSpace{}, err
	}

	return parkSpace, nil
}

func parseSpaceStatus(status string) (entities.ParkingSpaceStatus, error) {
	switch status {
	case "available":
		return entities.ParkingSpaceStatusAvailable, nil
	case "occupied":
		return entities.ParkingSpaceStatusOccupied, nil
	case "blocked":
		return entities.ParkingSpaceStatusBlocked, nil
	case "reserved":
		return entities.ParkingSpaceStatusReserved, nil
	default:
		return "", repositories.NewInvalidInputError("failed to parse given status")
	}
}
