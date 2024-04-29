package main

import (
	"github.com/IgorSteps/easypark/internal/drivers/websocketserver"
	"github.com/sirupsen/logrus"
)

type App struct {
	logger *logrus.Logger
	server *websocketserver.Server
}

func NewApp(s *websocketserver.Server, l *logrus.Logger) *App {
	return &App{
		logger: l,
		server: s,
	}
}
