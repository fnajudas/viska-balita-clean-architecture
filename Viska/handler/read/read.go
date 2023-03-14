package read

import (
	"net/http"
	"strconv"
	"viska/storage/models"

	"github.com/sirupsen/logrus"
	"github.com/thedevsaddam/renderer"
)

type listBalita interface {
	ListBalita(req models.ReqGetDataBalita) (resp models.Template2, err error)
}

type Handler struct {
	logger     *logrus.Logger
	listBalita listBalita
	render     *renderer.Render
}

func NewHandler(logger *logrus.Logger, listBalita listBalita, render *renderer.Render) *Handler {
	return &Handler{
		logger:     logger,
		listBalita: listBalita,
		render:     render,
	}
}

func (h *Handler) GetDataBalita(w http.ResponseWriter, r *http.Request) {
	logger := h.logger.WithFields(logrus.Fields{
		"Layer":     "Handler",
		"Func Name": "GetDataBalita",
	})

	param := r.URL.Query()
	Nik, _ := strconv.Atoi(param.Get("nik"))
	Nama := param.Get("nama")

	req := models.ReqGetDataBalita{
		NIK:  Nik,
		Nama: Nama,
	}

	data, err := h.listBalita.ListBalita(req)
	if err != nil {
		logger.Errorf(`Something error in handler: %v`, err)
		h.render.JSON(w, http.StatusBadRequest, models.Template1{
			Message: "Error",
			Data:    nil,
		})
		return
	}

	h.render.JSON(w, http.StatusOK, models.Template1{
		Message: "Success",
		Data:    data,
	})
}
