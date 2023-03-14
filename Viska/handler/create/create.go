package create

import (
	"encoding/json"
	"net/http"
	"viska/storage/models"

	"github.com/sirupsen/logrus"
	"github.com/thedevsaddam/renderer"
)

type createBalita interface {
	CreateBalitaService(req models.CreateDataBalita) (resp models.Template2, err error)
}

type Handler struct {
	logger       *logrus.Logger
	render       *renderer.Render
	createBalita createBalita
}

func NewHandler(logger *logrus.Logger, render *renderer.Render, createBalita createBalita) *Handler {
	return &Handler{
		logger:       logger,
		render:       render,
		createBalita: createBalita,
	}
}

func (h *Handler) CreateBalitaHandler(w http.ResponseWriter, r *http.Request) {
	logger := h.logger.WithFields(logrus.Fields{
		"Layer":     "Handler",
		"Func Name": "CreateBalitaHandler",
	})

	body := models.CreateDataBalita{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logger.Errorf(`Error decode body: %v`, err)
		h.render.JSON(w, http.StatusBadRequest, &models.Template1{
			Message: "Error decode body",
			Data:    nil,
		})
		return
	}

	if body.Nik == 0 || body.Nama == "" || body.Alamat == "" || body.Image == "" || body.TinggiBadan == "" || body.BeratBadan == "" || body.LingkarLenganAtas == "" {
		logger.Errorf(`Error, data not valid`)
		h.render.JSON(w, http.StatusBadRequest, &models.Template1{
			Message: "Error, data not valid",
			Data:    nil,
		})
		return
	}

	_, err := h.createBalita.CreateBalitaService(body)
	if err != nil {
		logger.Errorf(`Error, create balita`)
		h.render.JSON(w, http.StatusInternalServerError, &models.Template1{
			Message: "Error, create balita",
			Data:    nil,
		})
		return
	}

	h.render.JSON(w, http.StatusOK, &models.Template1{
		Message: "Succes, create balita",
		Data:    body,
	})
}
