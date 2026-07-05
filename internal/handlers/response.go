package handlers

import (
	"encoding/json"
	"net/http"
)

func writeRes(w http.ResponseWriter, statusCode int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}
