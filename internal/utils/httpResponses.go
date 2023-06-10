package utils

import (
	"encoding/json"
	"net/http"
)

func SendHttpResponseError(w http.ResponseWriter, err error, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(
		map[string]string{
			"status": "false",
			"error":  err.Error(),
		})
}

func SendHttpResponse(w http.ResponseWriter, data interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
