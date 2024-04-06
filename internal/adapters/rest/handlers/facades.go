package handlers

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/google/uuid"
)

// UserFacade is provides an interface implemented by usecasefacades.UserFacade.
type UserFacade interface {
	// CreateDriver is implemented by usecasefacades.UserFacade that wraps driver user creation usecase.
	CreateDriver(ctx context.Context, driver *entities.User) error

	// AuthoriseUser is implemented by usecasefacades.UserFacade that wraps user login usecase.
	AuthoriseUser(ctx context.Context, username, password string) (string, error)

	// GetAllDriverUsers is implemented by the usecasefacades.Userfacade that wraps getting all driver users usecase.
	GetAllDriverUsers(ctx context.Context) ([]entities.User, error)

	// BanDriver is implemented by the usecasefacades.Userfacade that wraps banning a driver usecase.
	BanDriver(ctx context.Context, id uuid.UUID) error
}

type ParkingLotFacade interface {
	CreateParkingLot(ctx context.Context, name string, capacity int) (entities.ParkingLot, error)
}

// ParkingRequestFacade is provides an interface implemented by usecasefacades.ParkingRequestFacade.
type ParkingRequestFacade interface {
	// CreateParkingRequest is implemented by usecasefacades.ParkingRequestFacade that wraps parking request creation usecase.
	CreateParkingRequest(ctx context.Context, parkingRequest *entities.ParkingRequest) (*entities.ParkingRequest, error)

	// UpdateParkingRequestStatus is implemented by usecasefacades.ParkingRequestFacade that wraps parking request update usecase.
	UpdateParkingRequestStatus(ctx context.Context, id uuid.UUID, status string) error

	// AssignParkingSpace is implemented by usecasefacades.ParkingRequestFacade that wraps parking space assigment usecase.
	AssignParkingSpace(ctx context.Context, requestID uuid.UUID, spaceID uuid.UUID) error

	// GetAllParkingRequests is implemented by usecasefacades.ParkingRequestFacade that wraps getting all parking requests.
	GetAllParkingRequests(ctx context.Context) ([]entities.ParkingRequest, error)
}

// Facade acts as a single entry point to access functionalities provided by all usecase facades.
type Facade struct {
	userFacade           UserFacade
	parkingRequestFacade ParkingRequestFacade
	parkingLotFacade     ParkingLotFacade
}

// NewFacade returns new instance of Facade.
func NewFacade(
	uFacade UserFacade,
	prFacade ParkingRequestFacade,
	plFacade ParkingLotFacade,
) *Facade {
	return &Facade{
		userFacade:           uFacade,
		parkingRequestFacade: prFacade,
		parkingLotFacade:     plFacade,
	}
}
