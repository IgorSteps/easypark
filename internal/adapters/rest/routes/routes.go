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
	DriverCreate() http.Handler
	UserAuthorise() http.Handler
	GetAllDrivers() http.Handler
	DriverBan() http.Handler
}

// RequestAuthoriser defines an interfaces for middleware that authorises users' tokens.
//
// TODO: This could be a factory similar to handlers, but we aren't using too many middlewares
type Middleware interface {
	Authorise(next http.Handler) http.Handler
	RequireRole(requiredRole entities.UserRole) func(next http.Handler) http.Handler
	CheckStatus(next http.Handler) http.Handler
}

// NewRouter constructs routes for our REST API.
func NewRouter(handlerFactory HandlerFactory, middleware Middleware, logger *logrus.Logger) chi.Router {
	router := chi.NewRouter()
	router.Use(lgr.Logger("router", logger))

	// Public routes
	router.Method(http.MethodPost, "/register", handlerFactory.DriverCreate())
	router.Method(http.MethodPost, "/login", handlerFactory.UserAuthorise())

	// Driver routes
	router.Group(func(r chi.Router) {
		r.Use(middleware.Authorise, middleware.RequireRole(entities.RoleDriver), middleware.CheckStatus)
		// Placeholder:
		r.Get("/driver", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Welcome, Driver!"))
		})
	})

	// Admin routes
	router.Group(func(r chi.Router) {
		r.Use(middleware.Authorise, middleware.RequireRole(entities.RoleAdmin))
		r.Method(http.MethodGet, "/drivers", handlerFactory.GetAllDrivers())
		r.Method(http.MethodPatch, "/drivers/{id}/status", handlerFactory.DriverBan())
	})

	return router
}
