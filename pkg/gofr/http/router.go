package http

import (
	"net/http"

	"github.com/vikash/gofr/pkg/gofr/logging"

	"github.com/rs/cors"
	"github.com/vikash/gofr/pkg/gofr/http/middleware"

	"github.com/gorilla/mux"
)

type Router struct {
	mux.Router
}

func NewRouter() *Router {
	muxRouter := mux.NewRouter().StrictSlash(false)
	muxRouter.Use(
		middleware.DefaultHeaders,
		middleware.Tracer,
		middleware.Logging(logging.NewLogger(logging.INFO)),
	)

	cors := cors.New(cors.Options{
		AllowedOrigins:         []string{"*"},
		AllowOriginRequestFunc: func(r *http.Request, origin string) bool { return true },
		AllowedMethods:         []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:         []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:         []string{"Link"},
		AllowCredentials:       true,
		OptionsPassthrough:     true,
		MaxAge:                 3599, // Maximum value not ignored by any of major browsers
	})

	muxRouter.Use(cors.Handler)

	return &Router{
		Router: *muxRouter,
	}
}

func (rou *Router) Add(method, pattern string, handler http.Handler) {
	rou.Router.NewRoute().Methods(method).Path(pattern).Handler(handler)
}
