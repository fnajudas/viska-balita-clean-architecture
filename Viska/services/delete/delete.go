package delete

import (
	"github.com/sirupsen/logrus"
)

type deleteBalita interface {
	DeleteBalita(id int) error
}

type Executor struct {
	logger       *logrus.Logger
	deleteBalita deleteBalita
}

func NewExecutor(logger *logrus.Logger, deleteBalita deleteBalita) *Executor {
	return &Executor{
		logger:       logger,
		deleteBalita: deleteBalita,
	}
}

func (e *Executor) DeleteBalita(id int) error {
	logger := e.logger.WithFields(logrus.Fields{
		"Layer":     "Service",
		"Func Name": "Service",
	})

	if err := e.deleteBalita.DeleteBalita(id); err != nil {
		logger.Errorf(`Error delete balita`)
		return err
	}

	return nil
}
