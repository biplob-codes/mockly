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

func findUserById(id int) (model.User, bool) {
	for _, u := range store.Users {
		if u.Id == id {
			return u, true
		}
	}
	return model.User{}, false
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	writeRes(w, http.StatusOK, store.Users)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
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
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var req model.NewUser
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	if req.Name == "" {
		writeRes(w, http.StatusBadRequest, "Name is required")
		return
	}
	if req.Email == "" {
		writeRes(w, http.StatusBadRequest, "Email is required")
		return
	}
	createdUser := model.User{
		Id:       len(store.Users) + 1,
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
		Website:  req.Website,
		Avatar:   req.Avatar,
		Address:  req.Address,
		Company:  req.Company,
		CreatedAt: time.Now().UTC().Format(time.RFC3339),
	}
	writeRes(w, http.StatusCreated, createdUser)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid number")
		return
	}
	var req model.NewUser
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	updatedUser := model.User{
		Id:       n,
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
		Website:  req.Website,
		Avatar:   req.Avatar,
		Address:  req.Address,
		Company:  req.Company,
	}
	writeRes(w, http.StatusOK, updatedUser)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid number")
		return
	}
	message := fmt.Sprintf("User with the Id: %d deleted", n)
	writeRes(w, http.StatusOK, message)
}