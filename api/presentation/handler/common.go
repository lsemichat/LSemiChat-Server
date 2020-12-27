package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

// ReadRequestBody req = pointer json unmarshalをやってくれる
func ReadRequestBody(r *http.Request, req interface{}) (interface{}, error) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read request body")
	}
	err = json.Unmarshal(reqBody, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal to json")
	}
	return req, nil
}

func ReadPathParam(r *http.Request, key string) string {
	return mux.Vars(r)[key]

}
