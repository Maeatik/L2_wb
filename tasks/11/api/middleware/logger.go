package middleware

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func LoggerMiddleware(log *logrus.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			latency := time.Now()
			next.ServeHTTP(resp, req)
			log.WithFields(logrus.Fields{
				"method":  req.Method,
				"uri":     req.RequestURI,
				"host":    req.Host,
				"ip":      req.RemoteAddr,
				"latency": time.Since(latency).Seconds(),
			}).Info()
		})
	}
}
