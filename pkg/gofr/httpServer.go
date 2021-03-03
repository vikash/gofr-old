package gofr

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"
	http2 "github.com/zopsmart/ezgo/pkg/gofr/http"
)

type httpServer struct {
	router *http2.Router
	port   int
}

func (s *httpServer) Run(container *Container) {
	var srv *http.Server

	container.Logf("Starting server on port: %d\n", s.port)
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"GET", "HEAD", "PUT", "POST", "DELETE", "OPTIONS"},
		Debug:          false,
	})

	srv = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: cors.Handler(s.router),
	}

	container.Error(srv.ListenAndServe())
}
