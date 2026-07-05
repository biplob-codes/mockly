package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/biplob-codes/mockly/internal/model"
	"github.com/biplob-codes/mockly/internal/store"
)

func findCommentById(id int) (model.Comment, bool) {
	for _, c := range store.Comments {
		if c.Id == id {
			return c, true
		}
	}
	return model.Comment{}, false
}

func GetAllComments(w http.ResponseWriter, r *http.Request) {
	writeRes(w, http.StatusOK, store.Comments)
}

func GetCommentById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid number")
		return
	}
	comment, found := findCommentById(n)
	if found {
		writeRes(w, http.StatusOK, comment)
		return
	}
	writeRes(w, http.StatusNotFound, "Comment with the Id not found")
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	var req model.NewComment
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	if req.Body == "" {
		writeRes(w, http.StatusBadRequest, "Body is required")
		return
	}
	createdComment := model.Comment{
		Id:     len(store.Comments) + 1,
		PostId: req.PostId,
		Name:   req.Name,
		Email:  req.Email,
		Body:   req.Body,
		CreatedAt: time.Now().UTC().Format(time.RFC3339),
	}
	writeRes(w, http.StatusCreated, createdComment)
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid number")
		return
	}
	var req model.NewComment
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	updatedComment := model.Comment{
		Id:     n,
		PostId: req.PostId,
		Name:   req.Name,
		Email:  req.Email,
		Body:   req.Body,
	}
	writeRes(w, http.StatusOK, updatedComment)
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid number")
		return
	}
	message := fmt.Sprintf("Comment with the Id: %d deleted", n)
	writeRes(w, http.StatusOK, message)
}