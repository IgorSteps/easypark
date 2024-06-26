package handlers

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// HandlerFactory implements HandlerFactory interface and helps provide dependencies for
// creating different REST Handlers.
type HandlerFactory struct {
	logger *logrus.Logger
	facade *Facade
}

// NewHandlerFactory provides new instance of the HandlerFactory.
func NewHandlerFactory(logger *logrus.Logger, facade *Facade) *HandlerFactory {
	return &HandlerFactory{
		logger: logger,
		facade: facade,
	}
}

// DriverCreate returns new REST handler for the creation of driver users.
func (s *HandlerFactory) DriverCreate() http.Handler {
	return NewDriverCreateHandler(s.facade.userFacade, s.logger)
}

// UserAuthorise returns new REST handler for user authentication.
func (s *HandlerFactory) UserAuthorise() http.Handler {
	return NewUserLoginHandler(s.facade.userFacade, s.logger)
}

// GetAllDrivers returns new REST handler for getting all driver users.
func (s *HandlerFactory) GetAllDrivers() http.Handler {
	return NewDriverUsersGetHandler(s.logger, s.facade.userFacade)
}

// DriverBan returns new REST handler to ban drivers.
func (s *HandlerFactory) DriverBan() http.Handler {
	return NewDriverStatusHandler(s.facade.userFacade, s.logger)
}

// ParkingRequestCreate returns new REST handler for creating parking requests.
func (s *HandlerFactory) ParkingRequestCreate() http.Handler {
	return NewParkingRequestCreateHandler(s.facade.parkingRequestFacade, s.logger)
}

// ParkingRequestStatusUpdate returns new REST handler to update parking request status.
func (s *HandlerFactory) ParkingRequestStatusUpdate() http.Handler {
	return NewParkingRequestStatusHandler(s.facade.parkingRequestFacade, s.logger)
}

// AssignParkingSpace returns a new REST handler to assign a space to a parking request.
func (s *HandlerFactory) AssignParkingSpace() http.Handler {
	return NewParkingRequestSpaceHandler(s.facade.parkingRequestFacade, s.logger)
}

// ParkingLotCreate returns a new REST handler to create parking lots.
func (s *HandlerFactory) ParkingLotCreate() http.Handler {
	return NewParkingLotCreateHandler(s.logger, s.facade.parkingLotFacade)
}

// GetAllParkingRequests returns a new REST handler to get all parking requests.
func (s *HandlerFactory) GetAllParkingRequests() http.Handler {
	return NewAllParkingRequestsGetHandler(s.logger, s.facade.parkingRequestFacade)
}

// GetAllParkingRequestsForDriver returns a new REST handler to get parking requests for a particular driver.
func (s *HandlerFactory) GetAllParkingRequestsForDriver() http.Handler {
	return NewDriversParkingRequestsGetHandler(s.logger, s.facade.parkingRequestFacade)
}

// GetAllParkingLots returns a new REST handler to get all parking lots.
func (s *HandlerFactory) GetAllParkingLots() http.Handler {
	return NewParkingLotGetAllHandler(s.logger, s.facade.parkingLotFacade)
}

// DeleteParkingLot returns a new REST handler to delete a parking lot.
func (s *HandlerFactory) DeleteParkingLot() http.Handler {
	return NewDeleteParkingLotHandler(s.logger, s.facade.parkingLotFacade)
}

// UpdateParkingSpaceStatus returns a new REST handler to update status of a parking space.
func (s *HandlerFactory) UpdateParkingSpaceStatus() http.Handler {
	return NewParkingSpaceStatusHandler(s.facade.parkingSpaceFacade, s.logger)
}

// CreateNotificaiton returns a new REST handler to create a notification.
func (s *HandlerFactory) CreateNotification() http.Handler {
	return NewNotificationCreateHandler(s.logger, s.facade.notificationFacade)
}

// GetSingleParkingSpace returns a new REST handler to get a single parking space.
func (s *HandlerFactory) GetSingleParkingSpace() http.Handler {
	return NewParkingSpaceGetSingleHandler(s.logger, s.facade.parkingSpaceFacade)
}

// GetAllNotifications returns a new REST handler to get all notifications.
func (s *HandlerFactory) GetAllNotifications() http.Handler {
	return NewNotificationGetAllHandler(s.logger, s.facade.notificationFacade)
}

// GetSingleAlert returns a new REST handler to get a single alert.
func (s *HandlerFactory) GetSingleAlert() http.Handler {
	return NewAlertGetSingleHandler(s.logger, s.facade.alertFacade)
}

// GetAllAlerts returns a new REST handler to get a all alerts.
func (s *HandlerFactory) GetAllAlerts() http.Handler {
	return NewAlertGetAllHandler(s.logger, s.facade.alertFacade)
}

// CheckForLateArrivals returns a new REST handler to check for late arrivals.
func (s *HandlerFactory) CheckForLateArrivals() http.Handler {
	return NewCheckLateArrivalHandler(s.logger, s.facade.alertFacade)
}

// CheckForOverStays returns a new REST handler to check for over staying users.
func (s *HandlerFactory) CheckForOverStays() http.Handler {
	return NewCheckOverStaysHandler(s.logger, s.facade.alertFacade)
}

// GetSingleParkingLot returns a new REST handler to get single parking lot.
func (s *HandlerFactory) GetSingleParkingLot() http.Handler {
	return NewParkingLotGetSingleHandler(s.facade.parkingLotFacade, s.logger)
}

// PaymentCreate returns a new REST handler to create payments
func (s *HandlerFactory) PaymentCreate() http.Handler {
	return NewPaymentCreateHandler(s.logger)
}

// AutomaticallyAssignParkingSpace returns a new REST handler to automatically assign parking spaces.
func (s *HandlerFactory) AutomaticallyAssignParkingSpace() http.Handler {
	return NewParkingRequestAutomaticSpaceHandler(s.facade.parkingRequestFacade, s.logger)
}
