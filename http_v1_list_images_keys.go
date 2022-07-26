package main

import (
	"net/http"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/ejuju/crud-aws/internal/awsutil"
	"github.com/ejuju/crud-aws/internal/httputils"
)

func (s *Service) httpV1HandleListImageKeys(awsSession *session.Session, region, bucket string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// retrive keys from aws
		keys, err := awsutil.ListFileKeys(awsSession, region, bucket)
		if err != nil {
			httputils.RespondJSON(w, http.StatusInternalServerError, err.Error(), s.Logger)
		}

		// send keys
		httputils.RespondJSON(w, http.StatusOK, keys, s.Logger)
	}
}
