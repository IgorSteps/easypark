package routes

import (
	"net/http"

	"github.com/IgorSteps/easypark/internal/adapters/websocket/handlers"
	chiLogger "github.com/chi-middleware/logrus-logger"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func NewRouter(logger *logrus.Logger) chi.Router {
	router := chi.NewRouter()
	router.Use(chiLogger.Logger("router", logger))

	// TODO: Add auth middleware.
	router.Method(http.MethodPost, "/ws", handlers.NewWebsocketHandler(logger))

	return router
}
