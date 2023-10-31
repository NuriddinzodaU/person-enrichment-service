package repository

import (
	"person-service/internal/models"
	"person-service/internal/pkg/db"
	"person-service/internal/pkg/logger"

	"go.uber.org/fx"
)

var NewRepository = fx.Provide(newRepository)

type IRepository interface {
	UpdatePerson(person *models.Person) error
	GetPerson(id int64) (person *models.Person, err error)
	GetAllPerson(count, page int) (person []*models.Person, err error)
	CreatePerson(person *models.Person) error
	DeletePerson(id int) error
}

type dependencies struct {
	fx.In
	Postgres db.IPostgres
	Logger   logger.ILogger
}

type repository struct {
	Postgres db.IPostgres
	Logger   logger.ILogger
}

func newRepository(dp dependencies) IRepository {
	return &repository{
		Postgres: dp.Postgres,
		Logger:   dp.Logger,
	}
}
