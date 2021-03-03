package gofr

import (
	"context"
	"mime/multipart"
)

type Request interface {
	Context() context.Context
	Param(string) string
	PathParam(string) string
	Bind(interface{}) error
	FormFile(string) (multipart.File, *multipart.FileHeader, error)
	ParseForm() error
	FormValue(string) string
	ParseMultiPartForm(int64) error
}
