package middleware

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

// RecoverMiddleware обрабатываем, если возникает паника
func RecoverMiddleware(log *logrus.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			//вызываем рековери-функцию в отложнной функции
			defer func() {
				var rec any
				if rec = recover(); rec != nil {
					var err, ok = rec.(error)
					if !ok {
						err = fmt.Errorf("%v", rec)
					}
					log.Errorf("panic: %v\n", err)
					writer.WriteHeader(http.StatusInternalServerError)
				}
			}()
			//если была паника, сообщаем о ней и отправляем ошибку на сервере
			next.ServeHTTP(writer, request)
		})
	}
}
