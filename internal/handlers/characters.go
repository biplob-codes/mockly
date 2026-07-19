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

type CharacterHandler struct {
	store store.CharacterStore
}
type createCharacterRequest struct {
	Name      string `json:"name"`
	Nickname  string `json:"nickname"`
	Clan      string `json:"clan"`
	Age       int64  `json:"age"`
	Rank      string `json:"rank"`
	Birthdate string `json:"birthdate"`
	VillageID int64  `json:"village_id"`
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
	writeRes(w, http.StatusOK, characters)
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

	writeRes(w, http.StatusOK, character)
}

func (h *CharacterHandler) CreateCharacter(w http.ResponseWriter, r *http.Request) {
	var req createCharacterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	birthdate, err := time.Parse(time.RFC3339, req.Birthdate)
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
	writeRes(w, http.StatusCreated, newCharacter)
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
	writeRes(w, http.StatusOK, character)
}