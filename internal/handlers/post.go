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

func findPostById(id int) (model.Post, bool) {
	for _, p := range store.Posts {
		if p.Id == id {
			return p, true
		}
	}
	return model.Post{}, false
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	writeRes(w, http.StatusOK, store.Posts)
}

func GetPostById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid number")
		return
	}
	post, found := findPostById(n)
	if found {
		writeRes(w, http.StatusOK, post)
		return
	}
	writeRes(w, http.StatusNotFound, "Post with the Id not found")
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var req model.NewPost
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	if req.Title == "" {
		writeRes(w, http.StatusBadRequest, "Title is required")
		return
	}
	createdPost := model.Post{
		Id:     len(store.Posts) + 1,
		UserId: req.UserId,
		Title:  req.Title,
		Body:   req.Body,
		Tags:   req.Tags,
		Views:  0,
		CreatedAt: time.Now().UTC().Format(time.RFC3339),
	}
	writeRes(w, http.StatusCreated, createdPost)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid number")
		return
	}
	var req model.NewPost
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	existing, found := findPostById(n)
	views := 0
	if found {
		views = existing.Views
	}
	updatedPost := model.Post{
		Id:     n,
		UserId: req.UserId,
		Title:  req.Title,
		Body:   req.Body,
		Tags:   req.Tags,
		Views:  views,
	}
	writeRes(w, http.StatusOK, updatedPost)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid number")
		return
	}
	message := fmt.Sprintf("Post with the Id: %d deleted", n)
	writeRes(w, http.StatusOK, message)
}