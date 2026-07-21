package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/biplob-codes/mockly/internal/database/sqlc"
	"github.com/biplob-codes/mockly/internal/store"
)

type MissionHandler struct {
	store store.MissionStore
}

type createMissionRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	AssignedTo  int64  `json:"assigned_to"`
	Rank        string `json:"rank"`
	Status      string `json:"status"`
	Reward      int64  `json:"reward"`
	StartsAt    string `json:"starts_at"`
}

func NewMissionHandler(s store.MissionStore) *MissionHandler {
	return &MissionHandler{store: s}
}

func (h *MissionHandler) CreateMission(w http.ResponseWriter, r *http.Request) {
	var req createMissionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	startsAt, err := time.Parse(time.RFC3339, req.StartsAt)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid date")
		return
	}

	// Do validation on the req

	params := sqlc.CreateMissionParams{
		Name:        req.Name,
		Description: req.Description,
		AssignedTo:  req.AssignedTo,
		Rank:        req.Rank,
		Status:      req.Status,
		Reward:      req.Reward,
		StartsAt:    startsAt,
	}

	newMission, err := h.store.Create(r.Context(), params)
	if err != nil {
		writeRes(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeRes(w, http.StatusCreated, newMission)
}

func (h *MissionHandler) GetMission(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid number")
		return
	}

	mission, err := h.store.Get(r.Context(), int64(n))
	if err != nil {
		writeRes(w, http.StatusNotFound, "Mission not found")
		return
	}
	writeRes(w, http.StatusOK, mission)
}

func (h *MissionHandler) ListMissions(w http.ResponseWriter, r *http.Request) {
	missions, err := h.store.List(r.Context())
	if err != nil {
		writeRes(w, http.StatusInternalServerError, "Something went wrong")
		return
	}
	writeRes(w, http.StatusOK, missions)
}

func (h *MissionHandler) GetMissionsByTeam(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("teamId")
	n, err := strconv.Atoi(id)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid number")
		return
	}

	missions, err := h.store.GetByTeam(r.Context(), int64(n))
	if err != nil {
		writeRes(w, http.StatusInternalServerError, "Something went wrong")
		return
	}
	writeRes(w, http.StatusOK, missions)
}

func (h *MissionHandler) DeleteMission(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid number")
		return
	}

	deleted, err := h.store.Delete(r.Context(), int64(n))
	if err != nil {
		writeRes(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeRes(w, http.StatusOK, deleted)
}