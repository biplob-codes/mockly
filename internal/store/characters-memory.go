package store

import (
	"context"
	"fmt"
	"time"

	"github.com/biplob-codes/mockly/internal/database/sqlc"
	"github.com/biplob-codes/mockly/internal/seed"
)

type MemoryCharacterStore struct {
	characters []sqlc.Character
}

func findCharacterById(id int64) (sqlc.Character, bool) {
	for _, c := range seed.CharacterData {
		if c.ID == id {
			return c, true
		}
	}
	return sqlc.Character{}, false
}

func CreateMemoryCharacterStore() *MemoryCharacterStore {
	return &MemoryCharacterStore{characters: seed.CharacterData}
}

func (s *MemoryCharacterStore) List(ctx context.Context) ([]sqlc.Character, error) {
	return s.characters, nil
}

func (s *MemoryCharacterStore) Get(ctx context.Context, id int64) (sqlc.Character, error) {
	character, found := findCharacterById(id)
	if found {
		return character, nil
	}
	return sqlc.Character{}, fmt.Errorf("Character not found")
}

func (s *MemoryCharacterStore) Create(ctx context.Context, nc sqlc.CreateCharacterParams) (sqlc.Character, error) {
	cc := sqlc.Character{
		ID:        int64(len(seed.CharacterData) + 1),
		Name:      nc.Name,
		Nickname:  nc.Nickname,
		Clan:      nc.Clan,
		Age:       nc.Age,
		Rank:      nc.Rank,
		Birthdate: nc.Birthdate,
		VillageID: nc.VillageID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return cc, nil
}

func (s *MemoryCharacterStore) Delete(ctx context.Context, id int64) (sqlc.Character, error) {
	character, found := findCharacterById(id)
	if !found {
		return sqlc.Character{}, fmt.Errorf("Character not found")
	}
	return character, nil
}