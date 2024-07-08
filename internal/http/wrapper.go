package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	defaultHTTPError = errors.New("Something getting wrong. Please visit ")
)

type ResponseError struct {
	Code          int
	VerboseErr    error
	ResponseError error
}

func (e ResponseError) Error() string {
	if e.VerboseErr != nil {
		return e.VerboseErr.Error()
	}

	return e.Error()
}

func newEndpoint(endpoint func(w http.ResponseWriter, r *http.Request) error, defaultHTTPCode ...int) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		if err = endpoint(w, r); err == nil {
			if defaultHTTPCode != nil {
				w.WriteHeader(defaultHTTPCode[0])
			}
			return
		}

		code := http.StatusInternalServerError
		resErr := defaultHTTPError

		var httpErr ResponseError
		if errors.As(err, &httpErr) {
			code = httpErr.Code
			resErr = httpErr.ResponseError
			if httpErr.VerboseErr != nil {
				// TODO Make logger call
				fmt.Println(httpErr.VerboseErr.Error())
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)

		errResponse := struct {
			ErrorInfo string
		}{
			ErrorInfo: resErr.Error(),
		}
		json.NewEncoder(w).Encode(errResponse)
	}
}
