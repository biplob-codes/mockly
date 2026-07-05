package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
 

	"github.com/biplob-codes/mockly/internal/model"
	"github.com/biplob-codes/mockly/internal/store"
)

func findTodoById(id int) (model.Todo, bool) {
	for _, t := range store.Todos {
		if t.Id == id {
			return t, true
		}
	}
	return model.Todo{}, false
}

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	writeRes(w, http.StatusOK, store.Todos)
}

func GetTodoById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid number")
		return
	}
	todo, found := findTodoById(n)
	if found {
		writeRes(w, http.StatusOK, todo)
		return
	}
	writeRes(w, http.StatusNotFound, "Todo with the Id not found")
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var req model.NewTodo
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	if req.Title == "" {
		writeRes(w, http.StatusBadRequest, "Title is required")
		return
	}
	createdTodo := model.Todo{
		Id:        len(store.Todos) + 1,
		UserId:    req.UserId,
		Title:     req.Title,
		Completed: req.Completed,
		Priority:  req.Priority,
		DueDate:   req.DueDate,
 	}
	writeRes(w, http.StatusCreated, createdTodo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid number")
		return
	}
	var req model.NewTodo
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	updatedTodo := model.Todo{
		Id:        n,
		UserId:    req.UserId,
		Title:     req.Title,
		Completed: req.Completed,
		Priority:  req.Priority,
		DueDate:   req.DueDate,
	}
	writeRes(w, http.StatusOK, updatedTodo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid number")
		return
	}
	message := fmt.Sprintf("Todo with the Id: %d deleted", n)
	writeRes(w, http.StatusOK, message)
}