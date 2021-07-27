package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseWithError(w http.ResponseWriter, statusCode int, msg string) {
	ResponseWithJson(w, statusCode, map[string]string{"error": msg})
}
func ResponseWithJson(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "Application-Json")
	w.WriteHeader(statusCode)
	w.Write(response)
}
