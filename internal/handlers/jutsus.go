package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/biplob-codes/mockly/internal/database/sqlc"
	"github.com/biplob-codes/mockly/internal/store"
)

type JutsuHandler struct {
	store store.JutsuStore
}
type createJutsuRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Rank        string `json:"rank"`
}

func NewJutsuHandler(s store.JutsuStore) *JutsuHandler {
	return &JutsuHandler{store: s}
}

func (h *JutsuHandler) ListJutsus(w http.ResponseWriter, r *http.Request) {
	jutsus, err := h.store.List(r.Context())
	if err != nil {
		writeRes(w, http.StatusInternalServerError, "Something went wrong")
		return
	}
	writeRes(w, http.StatusOK, jutsus)
}

func (h *JutsuHandler) GetJutsu(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid number")
		return
	}
	jutsu, err := h.store.Get(r.Context(), int64(n))
	if err != nil {
		writeRes(w, http.StatusNotFound, err.Error())
		return
	}
	writeRes(w, http.StatusOK, jutsu)
}

func (h *JutsuHandler) CreateJutsu(w http.ResponseWriter, r *http.Request) {
	var req createJutsuRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	//Do validation on the req

	jutsuParams := sqlc.CreateJutsuParams{
		Name:        req.Name,
		Description: sql.NullString{String: req.Description, Valid: req.Description != ""},
		Type:        req.Type,
		Rank:        req.Rank,
	}

	newJutsu, err := h.store.Create(r.Context(), jutsuParams)
	if err != nil {
		writeRes(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeRes(w, http.StatusCreated, newJutsu)
}

func (h *JutsuHandler) DeleteJutsu(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid number")
		return
	}
	jutsu, err := h.store.Delete(r.Context(), int64(n))
	if err != nil {
		writeRes(w, http.StatusInternalServerError, err)
		return
	}
	writeRes(w, http.StatusOK, jutsu)
}