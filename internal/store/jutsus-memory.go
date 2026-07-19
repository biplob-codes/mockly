package store

import (
	"context"
	"fmt"
	"time"

	"github.com/biplob-codes/mockly/internal/database/sqlc"
	"github.com/biplob-codes/mockly/internal/seed"
)



type MemoryJutsuStore struct {
	jutsus []sqlc.Jutsu
}

func findJutsuById(id int64) (sqlc.Jutsu, bool) {
	for _, j := range seed.JutsuData {
		if j.ID == id {
			return j, true
		}
	}
	return sqlc.Jutsu{}, false
}

func CreateMemoryJutsuStore() *MemoryJutsuStore {
	return &MemoryJutsuStore{jutsus: seed.JutsuData}
}

func (s *MemoryJutsuStore) List(ctx context.Context) ([]sqlc.Jutsu, error) {
	return s.jutsus, nil
}

func (s *MemoryJutsuStore) Get(ctx context.Context, id int64) (sqlc.Jutsu, error) {
	jutsu, found := findJutsuById(id)
	if found {
		return jutsu, nil
	}
	return sqlc.Jutsu{}, fmt.Errorf("Jutsu not found")
}

func (s *MemoryJutsuStore) Create(ctx context.Context, nj sqlc.CreateJutsuParams) (sqlc.Jutsu, error) {
	cj := sqlc.Jutsu{
		ID:          int64(len(seed.JutsuData) + 1),
		Name:        nj.Name,
		Description: nj.Description,
		Type:        nj.Type,
		Rank:        nj.Rank,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return cj, nil
}

func (s *MemoryJutsuStore) Delete(ctx context.Context, id int64) (sqlc.Jutsu, error) {
	jutsu, found := findJutsuById(id)
	if !found {
		return sqlc.Jutsu{}, fmt.Errorf("Jutsu not found")
	}
	return jutsu, nil
}