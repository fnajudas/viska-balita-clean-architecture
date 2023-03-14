package update

import (
	"encoding/json"
	"net/http"
	"viska/storage/models"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/thedevsaddam/renderer"
)

type updateData interface {
	UpdateDataBalita(req models.UpdateBalita) (resp models.Template2, err error)
}

type Handler struct {
	logger     *logrus.Logger
	render     *renderer.Render
	db         *gorm.DB
	updateData updateData
}

func NewHandler(logger *logrus.Logger, render *renderer.Render, db *gorm.DB, updateData updateData) *Handler {
	return &Handler{
		logger:     logger,
		render:     render,
		db:         db,
		updateData: updateData,
	}
}

func (h *Handler) UpdateDataBalita(w http.ResponseWriter, r *http.Request) {
	logger := h.logger.WithFields(logrus.Fields{
		"Layer":     "Controller",
		"Func Name": "UpdateDataBalita",
	})

	var body models.Balita
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		logger.Errorf(`Error: %v`, err)
		h.render.JSON(w, http.StatusBadRequest, &models.Template1{
			Message: "Bad request.",
			Data:    nil,
		})
		return
	}
	defer r.Body.Close()

	if body.Nama == "" || body.Nik == 0 {
		logger.Errorf(`Error.`)
		h.render.JSON(w, http.StatusBadRequest, &models.Template1{
			Message: "Bad request",
			Data:    nil,
		})
		return
	}

	var user models.Balita
	if h.db.Where("nik = ?", body.Nik).First(&user).RecordNotFound() {
		logger.Errorf(`Id not found`)
		h.render.JSON(w, http.StatusBadRequest, &models.Template1{
			Message: "Id not found",
			Data:    nil,
		})
		return
	}
	if err := h.db.Model(&user).Updates(&body).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			h.render.JSON(w, http.StatusBadRequest, &models.Template1{
				Message: "Bad request",
				Data:    nil,
			})
			return
		default:
			h.render.JSON(w, http.StatusInternalServerError, &models.Template1{
				Message: "Internal server error",
				Data:    nil,
			})
			return
		}
	}

	h.render.JSON(w, http.StatusOK, &models.Template1{
		Message: "Success",
		Data:    body,
	})
}
