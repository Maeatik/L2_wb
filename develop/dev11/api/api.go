package api

import (
	"L2/develop/dev11/api/handler"
	"L2/develop/dev11/api/middleware"
	"L2/develop/dev11/domain"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"time"
)

//Структура конфигурации сервера
type ServerConfig struct {
	Host           string
	Port           string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
}

//Инициализация роутера, указание путей для запросов и реализации методов GET и POST
func initRouter(h *handler.Handler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", h.Handle(h.MainHandler)).Methods(http.MethodGet)
	router.HandleFunc("/create_model", h.Handle(h.ModelCreateHandler)).Methods(http.MethodPost)
	router.HandleFunc("/update_model", h.Handle(h.ModelUpdateHandler)).Methods(http.MethodPost)
	router.HandleFunc("/delete_model", h.Handle(h.ModelDeleteHandler)).Methods(http.MethodPost)
	router.HandleFunc("/models_for_day", h.Handle(h.ModelsForDayHandler)).Methods(http.MethodGet)
	router.HandleFunc("/models_for_week", h.Handle(h.ModelsForWeekHandler)).Methods(http.MethodGet)
	router.HandleFunc("/models_for_month", h.Handle(h.ModelsForMonthHandler)).Methods(http.MethodGet)

	return router

}

//Запуск сервера
func StartServer(storage domain.StorageInterface, config *ServerConfig) {
	//Создание логгера
	logger := logrus.New()
	//для того чтобы возвращать что-либо с сервера - парсим сообщение в JSON
	logger.SetFormatter(&logrus.JSONFormatter{})

	//Создаем обработчик
	h := &handler.Handler{
		Storage: storage,
		Logger:  logger,
	}
	//Создаем роутер
	router := initRouter(h)

	//Подключаем миддлвайр
	router.Use(middleware.RecoverMiddleware(h.Logger))
	router.Use(middleware.LoggerMiddleware(h.Logger))

	//создадим и настроим сервер
	server := &http.Server{
		Addr:           net.JoinHostPort(config.Host, config.Port),
		Handler:        router,
		ReadTimeout:    config.ReadTimeout,
		WriteTimeout:   config.WriteTimeout,
		MaxHeaderBytes: config.MaxHeaderBytes,
	}

	//Наконец-то запускаем сервер
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logger.Fatalf("HTTP server error: %s", err)
	}

}
