package httputils

import "net/http"

// StatusRecorderResponseWriter implements the http.ResponseWriter interface
type StatusRecorderResponseWriter struct {
	http.ResponseWriter
	Status int
}

// Store the status when it gets set by the caller
func (r *StatusRecorderResponseWriter) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}
