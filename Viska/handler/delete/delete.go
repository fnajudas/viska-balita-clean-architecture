package delete

import (
	"encoding/json"
	"net/http"
	"viska/storage/models"

	"github.com/sirupsen/logrus"
	"github.com/thedevsaddam/renderer"
)

type deleteBalita interface {
	DeleteBalita(id int) error
}

type Handler struct {
	logger       *logrus.Logger
	render       *renderer.Render
	deleteBalita deleteBalita
}

func NewHandler(logger *logrus.Logger, render *renderer.Render, deleteBalita deleteBalita) *Handler {
	return &Handler{
		logger:       logger,
		render:       render,
		deleteBalita: deleteBalita,
	}
}

func (h *Handler) DeleteBalita(w http.ResponseWriter, r *http.Request) {
	logger := h.logger.WithFields(logrus.Fields{
		"Layer":     "Handler",
		"Func Name": "DeleteBalita",
	})

	// param := mux.Vars(r)
	// id, err := strconv.Atoi(param["id"])
	// if err != nil {
	// 	logger.Errorf(`Id dont nil`)
	// 	h.render.JSON(w, http.StatusBadRequest, &models.Template1{
	// 		Message: "Id dont nil",
	// 		Data:    nil,
	// 	})
	// 	return
	// }

	body := models.DeteleBalita{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logger.Errorf(`Error decode body: %v`, err)
		h.render.JSON(w, http.StatusBadRequest, &models.Template1{
			Message: "Error decode body",
			Data:    nil,
		})
		return
	}

	if body.Id == 0 {
		logger.Errorf(`Id dont nil`)
		h.render.JSON(w, http.StatusBadRequest, &models.Template1{
			Message: "Id dont nil",
			Data:    nil,
		})
		return
	}

	if err := h.deleteBalita.DeleteBalita(body.Id); err != nil {
		logger.Errorf(`Error delete balita`)
		h.render.JSON(w, http.StatusInternalServerError, &models.Template1{
			Message: "Error delete balita",
			Data:    nil,
		})
		return
	}

	h.render.JSON(w, http.StatusOK, &models.Template1{
		Message: "Success delete balita",
		Data:    "Success delete balita",
	})
}
