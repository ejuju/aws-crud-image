package httputils

import (
	"net/http"

	"github.com/gorilla/mux"
)

type GorillaRouter struct {
	router *mux.Router
}

func NewGorillaRouter() *GorillaRouter {
	router := mux.NewRouter()
	return &GorillaRouter{router: router}
}

func (gorillaRouter *GorillaRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	gorillaRouter.router.ServeHTTP(w, r)
}

func (gorillaRouter *GorillaRouter) Route(path, method string, handlerFunc http.HandlerFunc) {
	gorillaRouter.router.HandleFunc(path, handlerFunc).Methods(method)
}

func (gorillaRouter *GorillaRouter) Wrap(middleware func(http.Handler) http.Handler) {
	gorillaRouter.router.Use(middleware)
}

func (gorillaRouter *GorillaRouter) ParseURIParams(r *http.Request) map[string]string {
	return mux.Vars(r)
}

func (gorillaRouter *GorillaRouter) RouteNotFound(handlerFunc http.HandlerFunc) {
	// this looks a bit weird but it is done so the 404 appears in the request logger middleware
	gorillaRouter.router.NotFoundHandler = gorillaRouter.router.NewRoute().HandlerFunc(http.NotFound).GetHandler()
}

func (gorillaRouter *GorillaRouter) HandlePath(path string, handler http.Handler) {
	gorillaRouter.router.Handle(path, handler)
}
