package main

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/ejuju/crud-aws/internal/httputils"
	"github.com/ejuju/crud-aws/internal/logutil"
)

type Service struct {
	HTTPRouter httputils.Router
	Logger     logutil.Logger
}

func main() {
	// init logger
	logger := logutil.NewDefaultLogger(logutil.DefaultLoggerConfig{})

	// init router
	httpRouter := httputils.NewGorillaRouter()

	// connect to AWS
	awsS3Region := os.Getenv("AWS_S3_REGION")
	if awsS3Region == "" {
		logger.Log(logutil.LogLevelPanic, "missing AWS_S3_REGION environment variable")
		os.Exit(1)
	}
	awsS3Bucket := os.Getenv("AWS_S3_BUCKET")
	if awsS3Bucket == "" {
		logger.Log(logutil.LogLevelPanic, "missing AWS_S3_BUCKET environment variable")
		os.Exit(1)
	}
	session, err := session.NewSession(&aws.Config{Region: aws.String(awsS3Region)})
	if err != nil {
		logger.Log(logutil.LogLevelPanic, err.Error())
		os.Exit(1)
	}

	// init service
	service := &Service{
		HTTPRouter: httpRouter,
		Logger:     logger,
	}

	// register http endpoints
	httpRouter.Route("/v1/image", http.MethodPost, service.httpV1HandleUploadImage(session, awsS3Bucket))
	httpRouter.Route("/v1/image/all", http.MethodGet, service.httpV1HandleListImageKeys(session, awsS3Region, awsS3Bucket))
	httpRouter.Route("/v1/image/{id}", http.MethodGet, service.httpV1HandleGetImage(session, awsS3Region, awsS3Bucket))
	httpRouter.Route("/v1/image/{id}", http.MethodPut, func(w http.ResponseWriter, r *http.Request) {})
	httpRouter.Route("/v1/image/{id}", http.MethodDelete, func(w http.ResponseWriter, r *http.Request) {})
	httpRouter.RouteNotFound(service.httpHandleNotFound())

	// register http middleware
	httpRouter.Wrap(service.NewHTTPRequestLoggerMiddleware())
	httpRouter.Wrap(service.NewHTTPPanicRecoveryMiddleware())

	// start http server and listen for incoming connections
	httpServer := http.Server{
		Addr:              ":" + strconv.Itoa(8080),
		Handler:           httpRouter,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		ReadHeaderTimeout: 60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    8192,
	}
	logger.Log(logutil.LogLevelInfo, "starting HTTP server ("+httpServer.Addr+")")
	err = httpServer.ListenAndServe()
	if err != nil {
		logger.Log(logutil.LogLevelPanic, err.Error())
		os.Exit(1)
	}
}
