package store

import (
	"context"

	"github.com/biplob-codes/mockly/internal/database/sqlc"
)

type DBMissionStore struct {
	q *sqlc.Queries
}

func CreateDBMissionStore(q *sqlc.Queries) *DBMissionStore {
	return &DBMissionStore{q: q}
}

func (s *DBMissionStore) Create(ctx context.Context, p sqlc.CreateMissionParams) (sqlc.Mission, error) {
	return s.q.CreateMission(ctx, p)
}

func (s *DBMissionStore) Get(ctx context.Context, id int64) (sqlc.Mission, error) {
	return s.q.GetMission(ctx, id)
}

func (s *DBMissionStore) List(ctx context.Context) ([]sqlc.Mission, error) {
	return s.q.ListMissions(ctx)
}

func (s *DBMissionStore) GetByTeam(ctx context.Context, teamId int64) ([]sqlc.Mission, error) {
	return s.q.GetMissionsByTeam(ctx, teamId)
}

func (s *DBMissionStore) Delete(ctx context.Context, id int64) (sqlc.Mission, error) {
	return s.q.DeleteMission(ctx, id)
}