package routes

import (
	"net/http"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/sirupsen/logrus"

	lgr "github.com/chi-middleware/logrus-logger"
	"github.com/go-chi/chi/v5"
)

// HandlerFactory defines an interface for creating different REST handlers.
type HandlerFactory interface {
	// User handlers.
	DriverCreate() http.Handler
	UserAuthorise() http.Handler
	GetAllDrivers() http.Handler
	DriverBan() http.Handler

	// Parking request handlers.
	ParkingRequestCreate() http.Handler
	ParkingRequestStatusUpdate() http.Handler
	AssignParkingSpace() http.Handler
	GetAllParkingRequests() http.Handler
	GetAllParkingRequestsForDriver() http.Handler

	// Parking lots handlers.
	ParkingLotCreate() http.Handler
	GetAllParkingLots() http.Handler
	DeleteParkingLot() http.Handler

	// Parking space handlers.
	GetSingleParkingSpace() http.Handler
	UpdateParkingSpaceStatus() http.Handler

	// Notification handlers
	CreateNotification() http.Handler
	GetAllNotifications() http.Handler

	// Alert handlers
	GetSingleAlert() http.Handler
	CheckForLateArrivals() http.Handler
	GetAllAlerts() http.Handler
}

// RequestAuthoriser defines an interfaces for middleware that authorises users' tokens.
//
// TODO: This should be split up and it could be a factory similar to handlers?
type Middleware interface {
	Authorise(next http.Handler) http.Handler
	RequireRole(requiredRole entities.UserRole) func(next http.Handler) http.Handler
	CheckStatus(next http.Handler) http.Handler
	CorsMiddleware(next http.Handler) http.Handler
}

// NewRouter constructs routes for our REST API.
func NewRouter(handlerFactory HandlerFactory, middleware Middleware, logger *logrus.Logger) chi.Router {
	router := chi.NewRouter()
	router.Use(lgr.Logger("router", logger))
	//router.Use(middleware.CorsMiddleware)

	// Public routes
	router.Method(http.MethodPost, "/register", handlerFactory.DriverCreate())
	router.Method(http.MethodPost, "/login", handlerFactory.UserAuthorise())

	// Driver routes
	router.Group(func(r chi.Router) {
		r.Use(middleware.Authorise, middleware.RequireRole(entities.RoleDriver), middleware.CheckStatus)
		// Parking requests
		r.Method(http.MethodPost, "/drivers/{id}/parking-requests", handlerFactory.ParkingRequestCreate())
		r.Method(http.MethodGet, "/drivers/{id}/parking-requests", handlerFactory.GetAllParkingRequestsForDriver())
		// Notifications
		r.Method(http.MethodPost, "/drivers/{id}/notifications", handlerFactory.CreateNotification())
		// Parking spaces
		r.Method(http.MethodGet, "/parking-spaces/{id}", handlerFactory.GetSingleParkingSpace())
		// Park lots
		r.Method(http.MethodGet, "/driver-parking-lots", handlerFactory.GetAllParkingLots())
	})

	// Admin routes
	router.Group(func(r chi.Router) {
		r.Use(middleware.Authorise, middleware.RequireRole(entities.RoleAdmin))
		// Drivers
		r.Method(http.MethodGet, "/drivers", handlerFactory.GetAllDrivers())
		r.Method(http.MethodPatch, "/drivers/{id}/status", handlerFactory.DriverBan())
		// Parking lots
		r.Method(http.MethodPost, "/parking-lots", handlerFactory.ParkingLotCreate())
		r.Method(http.MethodGet, "/parking-lots", handlerFactory.GetAllParkingLots())
		r.Method(http.MethodDelete, "/parking-lots/{id}", handlerFactory.DeleteParkingLot())
		// Parking requests
		r.Method(http.MethodPatch, "/parking-requests/{id}/status", handlerFactory.ParkingRequestStatusUpdate())
		r.Method(http.MethodPatch, "/parking-requests/{id}/space", handlerFactory.AssignParkingSpace())
		r.Method(http.MethodGet, "/parking-requests", handlerFactory.GetAllParkingRequests())
		// Parking spaces
		r.Method(http.MethodPatch, "/parking-spaces/{id}/status", handlerFactory.UpdateParkingSpaceStatus())
		r.Method(http.MethodGet, "/parking-spaces/{id}", handlerFactory.GetSingleParkingSpace())
		// Notifications
		r.Method(http.MethodGet, "/notifications", handlerFactory.GetAllNotifications())
		// Alerts
		r.Method(http.MethodGet, "/alerts/{id}", handlerFactory.GetSingleAlert())
		r.Method(http.MethodGet, "/alerts", handlerFactory.GetAllAlerts())
		r.Method(http.MethodPost, "/alerts/late-arrivals", handlerFactory.CheckForLateArrivals())
	})

	return router
}
