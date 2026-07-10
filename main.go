package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/biplob-codes/mockly/internal/database"
	"github.com/biplob-codes/mockly/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handlers.Home)
	mux.HandleFunc("GET /health", handlers.Health)

	mux.HandleFunc("GET /users", handlers.GetAllUsers)
	mux.HandleFunc("GET /users/{id}", handlers.GetUserById)
	mux.HandleFunc("POST /users", handlers.CreateUser)
	mux.HandleFunc("PUT /users/{id}", handlers.UpdateUser)
	mux.HandleFunc("DELETE /users/{id}", handlers.DeleteUser)

	mux.HandleFunc("GET /posts", handlers.GetAllPosts)
	mux.HandleFunc("GET /posts/{id}", handlers.GetPostById)
	mux.HandleFunc("POST /posts", handlers.CreatePost)
	mux.HandleFunc("PUT /posts/{id}", handlers.UpdatePost)
	mux.HandleFunc("DELETE /posts/{id}", handlers.DeletePost)

	mux.HandleFunc("GET /todos", handlers.GetAllTodos)
	mux.HandleFunc("GET /todos/{id}", handlers.GetTodoById)
	mux.HandleFunc("POST /todos", handlers.CreateTodo)
	mux.HandleFunc("PUT /todos/{id}", handlers.UpdateTodo)
	mux.HandleFunc("DELETE /todos/{id}", handlers.DeleteTodo)

	mux.HandleFunc("GET /products", handlers.GetAllProducts)
	mux.HandleFunc("GET /products/{id}", handlers.GetProductById)
	mux.HandleFunc("POST /products", handlers.CreateProduct)
	mux.HandleFunc("PUT /products/{id}", handlers.UpdateProduct)
	mux.HandleFunc("DELETE /products/{id}", handlers.DeleteProduct)

	mux.HandleFunc("GET /comments", handlers.GetAllComments)
	mux.HandleFunc("GET /comments/{id}", handlers.GetCommentById)
	mux.HandleFunc("POST /comments", handlers.CreateComment)
	mux.HandleFunc("PUT /comments/{id}", handlers.UpdateComment)
	mux.HandleFunc("DELETE /comments/{id}", handlers.DeleteComment)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port
	srv := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}
	db,err:=database.Connect()
	if err!=nil{
		fmt.Println(err)
	}
	defer db.Close()
	fmt.Println("Database connection successfull")
	fmt.Println("Listening on port", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
