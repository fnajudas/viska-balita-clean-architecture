package create

import (
	"time"
	"viska/storage/models"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

const (
	defaultTimeout = 5 * time.Second
)

type createBalita struct {
	db      *gorm.DB
	logger  *logrus.Logger
	timeout time.Duration
}

func NewCreateBalita(db *gorm.DB, logger *logrus.Logger, timeout time.Duration) *createBalita {
	if timeout == 0 {
		timeout = defaultTimeout
	}
	return &createBalita{
		db:      db,
		logger:  logger,
		timeout: timeout,
	}
}

func (c *createBalita) CreateBalita(req models.CreateDataBalita) (data models.CreateDataBalita, err error) {
	logger := c.logger.WithFields(logrus.Fields{
		"Layer":     "Mysql-Create",
		"Func Name": "CreateBalita",
	})

	// data.Alamat = req.Alamat
	// data.BeratBadan = req.BeratBadan
	// data.Birthday = req.Birthday
	// data.Image = req.Image
	// data.LingkarLenganAtas = req.LingkarLenganAtas
	// data.Nama = req.Nama
	// data.TinggiBadan = req.TinggiBadan
	// data.Nik = req.Nik

	query := `
	INSERT INTO
		balita
		(nik, nama, alamat, birthday, image, tinggi_badan, berat_badan, lingkar_lengan_atas, deleted)
	VALUES
		(?,?,?,?,?,?,?,?,?)
	`

	// if err := c.db.Raw(query, req.Nik, req.Nama, req.Alamat, req.Birthday, req.Image,
	// 	req.TinggiBadan, req.BeratBadan, req.LingkarLenganAtas, 0).First(data).Error; err != nil {
	// 	log.Print("===========================", err)
	// 	if err == gorm.ErrRecordNotFound {
	// 		logger.Errorf(`Error, record not found: %v`, err)
	// 		return data, err
	// 	}
	// 	logger.Errorf(`Error something; %v`, err)
	// 	return data, err
	// }
	// ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	// defer cancel()

	trx := c.db.Begin()
	defer trx.Rollback()

	if err := trx.Exec(query, req.Nik, req.Nama, req.Alamat, req.Birthday, req.Image, req.TinggiBadan, req.BeratBadan, req.LingkarLenganAtas, 0).Error; err != nil {
		logger.Errorf(`Error insert balita: %v`, err)
		return data, err
	}

	if err := trx.Commit().Error; err != nil {
		logger.Errorf(`Error commit: %v`, err)
		return data, err
	}

	return data, nil
}
