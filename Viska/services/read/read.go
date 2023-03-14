package read

import (
	"viska/storage/models"

	"github.com/sirupsen/logrus"
)

type listBalita interface {
	ListDataBalita(req models.ReqGetDataBalita) (data []models.RespGetDataBalita, err error)
}

type Executor struct {
	logger     *logrus.Logger
	listBalita listBalita
}

func NewExecutor(logger *logrus.Logger, listBalita listBalita) *Executor {
	return &Executor{
		logger:     logger,
		listBalita: listBalita,
	}
}

func (e *Executor) ListBalita(req models.ReqGetDataBalita) (resp models.Template2, err error) {
	logger := e.logger.WithFields(logrus.Fields{
		"Layer":     "Service",
		"Func Name": "GetDataBalita",
	})

	dt, err := e.listBalita.ListDataBalita(req)
	if err != nil {
		logger.Errorf(`Something error in Services`, err)
		return resp, err
	}
	resp.Items = dt

	return
}
