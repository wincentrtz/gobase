package middleware

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		makeLogEntry(r).Info("incoming request")
		next.ServeHTTP(w, r)
	})
}

func makeLogEntry(r *http.Request) *log.Entry {
	return log.WithFields(log.Fields{
		"method": r.Method,
		"uri":    r.URL,
		"ip":     r.RemoteAddr,
	})
}
