package handler

import (
	"L2/develop/dev11/domain"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// Handler Структура обработчика. Содержит в себе хранилище с моделям и логгер для миддлвайра
type Handler struct {
	Storage domain.StorageInterface
	Logger  *logrus.Logger
}

type APIResponse func(resp http.ResponseWriter)

// Handle обработчик для установки response и request
func (h *Handler) Handle(fn func(req *http.Request) APIResponse) func(resp http.ResponseWriter, req *http.Request) {
	return func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Server", "GoCalendar")
		fn(req)(resp)
	}
}

// JSON отправляем json
func (h *Handler) JSON(code int, data interface{}) APIResponse {
	return h.sendJSON(code, map[string]interface{}{"result": data})
}

func (h *Handler) sendJSON(code int, data interface{}) APIResponse {
	var encData []byte
	var err error
	//если переданная модель не пустая парсим ее в JSON
	if data != nil {
		encData, err = json.Marshal(data)
		if err != nil {
			//Если не получилось, говорим, что на сервере произошла ошибка
			return h.Error(http.StatusInternalServerError, err)
		}
	}
	//если все хорошо, отправляем клиенту полученную/ые модель/и
	return func(resp http.ResponseWriter) {
		resp.Header().Set("Content-Type", "application/json")
		resp.Header().Set("Content-Length", strconv.Itoa(len(encData)))
		resp.WriteHeader(code)
		resp.Write(encData)
	}
}

//Если случилась ошибка, отправляем сообщение об ошибке
func (h *Handler) Error(code int, err error) APIResponse {
	return h.sendJSON(code, map[string]string{"error": err.Error()})
}
