package main

import (
	"net/http"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/ejuju/crud-aws/internal/awsutil"
	"github.com/ejuju/crud-aws/internal/httputils"
)

func (s *Service) httpV1HandleGetImage(awsSession *session.Session, region, bucket string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get ID from request uri
		fileKey := s.HTTPRouter.ParseURIParams(r)["id"]
		if fileKey == "" {
			httputils.RespondJSON(w, http.StatusBadRequest, "file key not provided", s.Logger)
			return
		}

		// download desired image
		rawdata, err := awsutil.DownloadFile(awsSession, bucket, fileKey)
		if err != nil {
			httputils.RespondJSON(w, http.StatusInternalServerError, err.Error(), s.Logger)
			return
		}

		// return raw file data
		w.Header().Set("Content-Type", "image/png")
		w.Write(rawdata)
		w.WriteHeader(http.StatusOK)
	}
}
