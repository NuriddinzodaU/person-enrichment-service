package internal

import (
	"go.uber.org/fx"
	"person-service/internal/config"
	"person-service/internal/handlers"
	"person-service/internal/pkg/db"
	"person-service/internal/pkg/logger"
	"person-service/internal/pkg/repository"
	"person-service/internal/pkg/service"
)

var Modules = fx.Options(
	service.NewService,
	logger.NewLogger,
	handlers.NewHandler,
	repository.NewRepository,
	db.NewPostgres,
	config.NewConfig,
)
