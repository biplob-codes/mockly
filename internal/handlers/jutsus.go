package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

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

type updateJutsuRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Type        *string `json:"type"`
	Rank        *string `json:"rank"`
}

type JutsuResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type"`
	Rank        string `json:"rank"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func toJutsuResponse(j sqlc.Jutsu) JutsuResponse {
	resp := JutsuResponse{
		ID:        j.ID,
		Name:      j.Name,
		Type:      j.Type,
		Rank:      j.Rank,
		CreatedAt: j.CreatedAt.Format(time.RFC3339),
		UpdatedAt: j.UpdatedAt.Format(time.RFC3339),
	}
	if j.Description.Valid {
		resp.Description = j.Description.String
	}
	return resp
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
	var jres []JutsuResponse
	for _, j := range jutsus {
		res := toJutsuResponse(j)
		jres = append(jres, res)
	}
	writeRes(w, http.StatusOK, jres)
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
	writeRes(w, http.StatusOK, toJutsuResponse(jutsu))
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
	writeRes(w, http.StatusCreated, toJutsuResponse(newJutsu))
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
	writeRes(w, http.StatusOK, toJutsuResponse(jutsu))
}

func (h *JutsuHandler) UpdateJutsu(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid jutsu ID")
		return
	}

	var req updateJutsuRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	// Do validation on the req

	jutsuParams := sqlc.UpdateJutsuParams{
		ID: id,
	}

	if req.Name != nil {
		jutsuParams.Name = sql.NullString{String: *req.Name, Valid: true}
	}
	if req.Description != nil {
		jutsuParams.Description = sql.NullString{String: *req.Description, Valid: true}
	}
	if req.Type != nil {
		jutsuParams.Type = sql.NullString{String: *req.Type, Valid: true}
	}
	if req.Rank != nil {
		jutsuParams.Rank = sql.NullString{String: *req.Rank, Valid: true}
	}

	updatedJutsu, err := h.store.Update(r.Context(), jutsuParams)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeRes(w, http.StatusNotFound, "Jutsu not found")
			return
		}
		writeRes(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeRes(w, http.StatusOK, toJutsuResponse(updatedJutsu))
}
