package handlers

import (
	"go.uber.org/fx"
	"net/http"
	"person-service/internal/pkg/logger"
	"person-service/internal/pkg/service"
)

var NewHandler = fx.Provide(newHandler)

type IHandler interface {
	CreatePerson() http.HandlerFunc
	UpdatePerson() http.HandlerFunc
	GetPerson() http.HandlerFunc
	GetAllPerson() http.HandlerFunc
	DeletePerson() http.HandlerFunc
}

type dependencies struct {
	fx.In
	SVC    service.IService
	Logger logger.ILogger
}

type Handler struct {
	svc    service.IService
	Logger logger.ILogger
}

func newHandler(d dependencies) IHandler {
	return Handler{
		svc:    d.SVC,
		Logger: d.Logger,
	}
}
