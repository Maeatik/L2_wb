package middleware

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

func RecoverMiddleware(log *logrus.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
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
			next.ServeHTTP(writer, request)
		})
	}
}
