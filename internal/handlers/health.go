package handlers

import "net/http"

func Health(w http.ResponseWriter, r *http.Request) {
	writeRes(w, http.StatusOK, map[string]any{
		"status": "healthy",
	})
}