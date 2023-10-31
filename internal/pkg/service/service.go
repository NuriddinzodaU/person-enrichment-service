package service

import (
	"go.uber.org/fx"
	"person-service/internal/config"
	"person-service/internal/pkg/logger"
	"person-service/internal/pkg/repository"
)

var NewService = fx.Provide(newService)

type IService interface {
	LoggerInstance() logger.ILogger
	RepositoryInstance() repository.IRepository
	ConfigInstance() config.IConfig
}

type dependencies struct {
	fx.In
	Repository repository.IRepository
	Logger     logger.ILogger
	Config     config.IConfig
}

type service struct {
	Repository repository.IRepository
	Logger     logger.ILogger
	Config     config.IConfig
}

func newService(d dependencies) IService {
	return &service{
		d.Repository,
		d.Logger,
		d.Config,
	}
}

func (s service) LoggerInstance() logger.ILogger {
	return s.Logger
}

func (s service) RepositoryInstance() repository.IRepository {
	return s.Repository
}

func (s service) ConfigInstance() config.IConfig {
	return s.Config
}
