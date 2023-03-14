package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type DefaultFieldHook struct {
}

func (h *DefaultFieldHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *DefaultFieldHook) Fire(e *logrus.Entry) error {
	e.Data["service_name"] = "Viska-Balita"
	return nil
}

func New() *logrus.Logger {
	logger := logrus.New()
	logger.Out = os.Stdout
	logger.Formatter = &logrus.JSONFormatter{
		DisableTimestamp: true,
	}
	logger.AddHook(&DefaultFieldHook{})

	return logger
}
