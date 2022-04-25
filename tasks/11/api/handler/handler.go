package handler

import (
	"L2-maste/tasks/11/domain"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type Handler struct {
	Storage domain.StorageInterface
	Logger  *logrus.Logger
}

type APIResponse func(resp http.ResponseWriter)

func (h *Handler) Handle(fn func(req *http.Request) APIResponse) func(resp http.ResponseWriter, req *http.Request) {
	return func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Server", "GoCalendar")
		fn(req)(resp)
	}
}

func (h *Handler) JSON(code int, data interface{}) APIResponse {
	return h.sendJSON(code, map[string]interface{}{"result": data})
}

func (h *Handler) sendJSON(code int, data interface{}) APIResponse {
	var encData []byte
	var err error

	if data != nil {
		encData, err = json.Marshal(data)
		if err != nil {
			return h.Error(http.StatusInternalServerError, err)
		}
	}
	return func(resp http.ResponseWriter) {
		resp.Header().Set("Content-Type", "application/json")
		resp.Header().Set("Content-Length", strconv.Itoa(len(encData)))
		resp.WriteHeader(code)
		resp.Write(encData)
	}
}

func (h *Handler) Error(code int, err error) APIResponse {
	return h.sendJSON(code, map[string]string{"error1": err.Error()})
}
