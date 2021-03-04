package http

import (
	"net/http"

	"github.com/vikash/gofr/pkg/gofr/logging"

	"github.com/vikash/gofr/pkg/gofr/http/middleware"

	"github.com/gorilla/mux"
)

type Router struct {
	mux.Router
}

func NewRouter() *Router {
	muxRouter := mux.NewRouter().StrictSlash(false)
	muxRouter.Use(
		middleware.Tracer,
		middleware.Logging(logging.NewLogger(logging.INFO)),
	)

	return &Router{
		Router: *muxRouter,
	}
}

func (rou *Router) Add(method, pattern string, handler http.Handler) {
	rou.Router.NewRoute().Methods(method).Path(pattern).Handler(handler)
}

func (rou *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, OPTIONS")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	rou.Router.ServeHTTP(w, r)
}
