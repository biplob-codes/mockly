package store

import (
	"context"
	"fmt"
	"time"

	"github.com/biplob-codes/mockly/internal/database/sqlc"
	"github.com/biplob-codes/mockly/internal/seed"
)

type MemoryMissionStore struct {
	missions []sqlc.Mission
}

func CreateMemoryMissionStore() *MemoryMissionStore {
	return &MemoryMissionStore{missions: seed.MissionData}
}

func (s *MemoryMissionStore) Create(ctx context.Context, p sqlc.CreateMissionParams) (sqlc.Mission, error) {
	newMission := sqlc.Mission{
		ID:          int64(len(s.missions) + 1),
		Name:        p.Name,
		Description: p.Description,
		AssignedTo:  p.AssignedTo,
		Rank:        p.Rank,
		Status:      p.Status,
		Reward:      p.Reward,
		StartsAt:    p.StartsAt,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return newMission, nil
}

func (s *MemoryMissionStore) Get(ctx context.Context, id int64) (sqlc.Mission, error) {
	for _, m := range s.missions {
		if m.ID == id {
			return m, nil
		}
	}
	return sqlc.Mission{}, fmt.Errorf("mission not found")
}

func (s *MemoryMissionStore) List(ctx context.Context) ([]sqlc.Mission, error) {
	return s.missions, nil
}

func (s *MemoryMissionStore) GetByTeam(ctx context.Context, teamId int64) ([]sqlc.Mission, error) {
	var result []sqlc.Mission
	for _, m := range s.missions {
		if m.AssignedTo == teamId {
			result = append(result, m)
		}
	}
	return result, nil
}

func (s *MemoryMissionStore) Delete(ctx context.Context, id int64) (sqlc.Mission, error) {
	for _, m := range s.missions {
		if m.ID == id {
			return m, nil
		}
	}
	return sqlc.Mission{}, fmt.Errorf("mission not found")
}