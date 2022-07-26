package main

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/ejuju/crud-aws/internal/awsutil"
	"github.com/ejuju/crud-aws/internal/httputils"
)

func (s *Service) httpV1HandleUploadImage(awsSession *session.Session, bucket string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// read image from request body
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			httputils.RespondJSON(w, http.StatusBadRequest, err.Error(), s.Logger)
			return
		}

		// upload to bucket
		key := time.Now().String()
		err = awsutil.UploadFile(awsSession, bucket, key, data)
		if err != nil {
			httputils.RespondJSON(w, http.StatusInternalServerError, err.Error(), s.Logger)
			return
		}

		// respond with upload object key
		httputils.RespondJSON(w, http.StatusOK, key, s.Logger)
	}
}
