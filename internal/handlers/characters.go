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

type CharacterHandler struct {
	store store.CharacterStore
}
type CharacterResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Nickname  string `json:"nickname,omitempty"`
	Clan      string `json:"clan,omitempty"`
	Age       *int64 `json:"age,omitempty"`
	Rank      string `json:"rank"`
	Birthdate string `json:"birthdate"`
	VillageID int64  `json:"villageId"`
}
type createCharacterRequest struct {
	Name      string `json:"name"`
	Nickname  string `json:"nickname"`
	Clan      string `json:"clan"`
	Age       int64  `json:"age"`
	Rank      string `json:"rank"`
	Birthdate string `json:"birthdate"`
	VillageID int64  `json:"villageId"`
}
type updateCharacterRequest struct {
	Name      *string `json:"name"`
	Nickname  *string `json:"nickname"`
	Clan      *string `json:"clan"`
	Age       *int64  `json:"age"`
	Rank      *string `json:"rank"`
	Birthdate *string `json:"birthdate"`
	VillageID *int64  `json:"villageId"`
}

func toCharacterResponse(c sqlc.Character) CharacterResponse {
	resp := CharacterResponse{
		ID:        c.ID,
		Name:      c.Name,
		Rank:      c.Rank,
		Birthdate: c.Birthdate.Format("2006-01-02"),
		VillageID: c.VillageID,
	}
	if c.Nickname.Valid {
		resp.Nickname = c.Nickname.String
	}
	if c.Clan.Valid {
		resp.Clan = c.Clan.String
	}
	if c.Age.Valid {
		resp.Age = &c.Age.Int64
	}
	return resp
}
func NewCharacterHandler(s store.CharacterStore) *CharacterHandler {
	return &CharacterHandler{store: s}
}

func (h *CharacterHandler) ListCharacters(w http.ResponseWriter, r *http.Request) {
	characters, err := h.store.List(r.Context())

	if err != nil {
		writeRes(w, http.StatusInternalServerError, "Something went wrong")
		return
	}
	var chres []CharacterResponse

	for _, ch := range characters {
		res := toCharacterResponse(ch)
		chres = append(chres, res)
	}
	writeRes(w, http.StatusOK, chres)
}

func (h *CharacterHandler) GetCharacter(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid number")
		return
	}
	character, err := h.store.Get(r.Context(), int64(n))

	if err != nil {
		writeRes(w, http.StatusNotFound, err.Error())
		return
	}
	writeRes(w, http.StatusOK, toCharacterResponse(character))
}

func (h *CharacterHandler) CreateCharacter(w http.ResponseWriter, r *http.Request) {
	var req createCharacterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	birthdate, err := time.Parse("2006-01-02", req.Birthdate)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid date")
		return
	}

	//Do validation on the req

	characterParams := sqlc.CreateCharacterParams{
		Name:      req.Name,
		Nickname:  sql.NullString{String: req.Nickname, Valid: req.Nickname != ""},
		Clan:      sql.NullString{String: req.Clan, Valid: req.Clan != ""},
		Age:       sql.NullInt64{Int64: req.Age, Valid: req.Age != 0},
		Rank:      req.Rank,
		Birthdate: birthdate,
		VillageID: req.VillageID,
	}

	newCharacter, err := h.store.Create(r.Context(), characterParams)

	if err != nil {
		writeRes(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeRes(w, http.StatusCreated, toCharacterResponse(newCharacter))
}

func (h *CharacterHandler) DeleteCharacter(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid number")
		return
	}
	character, err := h.store.Delete(r.Context(), int64(n))
	if err != nil {
		writeRes(w, http.StatusInternalServerError, err)
		return
	}
	writeRes(w, http.StatusOK, toCharacterResponse(character))
}

func (h *CharacterHandler) UpdateCharacter(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid character ID")
		return
	}

	var req updateCharacterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	// Do validation on the req

	characterParams := sqlc.UpdateCharacterParams{
		ID: id,
	}

	if req.Name != nil {
		characterParams.Name = sql.NullString{String: *req.Name, Valid: true}
	}
	if req.Nickname != nil {
		characterParams.Nickname = sql.NullString{String: *req.Nickname, Valid: true}
	}
	if req.Clan != nil {
		characterParams.Clan = sql.NullString{String: *req.Clan, Valid: true}
	}
	if req.Age != nil {
		characterParams.Age = sql.NullInt64{Int64: *req.Age, Valid: true}
	}
	if req.Rank != nil {
		characterParams.Rank = sql.NullString{String: *req.Rank, Valid: true}
	}
	if req.VillageID != nil {
		characterParams.VillageID = sql.NullInt64{Int64: *req.VillageID, Valid: true}
	}
	if req.Birthdate != nil {
		birthdate, err := time.Parse(time.RFC3339, *req.Birthdate)
		if err != nil {
			writeRes(w, http.StatusBadRequest, "Invalid date")
			return
		}
		characterParams.Birthdate = sql.NullTime{Time: birthdate, Valid: true}
	}

	updatedCharacter, err := h.store.Update(r.Context(), characterParams)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeRes(w, http.StatusNotFound, "Character not found")
			return
		}
		writeRes(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeRes(w, http.StatusOK, toCharacterResponse(updatedCharacter))
}
