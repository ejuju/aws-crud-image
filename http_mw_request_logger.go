package main

import (
	"net/http"
	"strconv"

	"github.com/ejuju/crud-aws/internal/httputils"
	"github.com/ejuju/crud-aws/internal/logutil"
)

// NewHTTPRequestLoggerMiddleware logs requests with the response status and URL path
func (s *Service) NewHTTPRequestLoggerMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			statusRecorderResponseWriter := &httputils.StatusRecorderResponseWriter{ResponseWriter: w}
			next.ServeHTTP(statusRecorderResponseWriter, r)
			logBody := r.Method + " " + strconv.Itoa(statusRecorderResponseWriter.Status) + " " + r.URL.Path
			_ = s.Logger.Log(logutil.LogLevelInfo, logBody)
		})
	}
}
