package utils

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type errorMessage struct {
	Code    int    `json:"code"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

type appError struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	HttpStatus int    `json:"status"`
}

type errorResource struct {
	Data errorMessage `json:"status"`
}

//Display error on err != nil
func DisplayAppError(w http.ResponseWriter, handlerError error, message string, code int) {

	errObj := errorMessage{
		Error:   handlerError.Error(),
		Code:    code,
		Message: message,
	}

	log.Error("Error processing request")

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.WriteHeader(code)

	if j, err := json.Marshal(errorResource{Data: errObj}); err == nil {
		w.Write(j)
	}

	return

}
