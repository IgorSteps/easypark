package main

import (
	"github.com/IgorSteps/easypark/internal/drivers/httpserver"
	"github.com/sirupsen/logrus"
)

type App struct {
	logger *logrus.Logger
	server *httpserver.Server
}

func NewApp(s *httpserver.Server, l *logrus.Logger) *App {
	return &App{
		logger: l,
		server: s,
	}
}
