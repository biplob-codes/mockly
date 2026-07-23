package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/biplob-codes/mockly/internal/database/sqlc"
	"github.com/biplob-codes/mockly/internal/seed"
)

type MemoryTeamStore struct {
	teams   []sqlc.Team
	members []sqlc.TeamMember
}

func CreateMemoryTeamStore() *MemoryTeamStore {
	return &MemoryTeamStore{teams: seed.TeamData, members: seed.TeamMemberData}
}

// --- internal lookup helpers (read-only against seed data) ---

func findTeamById(id int64) (sqlc.Team, bool) {
	for _, t := range seed.TeamData {
		if t.ID == id {
			return t, true
		}
	}
	return sqlc.Team{}, false
}

func findMembership(characterID int64) (sqlc.TeamMember, bool) {
	for _, m := range seed.TeamMemberData {
		if m.CharacterID == characterID {
			return m, true
		}
	}
	return sqlc.TeamMember{}, false
}

func findMembershipInTeam(teamID, characterID int64) (sqlc.TeamMember, bool) {
	for _, m := range seed.TeamMemberData {
		if m.TeamID == teamID && m.CharacterID == characterID {
			return m, true
		}
	}
	return sqlc.TeamMember{}, false
}

func findTeamSensei(teamID int64) (sqlc.TeamMember, bool) {
	for _, m := range seed.TeamMemberData {
		if m.TeamID == teamID && m.Role == "Sensei" {
			return m, true
		}
	}
	return sqlc.TeamMember{}, false
}

// --- store methods ---
// findSenseiMembership returns a TeamMember row where the given character
// holds the "Sensei" role on some team (any team), if one exists.
func findSenseiMembership(characterID int64) (sqlc.TeamMember, bool) {
	for _, m := range seed.TeamMemberData {
		if m.CharacterID == characterID && m.Role == "Sensei" {
			return m, true
		}
	}
	return sqlc.TeamMember{}, false
}

// Create mirrors the intended validation (sensei exists, character isn't
// already a Sensei of some other team — being a student elsewhere is fine)
// but never appends to s.teams — it only returns a plausible faked Team.
func (s *MemoryTeamStore) Create(ctx context.Context, name string, senseiID int64) (sqlc.Team, error) {
	if _, found := findCharacterById(senseiID); !found {
		return sqlc.Team{}, fmt.Errorf("sensei does not exist")
	}

	if _, found := findSenseiMembership(senseiID); found {
		return sqlc.Team{}, fmt.Errorf("already a sensei of another team")
	}

	newTeam := sqlc.Team{
		ID:        int64(len(s.teams) + 1),
		Name:      name,
		CreatedAt: time.Now(),
	}
	return newTeam, nil
}

func (s *MemoryTeamStore) Get(ctx context.Context, id int64) (sqlc.Team, error) {
	team, found := findTeamById(id)
	if !found {
		return sqlc.Team{}, sql.ErrNoRows
	}
	return team, nil
}

