package web

import (
	"encoding/json"
	"net/http"
)

type RequestError struct {
	Message string `json: message`
}

// https://golang.org/src/net/http/status.go
func JSONError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	err := RequestError{message}

	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}
