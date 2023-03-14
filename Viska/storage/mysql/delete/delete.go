package delete

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

const (
	defaultTimeout = 5 * time.Second
)

type deleteBalita struct {
	logger  *logrus.Logger
	db      *gorm.DB
	timeout time.Duration
}

func NewDeleteBalita(logger *logrus.Logger, db *gorm.DB, timeout time.Duration) *deleteBalita {
	if timeout == 0 {
		timeout = defaultTimeout
	}
	return &deleteBalita{
		logger:  logger,
		db:      db,
		timeout: timeout,
	}
}

func (d *deleteBalita) DeleteBalita(id int) error {
	logger := d.logger.WithFields(logrus.Fields{
		"Layer":     "Mysql-Delete",
		"Func Name": "DeleteBalita",
	})

	query := `
	DELETE FROM
		balita
	WHERE
		id = ?
	`

	// ctx, cancel := context.WithTimeout(context.Background(), d.timeout)
	// defer cancel()

	trx := d.db.Begin()
	defer trx.Rollback()

	if err := trx.Exec(query, id).Error; err != nil {
		logger.Errorf(`Error insert balita: %v`, err)
		return err
	}

	if err := trx.Commit().Error; err != nil {
		logger.Errorf(`Error commit`)
		return err
	}

	return nil
}
