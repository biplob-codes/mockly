package main

import (
	"fmt"
	"net/http"

	"github.com/biplob-codes/mockly/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health",handlers.Health)
	mux.HandleFunc("GET /users", handlers.GetAllUsers)
	mux.HandleFunc("GET /users/{id}", handlers.GetUserById)
	mux.HandleFunc("POST /users", handlers.CreateUser)
	mux.HandleFunc("PUT /users/{id}", handlers.UpdateUser)
	mux.HandleFunc("DELETE /users/{id}", handlers.DeleteUser)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", mux)
}
