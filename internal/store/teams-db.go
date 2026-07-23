package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/biplob-codes/mockly/internal/database/sqlc"
)

type DBTeamStore struct {
	db *sql.DB
	q  *sqlc.Queries
}

func CreateDBTeamStore(db *sql.DB, q *sqlc.Queries) *DBTeamStore {
	return &DBTeamStore{db: db, q: q}
}

func (s *DBTeamStore) Create(ctx context.Context, name string, senseiID int64) (sqlc.Team, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return sqlc.Team{}, err
	}
	defer tx.Rollback()

	qtx := s.q.WithTx(tx)

	count, err := qtx.CharacterIsSenseiOfAnyTeam(ctx, senseiID)
	if err != nil {
		return sqlc.Team{}, err
	}
	if count > 0 {
		return sqlc.Team{}, fmt.Errorf("Character is already a sensei")
	}

	team, err := qtx.CreateTeam(ctx, name)
	if err != nil {
		return sqlc.Team{}, err
	}

	_, err = qtx.AddTeamMember(ctx, sqlc.AddTeamMemberParams{
		TeamID:      team.ID,
		CharacterID: senseiID,
		Role:        "Sensei",
	})
	if err != nil {
		return sqlc.Team{}, err
	}

	if err := tx.Commit(); err != nil {
		return sqlc.Team{}, err
	}

	return team, nil
}

func (s *DBTeamStore) Get(ctx context.Context, id int64) (sqlc.Team, error) {
	return s.q.GetTeam(ctx, id)
}

func (s *DBTeamStore) GetDetails(ctx context.Context, id int64) ([]sqlc.GetTeamDetailsRow, error) {
	return s.q.GetTeamDetails(ctx, id)
}

func (s *DBTeamStore) List(ctx context.Context) ([]sqlc.Team, error) {
	return s.q.ListTeams(ctx)
}

func (s *DBTeamStore) Update(ctx context.Context, p sqlc.UpdateTeamParams) (sqlc.Team, error) {
	return s.q.UpdateTeam(ctx, p)
}

func (s *DBTeamStore) AddMember(ctx context.Context, p sqlc.AddTeamMemberParams) (sqlc.TeamMember, error) {
	return s.q.AddTeamMember(ctx, p)
}

func (s *DBTeamStore) RemoveMember(ctx context.Context, p sqlc.RemoveTeamMemberParams) (sqlc.TeamMember, error) {
	return s.q.RemoveTeamMember(ctx, p)
}

func (s *DBTeamStore) Delete(ctx context.Context, id int64) (sqlc.Team, error) {
	return s.q.DeleteTeam(ctx, id)
}

func (s *DBTeamStore) GetCharacterMembership(ctx context.Context, characterID int64) (sqlc.TeamMember, error) {
	return s.q.GetCharacterMembership(ctx, characterID)
}

func (s *DBTeamStore) GetTeamSensei(ctx context.Context, teamID int64) (sqlc.TeamMember, error) {
	return s.q.GetTeamSensei(ctx, teamID)
}
