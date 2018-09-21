package server

import (
	"fmt"
	"net/http"
	"time"

	"memefy/server/pkg/config"

	log "github.com/Sirupsen/logrus"
)

func accessLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		next.ServeHTTP(w, r)

		log.WithFields(log.Fields{
			"log_type":       "access",
			"remote_address": r.RemoteAddr,
			"protocol":       r.Proto,
			"request_method": r.Method,
			"uri":            r.URL.Path,
			"query_string":   r.URL.RawQuery,
			"bytes_received": r.ContentLength,
			"response_time":  uint64(time.Since(startTime) / time.Millisecond),
			"user_agent":     r.Header.Get("User-Agent"),
		}).Info("")
	})
}

func basicAuthMiddleware(config config.Security) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !config.EnableBasicAuth {
				next.ServeHTTP(w, r)
				return
			}

			realm := "admin"
			user, pass, ok := r.BasicAuth()
			if !ok {
				unauthorized(w, realm)
				return
			}
			valid := config.BasicAuthUser == user && config.BasicAuthPass == pass
			if !valid {
				unauthorized(w, realm)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func unauthorized(w http.ResponseWriter, realm string) {
	w.Header().Add("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, realm))
	w.WriteHeader(http.StatusUnauthorized)
}
