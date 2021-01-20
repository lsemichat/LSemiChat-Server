package response

import (
	"app/api/llog"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

func Success(w http.ResponseWriter, v interface{}) {
	jsonData, err := json.Marshal(v)
	if err != nil {
		InternalServerError(w, err, "failed to marshal to json")
	}
	w.Write(jsonData)
}

func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func BadRequest(w http.ResponseWriter, err error, message string) {
	httpError(w, http.StatusBadRequest, err, message)
}

func Unauthorized(w http.ResponseWriter, err error, message string) {
	httpError(w, http.StatusUnauthorized, err, message)
}

func NotFound(w http.ResponseWriter, err error, message string) {
	httpError(w, http.StatusNotFound, err, message)
}

func InternalServerError(w http.ResponseWriter, err error, message string) {
	httpError(w, http.StatusInternalServerError, err, message)
}

func NotImplemented(w http.ResponseWriter) {
	err := errors.New("not implemented")
	httpError(w, http.StatusNotImplemented, err, err.Error())
}

func httpError(w http.ResponseWriter, statusCode int, err error, message string) {
	// TODO: logをmiddlewareで挿入したい
	llog.Error(err.Error())
	res := &httpErrorResponse{
		StatusCode: statusCode,
		Message:    message,
	}
	jsonData, _ := json.Marshal(res)
	w.WriteHeader(statusCode)
	w.Write(jsonData)
}

type httpErrorResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}
