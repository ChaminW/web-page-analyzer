package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// LoggingMiddleware creates a new logging middleware
func LoggingMiddleware(logger *logrus.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.WithFields(logrus.Fields{
				"method": r.Method,
				"path":   r.URL.Path,
				"ip":     r.RemoteAddr,
			}).Info("HTTP Request")
			next.ServeHTTP(w, r)
		})
	}
}
