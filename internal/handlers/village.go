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

 
type VillageHandler struct{
	store store.VillageStore
}
type createVillageRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Land        string `json:"land"`
	Population  int64  `json:"population"`
	KageID      *int64 `json:"kage_id"` 
	FoundedAt   string `json:"founded_at"`
}
func NewVillageHandler(s store.VillageStore) *VillageHandler{
	return  &VillageHandler{store:s }
}

func (h *VillageHandler) ListVillages(w http.ResponseWriter, r *http.Request) {
	villages,err:=h.store.List()

	if err!=nil{
		writeRes(w,http.StatusInternalServerError,"Something went wrong")
	}
	writeRes(w, http.StatusOK, villages)
}

func (h *VillageHandler) GetVillage(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid number")
		return
	}
	village,err:=h.store.Get(int64(n))

	if err!=nil{
     writeRes(w,http.StatusNotFound,err.Error())
	 return
	}

	writeRes(w, http.StatusOK, village)
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
	newVillage,err:=h.store.Create(villageParams)

	if err!=nil{
		writeRes(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeRes(w, http.StatusCreated, newVillage)
}


func (h *VillageHandler) DeleteVillage(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid number")
		return
	}
	village,err:=h.store.Delete(int64(n))
	if err!=nil{
		writeRes(w, http.StatusInternalServerError, err)
	}
	writeRes(w, http.StatusOK, village)
}
