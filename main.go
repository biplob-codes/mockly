package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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
    port:=os.Getenv("PORT")
	if port==""{
		port="8080"
	}
	addr:=":"+port
	srv:=&http.Server{
		Addr: addr,
		Handler: mux,
		ReadHeaderTimeout: 5*time.Second,
		ReadTimeout: 10*time.Second,
		WriteTimeout: 10*time.Second,
	}
	fmt.Println("Listening on port",port)
	if err:=srv.ListenAndServe();err!=nil{
		log.Fatal(err)
	}

	
}
