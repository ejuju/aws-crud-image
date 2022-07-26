package httputils

import "net/http"

type Router interface {
	http.Handler
	Route(method string, path string, handlerFunc http.HandlerFunc)
	HandlePath(path string, handler http.Handler)
	Wrap(middlewareFunc func(http.Handler) http.Handler)
	ParseURIParams(*http.Request) map[string]string
	RouteNotFound(http.HandlerFunc)
}
