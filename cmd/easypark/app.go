package main

import (
	"github.com/IgorSteps/easypark/internal/drivers/httpserver"
	"github.com/IgorSteps/easypark/internal/drivers/scheduler"
	"github.com/sirupsen/logrus"
)

type App struct {
	logger    *logrus.Logger
	server    *httpserver.Server
	scheduler *scheduler.Scheduler
}

func NewApp(s *httpserver.Server, l *logrus.Logger, scheduler *scheduler.Scheduler) *App {
	return &App{
		logger:    l,
		server:    s,
		scheduler: scheduler,
	}
}
