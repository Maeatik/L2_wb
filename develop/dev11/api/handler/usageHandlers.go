package handler

import (
	"L2/develop/dev11/domain"
	"errors"
	"fmt"
	"net/http"
	"time"
)

//Модели для запросов и респонсов

type modelCreateReq struct {
	Id       string
	Title    string
	DateFrom time.Time
	DateTo   time.Time
}
type modelDeleteReq struct {
	Id string
}

type modelUpdateReq struct {
	Id       string
	Title    string
	DateFrom time.Time
	DateTo   time.Time
}

//Указываем теги JSON'а для отправки клиенту
type modelCreateResp struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	DateFrom string `json:"date_from"`
	DateTo   string `json:"date_to"`
}

type modelUpdateResp struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	DateFrom string `json:"date_from"`
	DateTo   string `json:"date_to"`
}

// MainHandler Обрабочик пустого запроса
func (h *Handler) MainHandler(_ *http.Request) APIResponse {
	return h.JSON(http.StatusOK, map[string]string{"main": "/ server route"})
}

// ModelCreateHandler Обработчик запроса на создание модели
func (h *Handler) ModelCreateHandler(req *http.Request) APIResponse {
	data := &modelCreateReq{}
	//парсим запрос, еслм ошибка, говорим, что запрос был очень-очень плохим
	if err := data.parse(req); err != nil {
		return h.Error(http.StatusBadRequest, err)
	}
	//Создаем новую модель
	model := domain.NewModel(data.Id, data.Title, data.DateFrom, data.DateTo)

	//если не получилось ее сохранить - сообщаем, что на сервере произошла ошибка
	if err := h.Storage.Save(model); err != nil {
		h.Error(http.StatusInternalServerError, err)
	}

	//отправляем JSON
	return h.JSON(http.StatusCreated, &modelCreateResp{
		Id:       model.Id,
		Title:    model.Title,
		DateFrom: model.DateFrom.Format(time.RFC3339),
		DateTo:   model.DateTo.Format(time.RFC3339),
	})
}

//Парсер запроса на создание модели
func (e *modelCreateReq) parse(req *http.Request) error {
	var err error
	if err = req.ParseForm(); err != nil {
		return err
	}

	e.Id = req.FormValue("id")
	fmt.Println("e", e.Id)
	if e.Id == "" {
		return errors.New("model id is required")
	}

	e.Title = req.FormValue("title")
	fmt.Println("e", e.Title)
	if e.Title == "" {
		return errors.New("model title is required")
	}

	if e.DateFrom, err = time.Parse(time.RFC3339, req.FormValue("date_from")); err != nil {
		return err
	}

	if e.DateTo, err = time.Parse(time.RFC3339, req.FormValue("date_to")); err != nil {
		return err
	}

	if e.DateFrom.After(e.DateTo) {
		return errors.New("why dateLate after dateTo")
	}
	return nil
}

// ModelDeleteHandler Запрос на удаление модели
func (h *Handler) ModelDeleteHandler(req *http.Request) APIResponse {
	data := &modelDeleteReq{}

	if err := data.parse(req); err != nil {
		return h.Error(http.StatusBadRequest, err)
	}
	//получаем модель по айди
	model, err := h.Storage.GetById(data.Id)

	if err != nil {
		return h.Error(http.StatusInternalServerError, err)
	}

	//Если модель не нашлась, то 404
	if model == nil {
		return h.Error(http.StatusNotFound, errors.New("model not found"))
	}
	//если нашлось, то 200 и обнуляем модель
	return h.sendJSON(http.StatusAccepted, nil)
}

//поиск модели по айди
func (e *modelDeleteReq) parse(req *http.Request) error {

	if err := req.ParseForm(); err != nil {
		return err
	}
	e.Id = req.FormValue("id")
	if e.Id == "" {
		return errors.New("model ID is required")
	}

	return nil
}

//обрабочик изменения модели
func (h *Handler) ModelUpdateHandler(req *http.Request) APIResponse {
	data := &modelUpdateReq{}
	if err := data.parse(req); err != nil {
		return h.Error(http.StatusBadRequest, err)
	}

	model, err := h.Storage.GetById(data.Id)

	if err != nil {
		return h.Error(http.StatusNotFound, errors.New("404 where your model"))
	}

	model.Title = data.Title
	model.DateFrom = data.DateFrom
	model.DateTo = data.DateTo

	if err = h.Storage.Save(model); err != nil {
		return h.Error(http.StatusInternalServerError, err)
	}

	return h.JSON(http.StatusOK, &modelUpdateResp{
		Id:       model.Id,
		Title:    model.Title,
		DateFrom: model.DateFrom.Format(time.RFC3339),
		DateTo:   model.DateTo.Format(time.RFC3339),
	})
}

func (e *modelUpdateReq) parse(req *http.Request) error {
	var err error

	if err = req.ParseForm(); err != nil {
		return err
	}

	e.Id = req.FormValue("id")
	if e.Id == "" {
		return errors.New("model id is required")
	}

	e.Title = req.FormValue("title")
	if e.Title == "" {
		return errors.New("model title is required")
	}

	if e.DateFrom, err = time.Parse(time.RFC3339, req.FormValue("date_from")); err != nil {
		return err
	}

	if e.DateTo, err = time.Parse(time.RFC3339, req.FormValue("date_to")); err != nil {
		return err
	}

	if e.DateFrom.After(e.DateTo) {
		return errors.New("why dateLate after dateTo")
	}
	return nil
}
