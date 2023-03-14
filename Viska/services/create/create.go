package create

import (
	"viska/storage/models"

	"github.com/sirupsen/logrus"
)

type createBalita interface {
	CreateBalita(req models.CreateDataBalita) (data models.CreateDataBalita, err error)
}

type Executor struct {
	logger       *logrus.Logger
	createBalita createBalita
}

func NewExecutor(logger *logrus.Logger, createBalita createBalita) *Executor {
	return &Executor{
		logger:       logger,
		createBalita: createBalita,
	}
}

func (e *Executor) CreateBalitaService(req models.CreateDataBalita) (resp models.Template2, err error) {
	logger := e.logger.WithFields(logrus.Fields{
		"Layer":     "Service",
		"Func Name": "CreateBalita",
	})

	dt, err := e.createBalita.CreateBalita(req)
	if err != nil {
		logger.Errorf(`Error service: %v`, err)
		return resp, err
	}

	resp.Items = dt

	return
}
