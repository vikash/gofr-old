package gofr

import (
	"context"
	"net/http"
)

type Request interface {
	Context() context.Context
	Param(string) string
	PathParam(string) string
	Bind(interface{}) error
	Header() http.Header
}
