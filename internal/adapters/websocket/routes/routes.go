package routes

import (
	"net/http"

	"github.com/IgorSteps/easypark/internal/adapters/websocket/client"
	"github.com/IgorSteps/easypark/internal/adapters/websocket/handlers"
	chiLogger "github.com/chi-middleware/logrus-logger"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func NewRouter(logger *logrus.Logger, hub *client.Hub) chi.Router {
	router := chi.NewRouter()
	router.Use(chiLogger.Logger("router", logger))

	// TODO: Add auth middleware.
	router.Method(http.MethodGet, "/ws/{id}", handlers.NewWebsocketHandler(logger, hub))

	return router
}
