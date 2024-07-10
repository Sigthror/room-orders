package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"top-selection-test/internal/logger"
)

var (
	errDefaultHTTP = errors.New("something getting wrong")
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
		resErr := errDefaultHTTP

		var httpErr ResponseError

		// Probably overhead, try to make it simplier in the future
		logErr := err
		if errors.As(err, &httpErr) {
			code = httpErr.Code
			resErr = httpErr.ResponseError
			logErr = httpErr.VerboseErr
			if logErr == nil {
				logErr = httpErr.ResponseError
			}
		}
		logger.FromContext(r.Context()).Debug("%s", logErr)

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
