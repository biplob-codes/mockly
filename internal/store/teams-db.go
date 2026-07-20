package store

import (
	"context"
	"database/sql"

	"github.com/biplob-codes/mockly/internal/database/sqlc"
)

type DBTeamStore struct {
	q *sqlc.Queries
}

func CreateDBTeamStore(q *sqlc.Queries) *DBTeamStore {
	return &DBTeamStore{q: q}
}

func (s *DBTeamStore) Create(ctx context.Context, p sqlc.CreateTeamParams) (sqlc.Team, error) {
	return s.q.CreateTeam(ctx, p)
}

func (s *DBTeamStore) Delete(ctx context.Context, id int64) (sqlc.Team, error) {
	return s.q.DeleteTeam(ctx, id)
}

func (s *DBTeamStore) Get(ctx context.Context, id int64) (sqlc.Team, error) {
	return s.q.GetTeam(ctx, id)
}

func (s *DBTeamStore) GetMembers(ctx context.Context, teamId int64) ([]sqlc.Character, error) {
	return s.q.GetTeamMembers(ctx, sql.NullInt64{Int64: teamId, Valid: true})
}

func (s *DBTeamStore) List(ctx context.Context) ([]sqlc.Team, error) {
	return s.q.ListTeams(ctx)
}