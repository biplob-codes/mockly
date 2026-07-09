package handlers

import (
	"net/http"
	"time"
)

func Home(w http.ResponseWriter, r *http.Request) {
	writeRes(w, http.StatusOK, map[string]any{
		"message":     "Welcome to Mockly — a fake REST API, Konoha edition.",
		"description": "A mock REST API serving Naruto-themed fake data for frontend prototyping, testing, and learning Go.",
		"version":     "1.0.0",
		"repository":  "https://github.com/biplob-codes/mockly",
		"timestamp":   time.Now().UTC().Format(time.RFC3339),
		"resources": map[string]any{
			"users": map[string]any{
				"path":    "/users",
				"methods": []string{"GET"},
				"count":   14,
			},
			"todos": map[string]any{
				"path":    "/todos",
				"methods": []string{"GET"},
				"count":   15,
			},
			"posts": map[string]any{
				"path":    "/posts",
				"methods": []string{"GET"},
				"count":   15,
			},
			"comments": map[string]any{
				"path":    "/comments",
				"methods": []string{"GET"},
				"count":   15,
			},
			"products": map[string]any{
				"path":    "/products",
				"methods": []string{"GET"},
				"count":   15,
			},
		},
		"health": "/health",
	})
}