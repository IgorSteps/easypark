package main

import (
	"github.com/IgorSteps/easypark/internal/adapters/websocket/client"
	"github.com/IgorSteps/easypark/internal/drivers/websocketserver"
	"github.com/sirupsen/logrus"
)

type App struct {
	logger *logrus.Logger
	server *websocketserver.Server
	hub    *client.Hub
}

func NewApp(s *websocketserver.Server, l *logrus.Logger, h *client.Hub) *App {
	return &App{
		logger: l,
		server: s,
		hub:    h,
	}
}
