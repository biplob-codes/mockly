package store

import (
	"context"
	"fmt"
	// "time"

	"github.com/biplob-codes/mockly/internal/database/sqlc"
	"github.com/biplob-codes/mockly/internal/seed"
)

type MemoryTeamStore struct {
	teams []sqlc.Team
}

func CreateMemoryTeamStore() *MemoryTeamStore {
	return &MemoryTeamStore{teams: seed.TeamData}
}

// func (s *MemoryTeamStore) Create(ctx context.Context, p sqlc.CreateTeamParams) (sqlc.Team, error) {
// 	senseiExists := false
// 	for _, c := range seed.CharacterData {
// 		if c.ID == p.SenseiID {
// 			senseiExists = true
// 			break
// 		}
// 	}
// 	if !senseiExists {
// 		return sqlc.Team{}, fmt.Errorf("sensei does not exist")
// 	}

// 	for _, t := range s.teams {
// 		if t.SenseiID == p.SenseiID {
// 			return sqlc.Team{}, fmt.Errorf("sensei already leads a team")
// 		}
// 	}

// 	newTeam := sqlc.Team{
// 		ID:        int64(len(s.teams) + 1),
// 		Name:      p.Name,
// 		SenseiID:  p.SenseiID,
// 		CreatedAt: time.Now(),
// 	}
// 	return newTeam, nil
// }

func (s *MemoryTeamStore) Delete(ctx context.Context, id int64) (sqlc.Team, error) {
	for _, t := range s.teams {
		if t.ID == id {
			return t, nil
		}
	}
	return sqlc.Team{}, fmt.Errorf("team not found")
}
func (s *MemoryTeamStore) Get(ctx context.Context, id int64) (sqlc.Team, error) {
	for _, t := range s.teams {
		if t.ID == id {
			return t, nil
		}
	}
	return sqlc.Team{}, fmt.Errorf("team not found")
}

//	func (s *MemoryTeamStore) GetMembers(ctx context.Context, teamId int64) ([]sqlc.Character, error) {
//		var members []sqlc.Character
//		for _, c := range seed.CharacterData {
//			if c.TeamID.Valid && c.TeamID.Int64 == teamId {
//				members = append(members, c)
//			}
//		}
//		return members, nil
//	}
func (s *MemoryTeamStore) List(ctx context.Context) ([]sqlc.Team, error) {
	return s.teams, nil
}
