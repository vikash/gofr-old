package http

import (
	"encoding/json"
	"net/http"

	"github.com/zopsmart/ezgo/pkg/gofr/exists"

	"github.com/zopsmart/ezgo/pkg/gofr/errors"
	resTypes "github.com/zopsmart/ezgo/pkg/gofr/http/response"
)

func NewResponder(w http.ResponseWriter) *Responder {
	return &Responder{w: w}
}

type Responder struct {
	w http.ResponseWriter
}

func (r Responder) Respond(data interface{}, err error) {
	r.w.Header().Set("Content-type", "application/json")

	statusCode, errorObj := r.HTTPStatusFromError(err)
	r.w.WriteHeader(statusCode)

	var resp interface{}
	switch v := data.(type) {
	case resTypes.Raw:
		resp = v.Data
	default:
		resp = response{
			Data:  v,
			Error: errorObj,
		}
	}

	_ = json.NewEncoder(r.w).Encode(resp)
}

func (r Responder) HTTPStatusFromError(err error) (int, interface{}) {

	var statusCode int
	message := make(map[string]interface{})

	switch v := err.(type) {
	case nil:
		return http.StatusOK, nil

	case exists.DuplicateEntity:
		return http.StatusOK, nil

	case exists.Entity:
		return http.StatusCreated, nil

	case errors.InvalidParam:
		statusCode = http.StatusBadRequest

	case errors.MissingParam:
		statusCode = http.StatusBadRequest

	case errors.Response:
		statusCode = v.Code

	default:
		statusCode = http.StatusInternalServerError
	}

	message["message"] = err.Error()

	return statusCode, message

}

type response struct {
	Error interface{} `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}
