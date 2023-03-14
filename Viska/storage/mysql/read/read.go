package read

import (
	"fmt"
	"viska/storage/models"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type ListDataBalita struct {
	logger *logrus.Logger
	db     *gorm.DB
}

func NewDataBalita(logger *logrus.Logger, db *gorm.DB) *ListDataBalita {
	return &ListDataBalita{
		logger: logger,
		db:     db,
	}
}

func (g *ListDataBalita) ListDataBalita(req models.ReqGetDataBalita) (data []models.RespGetDataBalita, err error) {
	logger := g.logger.WithFields(logrus.Fields{
		"Layer":     "Storage-Mysql",
		"Func Name": "GetDataBalita",
	})

	query := `
	SELECT
		*
	FROM
		balita
	WHERE
		deleted = 0
	`

	if req.NIK != 0 {
		query = query + fmt.Sprintf(` AND NIK = %d`, req.NIK)
	}

	if req.Nama != "" {
		query = query + fmt.Sprintln(`AND Nama LIKE '%`+req.Nama+`%'`)
	}

	listData := []models.RespGetDataBalita{}
	if err := g.db.Raw(query).Scan(&listData).Error; err != nil {
		logger.Errorf(`Query error: %v`, err)
		return data, err
	}

	data = listData
	return data, nil
}
