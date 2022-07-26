package main

import (
	"fmt"
	"net/http"

	"github.com/ejuju/crud-aws/internal/httputils"
	"github.com/ejuju/crud-aws/internal/logutil"
)

func (s *Service) NewHTTPPanicRecoveryMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				err := recover()
				if err != nil {
					_ = s.Logger.Log(logutil.LogLevelPanic, fmt.Sprintf("recovered from panic: %v", err))
					httputils.RespondJSON(w, http.StatusInternalServerError, nil, s.Logger)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
