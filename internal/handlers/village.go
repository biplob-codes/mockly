package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/biplob-codes/mockly/internal/database/sqlc"
	"github.com/biplob-codes/mockly/internal/store"
)

type VillageHandler struct {
	store store.VillageStore
}
type createVillageRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Land        string `json:"land"`
	Population  int64  `json:"population"`
	KageID      *int64 `json:"kageId"`
	FoundedAt   string `json:"foundedAt"`
}
type VillageResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Land        string `json:"land"`
	Population  *int64 `json:"population,omitempty"`
	KageID      *int64 `json:"kageId,omitempty"`
	FoundedAt   string `json:"foundedAt"`
}
type updateVillageRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Land        *string `json:"land"`
	Population  *int64  `json:"population"`
	KageID      *int64  `json:"kageId"`
	FoundedAt   *string `json:"foundedAt"`
}

func toVillageResponse(v sqlc.Village) VillageResponse {
	resp := VillageResponse{
		ID:        v.ID,
		Name:      v.Name,
		Land:      v.Land,
		FoundedAt: v.FoundedAt.Format("2006-01-02"),
	}
	if v.Description.Valid {
		resp.Description = v.Description.String
	}
	if v.Population.Valid {
		resp.Population = &v.Population.Int64
	}
	if v.KageID.Valid {
		resp.KageID = &v.KageID.Int64
	}
	return resp
}
func NewVillageHandler(s store.VillageStore) *VillageHandler {
	return &VillageHandler{store: s}
}

func (h *VillageHandler) ListVillages(w http.ResponseWriter, r *http.Request) {
	villages, err := h.store.List(r.Context())

	if err != nil {
		writeRes(w, http.StatusInternalServerError, "Something went wrong")
	}
	var vres []VillageResponse
	for _, vl := range villages {
		res := toVillageResponse(vl)
		vres = append(vres, res)
	}
	writeRes(w, http.StatusOK, vres)
}

func (h *VillageHandler) GetVillage(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid number")
		return
	}
	village, err := h.store.Get(r.Context(), int64(n))

	if err != nil {
		writeRes(w, http.StatusNotFound, err.Error())
		return
	}

	writeRes(w, http.StatusOK, toVillageResponse(village))
}

func (h *VillageHandler) CreateVillage(w http.ResponseWriter, r *http.Request) {
	var req createVillageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	foundedAt, err := time.Parse(time.RFC3339, req.FoundedAt)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid date")
		return
	}

	//Do validation on the req

	villageParams := sqlc.CreateVillageParams{
		Name:        req.Name,
		Description: sql.NullString{String: req.Description, Valid: req.Description != ""},
		Land:        req.Land,
		Population:  sql.NullInt64{Int64: req.Population, Valid: req.Population != 0},
		FoundedAt:   foundedAt,
	}

	if req.KageID != nil {
		villageParams.KageID = sql.NullInt64{Int64: *req.KageID, Valid: true}
	}
	newVillage, err := h.store.Create(r.Context(), villageParams)

	if err != nil {
		writeRes(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeRes(w, http.StatusCreated, toVillageResponse(newVillage))
}

func (h *VillageHandler) DeleteVillage(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid number")
		return
	}
	village, err := h.store.Delete(r.Context(), int64(n))
	if err != nil {
		writeRes(w, http.StatusInternalServerError, err)
	}
	writeRes(w, http.StatusOK, toVillageResponse(village))
}

func (h *VillageHandler) UpdateVillage(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid village ID")
		return
	}

	var req updateVillageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	// Do validation on the req

	villageParams := sqlc.UpdateVillageParams{
		ID: id,
	}

	if req.Name != nil {
		villageParams.Name = sql.NullString{String: *req.Name, Valid: true}
	}
	if req.Description != nil {
		villageParams.Description = sql.NullString{String: *req.Description, Valid: true}
	}
	if req.Land != nil {
		villageParams.Land = sql.NullString{String: *req.Land, Valid: true}
	}
	if req.Population != nil {
		villageParams.Population = sql.NullInt64{Int64: *req.Population, Valid: true}
	}
	if req.KageID != nil {
		villageParams.KageID = sql.NullInt64{Int64: *req.KageID, Valid: true}
	}
	if req.FoundedAt != nil {
		foundedAt, err := time.Parse(time.RFC3339, *req.FoundedAt)
		if err != nil {
			writeRes(w, http.StatusBadRequest, "Invalid date")
			return
		}
		villageParams.FoundedAt = sql.NullTime{Time: foundedAt, Valid: true}
	}

	updatedVillage, err := h.store.Update(r.Context(), villageParams)
	if err != nil {
		writeRes(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeRes(w, http.StatusOK, toVillageResponse(updatedVillage))
}
