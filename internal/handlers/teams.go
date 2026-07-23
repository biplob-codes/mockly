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

type TeamHandler struct {
	store          store.TeamStore
	characterStore store.CharacterStore
}

type createTeamRequest struct {
	Name     string `json:"name"`
	SenseiID int64  `json:"senseiId"`
}
type updateTeamRequest struct {
	Name *string `json:"name"`
}
type addTeamMemberRequest struct {
	CharacterID int64 `json:"characterId"`
}

type TeamResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}
type teamMemberResponse struct {
	Character CharacterResponse `json:"character"`
	Role      string            `json:"role"`
	JoinedAt  string            `json:"joinedAt"`
}
type teamDetailsResponse struct {
	ID      int64                `json:"id"`
	Name    string               `json:"name"`
	Members []teamMemberResponse `json:"members"`
}

func toTeamResponse(t sqlc.Team) TeamResponse {
	return TeamResponse{
		ID:        t.ID,
		Name:      t.Name,
		CreatedAt: t.CreatedAt.Format(time.RFC3339),
	}
}

func NewTeamHandler(s store.TeamStore, cs store.CharacterStore) *TeamHandler {
	return &TeamHandler{store: s, characterStore: cs}
}

func (h *TeamHandler) ListTeams(w http.ResponseWriter, r *http.Request) {
	teams, err := h.store.List(r.Context())
	if err != nil {
		writeRes(w, http.StatusInternalServerError, "Something went wrong")
		return
	}
	var tres []TeamResponse
	for _, t := range teams {
		tres = append(tres, toTeamResponse(t))
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

	newTeam, err := h.store.Create(r.Context(), req.Name, req.SenseiID)
	if err != nil {
		writeRes(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeRes(w, http.StatusCreated, toTeamResponse(newTeam))
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

func (h *TeamHandler) DeleteTeam(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid team ID")
		return
	}

	deleted, err := h.store.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeRes(w, http.StatusNotFound, "Team not found")
			return
		}
		writeRes(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeRes(w, http.StatusOK, toTeamResponse(deleted))
}

func (h *TeamHandler) GetTeamDetails(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid team ID")
		return
	}

	rows, err := h.store.GetDetails(r.Context(), id)
	if err != nil {
		writeRes(w, http.StatusInternalServerError, "Something went wrong")
		return
	}
	if len(rows) == 0 {
		writeRes(w, http.StatusNotFound, "Team not found")
		return
	}

	resp := teamDetailsResponse{ID: rows[0].TeamID, Name: rows[0].TeamName}
	for _, row := range rows {
		if !row.ID.Valid { // no members yet
			continue
		}
		resp.Members = append(resp.Members, teamMemberResponse{
			Character: toCharacterResponse(sqlc.Character{
				ID:        row.ID.Int64,
				Name:      row.Name.String,
				Nickname:  row.Nickname,
				Clan:      row.Clan,
				Age:       row.Age,
				Rank:      row.Rank.String,
				Birthdate: row.Birthdate.Time,
				VillageID: row.VillageID.Int64,
				CreatedAt: row.CreatedAt.Time,
				UpdatedAt: row.UpdatedAt.Time,
			}),
			Role:     row.Role.String,
			JoinedAt: row.JoinedAt.Time.Format(time.RFC3339),
		})
	}
	writeRes(w, http.StatusOK, resp)
}

func (h *TeamHandler) UpdateTeam(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid team ID")
		return
	}

	var req updateTeamRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	// Do validation on the req

	teamParams := sqlc.UpdateTeamParams{ID: id, Name: *req.Name}

	updatedTeam, err := h.store.Update(r.Context(), teamParams)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeRes(w, http.StatusNotFound, "Team not found")
			return
		}
		writeRes(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeRes(w, http.StatusOK, toTeamResponse(updatedTeam))
}

func (h *TeamHandler) AddTeamMember(w http.ResponseWriter, r *http.Request) {
	teamIDStr := r.PathValue("id")
	teamID, err := strconv.ParseInt(teamIDStr, 10, 64)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid team ID")
		return
	}

	var req addTeamMemberRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	// Do validation on the req (role must be "sensei" or "student")

	character, err := h.characterStore.Get(r.Context(), req.CharacterID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeRes(w, http.StatusNotFound, "Character not found")
			return
		}
		writeRes(w, http.StatusInternalServerError, err.Error())
		return
	}

	existing, err := h.store.GetCharacterMembership(r.Context(), req.CharacterID)
	if err == nil {
		if existing.TeamID == teamID {
			writeRes(w, http.StatusConflict, "Character is already a member of this team")
			return
		}
		writeRes(w, http.StatusConflict, map[string]any{
			"message":   "Character is already a member of a different team",
			"character": toCharacterResponse(character),
		})
		return
	} else if !errors.Is(err, sql.ErrNoRows) {
		writeRes(w, http.StatusInternalServerError, err.Error())
		return
	}

	member, err := h.store.AddMember(r.Context(), sqlc.AddTeamMemberParams{
		TeamID:      teamID,
		CharacterID: req.CharacterID,
		Role:        "Student",
	})
	if err != nil {
		writeRes(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeRes(w, http.StatusCreated, teamMemberResponse{
		Character: toCharacterResponse(character),
		Role:      member.Role,
		JoinedAt:  member.JoinedAt.Format(time.RFC3339),
	})
}

func (h *TeamHandler) RemoveTeamMember(w http.ResponseWriter, r *http.Request) {
	teamIDStr := r.PathValue("id")
	teamID, err := strconv.ParseInt(teamIDStr, 10, 64)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid team ID")
		return
	}

	characterIDStr := r.PathValue("characterId")
	characterID, err := strconv.ParseInt(characterIDStr, 10, 64)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid character ID")
		return
	}

	removed, err := h.store.RemoveMember(r.Context(), sqlc.RemoveTeamMemberParams{
		TeamID:      teamID,
		CharacterID: characterID,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeRes(w, http.StatusNotFound, "Character is not a member of this team")
			return
		}
		writeRes(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeRes(w, http.StatusOK, removed)
}
