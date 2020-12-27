package response

import (
	"app/api/llog"
	"encoding/json"
	"net/http"
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

func InternalServerError(w http.ResponseWriter, err error, message string) {
	httpError(w, http.StatusInternalServerError, err, message)
}

func httpError(w http.ResponseWriter, statusCode int, err error, message string) {
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
