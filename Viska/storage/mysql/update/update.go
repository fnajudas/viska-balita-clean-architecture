package update

import (
	"context"
	"time"
	"viska/storage/models"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

const (
	defaultTimeOut = 5 * time.Second
)

type UpdateDataBalita struct {
	logger  *logrus.Logger
	db      *gorm.DB
	timeout time.Duration
}

func NewUpdateDataBalita(logger *logrus.Logger, db *gorm.DB, timeout time.Duration) *UpdateDataBalita {
	if timeout == 0 {
		timeout = defaultTimeOut
	}
	return &UpdateDataBalita{
		logger:  logger,
		db:      db,
		timeout: timeout,
	}
}

func (u *UpdateDataBalita) UpdateDataBalita(req models.UpdateBalita) (data models.RespGetDataBalita, err error) {
	logger := u.logger.WithFields(logrus.Fields{
		"Layer":     "Mysql-Uppdate",
		"Func Name": "UpdateDataBalita",
	})

	query := `
	UPDATE
		balita
	SET
		nama = ?
	WHERE
		id = ?
	`

	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	trx := u.db.BeginTx(ctx, nil)
	defer trx.Rollback()

	if err := trx.Exec(query, req.Nama, req.Id).Error; err != nil {
		logger.Errorf(`Error update data Balita: %v`, err)
		return data, err
	}

	if err := trx.Commit().Error; err != nil {
		logger.Errorf(`Error commit: %v`, err)
		return data, err
	}

	return data, nil
}
