package gofr

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"

	http2 "github.com/zopsmart/ezgo/pkg/gofr/http"
)

type httpServer struct {
	router *http2.Router
	port   int
}

func (s *httpServer) Run(container *Container) {
	var srv *http.Server

	container.Logf("Starting server on port: %d\n", s.port)
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "PUT", "POST", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	srv = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: handlers.CORS(headers, methods, origins)(s.router),
	}

	container.Error(srv.ListenAndServe())
}
