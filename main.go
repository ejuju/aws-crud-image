package main

import (
	"net/http"
	"os"
	"strconv"
	"time"

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

	// init service
	service := &Service{
		HTTPRouter: httpRouter,
		Logger:     logger,
	}

	// register http endpoints
	httpRouter.Route("/v1/image", http.MethodPost, func(w http.ResponseWriter, r *http.Request) {})
	httpRouter.Route("/v1/image/{id}", http.MethodGet, func(w http.ResponseWriter, r *http.Request) {})
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
	err := httpServer.ListenAndServe()
	if err != nil {
		logger.Log(logutil.LogLevelPanic, err.Error())
		os.Exit(1)
	}
}
