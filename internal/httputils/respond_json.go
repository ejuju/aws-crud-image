package httputils

import (
	"encoding/json"
	"net/http"

	"github.com/ejuju/crud-aws/internal/logutil"
)

func RespondJSON(w http.ResponseWriter, status int, body interface{}, logger logutil.Logger) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		logger.Log(logutil.LogLevelErr, err.Error())
	}
}
