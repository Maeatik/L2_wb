package handler

import (
	"net/http"
	"time"
)

type modelsDateReq struct {
	StartDay time.Time
}

func (h *Handler) ModelsForDayHandler(req *http.Request) APIResponse {
	data := &modelsDateReq{}
	if err := data.parse(req); err != nil {
		return h.Error(http.StatusBadRequest, err)
	}

	models, err := h.Storage.GetByPeriod(data.StartDay, data.StartDay.Add(time.Hour*24))

	if err != nil {
		return h.Error(http.StatusInternalServerError, err)
	}
	return h.JSON(http.StatusOK, models)
}

func (h *Handler) ModelsForMonthHandler(req *http.Request) APIResponse {
	data := &modelsDateReq{}
	if err := data.parse(req); err != nil {
		return h.Error(http.StatusBadRequest, err)
	}
	models, err := h.Storage.GetByPeriod(data.StartDay, data.StartDay.Add(time.Hour*24*30))
	if err != nil {
		return h.Error(http.StatusInternalServerError, err)
	}

	return h.JSON(http.StatusOK, models)
}

func (h *Handler) ModelsForWeekHandler(req *http.Request) APIResponse {
	data := &modelsDateReq{}
	if err := data.parse(req); err != nil {
		return h.Error(http.StatusBadRequest, err)
	}

	models, err := h.Storage.GetByPeriod(data.StartDay, data.StartDay.Add(time.Hour*24*7))
	if err != nil {
		return h.Error(http.StatusInternalServerError, err)
	}

	return h.JSON(http.StatusOK, models)
}

func (e *modelsDateReq) parse(req *http.Request) error {
	var err error

	if err = req.ParseForm(); err != nil {
		return err
	}

	if e.StartDay, err = time.Parse("2006-01-02", req.FormValue("start_day")); err != nil {
		return err
	}
	return nil
}
