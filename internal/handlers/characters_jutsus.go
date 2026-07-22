package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/biplob-codes/mockly/internal/database/sqlc"
	"github.com/biplob-codes/mockly/internal/store"
)

type CharacterJutsuHandler struct {
	store store.CharacterJutsuStore
}

type createCharacterJutsuRequest struct {
	JutsuID int64 `json:"jutsu_id"`
}

// internal/characterjutsu/response.go
type CharacterJutsuResponse struct {
	ID          int64  `json:"id"`
	CharacterID int64  `json:"characterId"`
	JutsuID     int64  `json:"jutsuId"`
	CreatedAt   string `json:"createdAt"`
}

func toCharacterJutsuResponse(cj sqlc.CharactersJutsu) CharacterJutsuResponse {
	return CharacterJutsuResponse{
		ID:          cj.ID,
		CharacterID: cj.CharacterID,
		JutsuID:     cj.JutsuID,
		CreatedAt:   cj.CreatedAt.Format(time.RFC3339),
	}
}
func NewCharacterJutsuHandler(s store.CharacterJutsuStore) *CharacterJutsuHandler {
	return &CharacterJutsuHandler{store: s}
}

func (h *CharacterJutsuHandler) CreateCharacterJutsu(w http.ResponseWriter, r *http.Request) {
	characterIdStr := r.PathValue("id")
	characterId, err := strconv.Atoi(characterIdStr)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid character id")
		return
	}

	var req createCharacterJutsuRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	// Do validation on the req

	params := sqlc.CreateCharacterJutsuParams{
		CharacterID: int64(characterId),
		JutsuID:     req.JutsuID,
	}

	newCharacterJutsu, err := h.store.Create(r.Context(), params)
	if err != nil {
		writeRes(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeRes(w, http.StatusCreated, toCharacterJutsuResponse(newCharacterJutsu))
}

func (h *CharacterJutsuHandler) ListJutsusByCharacter(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	n, err := strconv.Atoi(id)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid number")
		return
	}

	jutsus, err := h.store.ListJutsusByCharacter(r.Context(), int64(n))
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

func (h *CharacterJutsuHandler) ListCharactersByJutsu(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid number")
		return
	}

	characters, err := h.store.ListCharactersByJutsu(r.Context(), int64(n))
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

func (h *CharacterJutsuHandler) DeleteCharacterJutsu(w http.ResponseWriter, r *http.Request) {
	characterIdStr := r.PathValue("characterId")
	characterId, err := strconv.Atoi(characterIdStr)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid character id")
		return
	}

	jutsuIdStr := r.PathValue("jutsuId")
	jutsuId, err := strconv.Atoi(jutsuIdStr)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid jutsu id")
		return
	}

	deleted, err := h.store.Delete(r.Context(), int64(characterId), int64(jutsuId))
	if err != nil {
		writeRes(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeRes(w, http.StatusOK, toCharacterJutsuResponse(deleted))
}
