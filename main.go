package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/biplob-codes/mockly/internal/store"
)

type NewUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func writeRes(w http.ResponseWriter, statusCode int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}
func findUserById(id int) (store.User, bool) {
	for _, u := range store.Users {
		if u.Id == id {
			return u, true
		}
	}
	return store.User{}, false
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "The server is healthy.")
	})

	mux.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		writeRes(w, http.StatusOK, store.Users)
	})
	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		n, err := strconv.Atoi(id)
		if err != nil {
			writeRes(w, http.StatusBadRequest, "Invalid number")
			return
		}
		user, found := findUserById(n)
		if found {
			writeRes(w, http.StatusOK, user)
			return
		}
		writeRes(w, http.StatusNotFound, "User with the Id not found")

	})
	mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		var req NewUser
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeRes(w, http.StatusBadRequest, "Invalid JSON")
			return
		}
		if req.Name ==""{
			writeRes(w, http.StatusBadRequest, "Name is required")
			return
		}
		if req.Email ==""{
			writeRes(w, http.StatusBadRequest, "Email is required")
			return
		}
		createdUser := store.User{
			Id:    11,
			Name:  req.Name,
			Email: req.Email,
		}
		writeRes(w, http.StatusCreated, createdUser)
	})
	mux.HandleFunc("PUT /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		n, err := strconv.Atoi(id)
		if err != nil {
			writeRes(w, http.StatusBadRequest, "Invalid number")
			return
		}
		var req NewUser
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeRes(w, http.StatusBadRequest, "Invalid JSON")
			return
		}
		updatedUser := store.User{
			Id:    n,
			Name:  req.Name,
			Email: req.Email,
		}
		writeRes(w, http.StatusOK, updatedUser)
	})
	mux.HandleFunc("DELETE /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		n, err := strconv.Atoi(id)
		if err != nil {
			writeRes(w, http.StatusBadRequest, "Invalid number")
			return
		}
		message := fmt.Sprintf("User with the Id: %d deleted", n)
		writeRes(w, http.StatusOK, message)

	})
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", mux)
}
