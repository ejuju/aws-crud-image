package main

import (
	"net/http"

	"github.com/ejuju/crud-aws/internal/httputils"
)

func (s *Service) httpHandleNotFound() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httputils.RespondJSON(w, http.StatusNotFound, r.URL.String()+"not found", s.Logger)
	}
}
