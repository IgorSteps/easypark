package handlers

import (
	"context"
	"time"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/google/uuid"
)

// UserFacade is provides an interface implemented by usecasefacades.UserFacade.
type UserFacade interface {
	// CreateDriver is implemented by usecasefacades.UserFacade that wraps driver user creation usecase.
	CreateDriver(ctx context.Context, driver *entities.User) (*entities.User, error)

	// AuthoriseUser is implemented by usecasefacades.UserFacade that wraps user login usecase.
	AuthoriseUser(ctx context.Context, username, password string) (string, error)

	// GetAllDriverUsers is implemented by the usecasefacades.Userfacade that wraps getting all driver users usecase.
	GetAllDriverUsers(ctx context.Context) ([]entities.User, error)

	// BanDriver is implemented by the usecasefacades.Userfacade that wraps banning a driver usecase.
	BanDriver(ctx context.Context, id uuid.UUID) error
}

// ParkingLotFacade provides an interface implemented by usecasefacades.ParkingLotsFacade.
type ParkingLotFacade interface {
	// CreateParkingLot is implemented by usecasefacades.ParkingLotsFacade that wraps parking lot creation usecase.
	CreateParkingLot(ctx context.Context, name string, capacity int) (entities.ParkingLot, error)

	// GetAllPakringLots is implemented by usecasefacades.ParkingLotsFacade that wraps getting all parking lots usecase.
	GetAllParkingLots(ctx context.Context) ([]entities.ParkingLot, error)

	// DeleteParkingLot is implemented by usecasefacades.ParkingLotFacade that wraps deleting parking lot usecase.
	DeleteParkingLot(ctx context.Context, id uuid.UUID) error
}

// ParkingSpaceFacade provides an interface implemented by usecasefacades.ParkingSpaceFacade.
type ParkingSpaceFacade interface {
	// UpdateParkingSpaceStatus is implemented by usecasefacades.ParkingSpaceFacade that wraps updating parking space status usecase.
	UpdateParkingSpaceStatus(ctx context.Context, id uuid.UUID, status string) (entities.ParkingSpace, error)

	// GetSingleParkingSpace is implemented by usecasefacades.ParkingSpaceFacade that wraps getting single parking space usecase.
	GetSingleParkingSpace(ctx context.Context, id uuid.UUID) (entities.ParkingSpace, error)
}

// ParkingRequestFacade provides an interface implemented by usecasefacades.ParkingRequestFacade.
type ParkingRequestFacade interface {
	// CreateParkingRequest is implemented by usecasefacades.ParkingRequestFacade that wraps parking request creation usecase.
	CreateParkingRequest(ctx context.Context, parkingRequest *entities.ParkingRequest) (*entities.ParkingRequest, error)

	// UpdateParkingRequestStatus is implemented by usecasefacades.ParkingRequestFacade that wraps parking request update usecase.
	UpdateParkingRequestStatus(ctx context.Context, id uuid.UUID, status string) error

	// AssignParkingSpace is implemented by usecasefacades.ParkingRequestFacade that wraps parking space assigment usecase.
	AssignParkingSpace(ctx context.Context, requestID uuid.UUID, spaceID uuid.UUID) error

	// GetAllParkingRequests is implemented by usecasefacades.ParkingRequestFacade that wraps getting all parking requests.
	GetAllParkingRequests(ctx context.Context) ([]entities.ParkingRequest, error)

	// GetDriversParkingRequests is implemented by usecasefacades.ParkingRequestFacade that wraps getting parking requests for particular driver.
	GetDriversParkingRequests(ctx context.Context, id uuid.UUID) ([]entities.ParkingRequest, error)
}

// NotificationFacade provides an interface implemented by usecasefacades.NotificationFacade.
type NotificationFacade interface {
	// CreateNotification is implemented by usecasefacades.NotificationFacade that wraps creating a notification.
	CreateNotification(ctx context.Context, driverID, parkingReqID, spaceID uuid.UUID, location string, notificationType int) (entities.Notification, error)

	// GetAll is implemented by usecasefacades.NotificationFacade that wraps getting all notifications usecase.
	GetAllNotifications(ctx context.Context) ([]entities.Notification, error)
}

// AlertFacade provides an interface implemented by usecasefacades.AlertFacade.
type AlertFacade interface {
	// GetAlert is implemented by usecasefacades.AlertFacade that wraps get a single alert usecase.
	GetAlert(ctx context.Context, id uuid.UUID) (entities.Alert, error)

	GetAllAlerts(ctx context.Context) ([]entities.Alert, error)

	// CheckForLateArrivals is implemented by usecasefacades.AlertFacade that wraps checking for late arrivals usecase.
	CheckForLateArrivals(ctx context.Context, threshold time.Duration) ([]entities.Alert, error)
}

// Facade acts as a single entry point to access functionalities provided by all usecase facades.
type Facade struct {
	userFacade           UserFacade
	parkingRequestFacade ParkingRequestFacade
	parkingLotFacade     ParkingLotFacade
	parkingSpaceFacade   ParkingSpaceFacade
	notificationFacade   NotificationFacade
	alertFacade          AlertFacade
}

// NewFacade returns new instance of Facade.
func NewFacade(
	uFacade UserFacade,
	prFacade ParkingRequestFacade,
	plFacade ParkingLotFacade,
	psFacade ParkingSpaceFacade,
	nFacade NotificationFacade,
	aFacade AlertFacade,
) *Facade {
	return &Facade{
		userFacade:           uFacade,
		parkingRequestFacade: prFacade,
		parkingLotFacade:     plFacade,
		parkingSpaceFacade:   psFacade,
		notificationFacade:   nFacade,
		alertFacade:          aFacade,
	}
}