// GetDetails joins the team against its members and each member's character
// row, built fresh on every call from the seed data — nothing is stored.
func (s *MemoryTeamStore) GetDetails(ctx context.Context, id int64) ([]sqlc.GetTeamDetailsRow, error) {
	team, found := findTeamById(id)
	if !found {
		return nil, sql.ErrNoRows
	}

	var rows []sqlc.GetTeamDetailsRow

	for _, m := range s.members {
		if m.TeamID != id {
			continue
		}

		row := sqlc.GetTeamDetailsRow{
			TeamID:        team.ID,
			TeamName:      team.Name,
			TeamCreatedAt: team.CreatedAt,
			Role:          sql.NullString{String: m.Role, Valid: true},
			JoinedAt:      sql.NullTime{Time: m.JoinedAt, Valid: true},
		}

		if c, found := findCharacterById(m.CharacterID); found {
			row.ID = sql.NullInt64{Int64: c.ID, Valid: true}
			row.Name = sql.NullString{String: c.Name, Valid: true}
			row.Nickname = c.Nickname
			row.Clan = c.Clan
			row.Age = c.Age
			row.Rank = sql.NullString{String: c.Rank, Valid: true}
			row.Birthdate = sql.NullTime{Time: c.Birthdate, Valid: true}
			row.VillageID = sql.NullInt64{Int64: c.VillageID, Valid: true}
			row.CreatedAt = sql.NullTime{Time: c.CreatedAt, Valid: true}
			row.UpdatedAt = sql.NullTime{Time: c.UpdatedAt, Valid: true}
		}

		rows = append(rows, row)
	}

	// A team with no members still exists — return a single row with
	// null character fields, same shape a LEFT JOIN would give you.
	if len(rows) == 0 {
		rows = append(rows, sqlc.GetTeamDetailsRow{
			TeamID:        team.ID,
			TeamName:      team.Name,
			TeamCreatedAt: team.CreatedAt,
		})
	}

	return rows, nil
}

func (s *MemoryTeamStore) List(ctx context.Context) ([]sqlc.Team, error) {
	return s.teams, nil
}

// Update returns a copy of the found team with the new name applied.
// s.teams (and seed.TeamData) is left untouched.
func (s *MemoryTeamStore) Update(ctx context.Context, p sqlc.UpdateTeamParams) (sqlc.Team, error) {
	team, found := findTeamById(p.ID)
	if !found {
		return sqlc.Team{}, sql.ErrNoRows
	}

	updated := team
	updated.Name = p.Name
	return updated, nil
}

// AddMember validates like the DB would (team exists, character exists,
// character not already a member of this team) and returns a faked
// TeamMember without touching seed.TeamMemberData.
func (s *MemoryTeamStore) AddMember(ctx context.Context, p sqlc.AddTeamMemberParams) (sqlc.TeamMember, error) {
	if _, found := findTeamById(p.TeamID); !found {
		return sqlc.TeamMember{}, fmt.Errorf("team not found")
	}
	if _, found := findCharacterById(p.CharacterID); !found {
		return sqlc.TeamMember{}, fmt.Errorf("character not found")
	}
	if _, found := findMembershipInTeam(p.TeamID, p.CharacterID); found {
		return sqlc.TeamMember{}, fmt.Errorf("character already a member of this team")
	}

	newMember := sqlc.TeamMember{
		TeamID:      p.TeamID,
		CharacterID: p.CharacterID,
		Role:        p.Role,
		JoinedAt:    time.Now(),
	}
	return newMember, nil
}

// RemoveMember returns the existing membership as proof it existed,
// without removing it from seed.TeamMemberData.
func (s *MemoryTeamStore) RemoveMember(ctx context.Context, p sqlc.RemoveTeamMemberParams) (sqlc.TeamMember, error) {
	member, found := findMembershipInTeam(p.TeamID, p.CharacterID)
	if !found {
		return sqlc.TeamMember{}, sql.ErrNoRows
	}
	return member, nil
}

// Delete returns the found team without removing it from s.teams.
func (s *MemoryTeamStore) Delete(ctx context.Context, id int64) (sqlc.Team, error) {
	team, found := findTeamById(id)
	if !found {
		return sqlc.Team{}, sql.ErrNoRows
	}
	return team, nil
}

func (s *MemoryTeamStore) GetCharacterMembership(ctx context.Context, characterID int64) (sqlc.TeamMember, error) {
	member, found := findMembership(characterID)
	if !found {
		return sqlc.TeamMember{}, sql.ErrNoRows
	}
	return member, nil
}

func (s *MemoryTeamStore) GetTeamSensei(ctx context.Context, teamID int64) (sqlc.TeamMember, error) {
	member, found := findTeamSensei(teamID)
	if !found {
		return sqlc.TeamMember{}, sql.ErrNoRows
	}
	return member, nil
}
