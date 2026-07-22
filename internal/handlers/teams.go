package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/biplob-codes/mockly/internal/database/sqlc"
	"github.com/biplob-codes/mockly/internal/store"
)

type TeamHandler struct {
	store store.TeamStore
}

type createTeamRequest struct {
	Name     string `json:"name"`
	SenseiID int64  `json:"sensei_id"`
}

type teamDetailsResponse struct {
	ID      int64               `json:"id"`
	Name    string              `json:"name"`
	Sensei  CharacterResponse   `json:"sensei"`
	Members []CharacterResponse `json:"members"`
}

// internal/team/response.go
type TeamResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	SenseiID  int64  `json:"senseiId"`
	CreatedAt string `json:"createdAt"`
}

func toTeamResponse(t sqlc.Team) TeamResponse {
	return TeamResponse{
		ID:        t.ID,
		Name:      t.Name,
		SenseiID:  t.SenseiID,
		CreatedAt: t.CreatedAt.Format(time.RFC3339),
	}
}
func NewTeamHandler(s store.TeamStore) *TeamHandler {
	return &TeamHandler{store: s}
}
func (h *TeamHandler) ListTeams(w http.ResponseWriter, r *http.Request) {
	teams, err := h.store.List(r.Context())
	if err != nil {
		writeRes(w, http.StatusInternalServerError, "Something went wrong")
		return
	}
	var tres []TeamResponse
	for _, t := range teams {
		res := toTeamResponse(t)
		tres = append(tres, res)
	}
	writeRes(w, http.StatusOK, tres)
}
func (h *TeamHandler) CreateTeam(w http.ResponseWriter, r *http.Request) {
	var req createTeamRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	// Do validation on the req

	params := sqlc.CreateTeamParams{
		Name:     req.Name,
		SenseiID: req.SenseiID,
	}

	newTeam, err := h.store.Create(r.Context(), params)
	if err != nil {
		writeRes(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeRes(w, http.StatusCreated, toTeamResponse(newTeam))
}

func (h *TeamHandler) DeleteTeam(w http.ResponseWriter, r *http.Request) {
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
	writeRes(w, http.StatusOK, toTeamResponse(deleted))
}

func (h *TeamHandler) GetTeam(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid number")
		return
	}

	team, err := h.store.Get(r.Context(), int64(n))
	if err != nil {
		writeRes(w, http.StatusNotFound, "Team not found")
		return
	}
	writeRes(w, http.StatusOK, toTeamResponse(team))
}

func (h *TeamHandler) GetTeamDetails(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid number")
		return
	}

	team, err := h.store.Get(r.Context(), int64(n))
	if err != nil {
		writeRes(w, http.StatusNotFound, "Team not found")
		return
	}

	allMembers, err := h.store.GetMembers(r.Context(), team.ID)
	if err != nil {
		writeRes(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	var sensei CharacterResponse
	var members []CharacterResponse
	for _, c := range allMembers {
		if c.ID == team.SenseiID {
			sensei = toCharacterResponse(c)
		} else {
			r := toCharacterResponse(c)
			members = append(members, r)
		}
	}

	writeRes(w, http.StatusOK, teamDetailsResponse{
		ID:      team.ID,
		Name:    team.Name,
		Sensei:  sensei,
		Members: members,
	})
}
