package update

import (
	"viska/storage/models"

	"github.com/sirupsen/logrus"
)

type updateData interface {
	UpdateDataBalita(req models.UpdateBalita) (data models.RespGetDataBalita, err error)
}

type Executor struct {
	logger     *logrus.Logger
	updateData updateData
}

func NewExecutor(logger *logrus.Logger, updateData updateData) *Executor {
	return &Executor{
		logger:     logger,
		updateData: updateData,
	}
}

func (e *Executor) UpdateDataBalita(req models.UpdateBalita) (resp models.Template2, err error) {
	logger := e.logger.WithFields(logrus.Fields{
		"Layer":     "Service",
		"Func Name": "UpdateDataBalita",
	})

	dt, errs := e.updateData.UpdateDataBalita(req)
	if errs != nil {
		logger.Errorf(`Error delete: %v`, err)
		return resp, err
	}

	resp.Items = dt

	return resp, nil
}
