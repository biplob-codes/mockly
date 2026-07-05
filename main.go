package main

import (
	"fmt"
	"net/http"
)
func main() {
	mux:= http.NewServeMux()
	 mux.HandleFunc("GET /health",func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w,"The server is healthy.")
	 })
	 
	 fmt.Println("Listening on port 8080")
	 http.ListenAndServe(":8080",mux)
}
