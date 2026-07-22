package store

import (
	"context"

	"github.com/biplob-codes/mockly/internal/database/sqlc"
)

type VillageStore interface {
	Create(ctx context.Context, v sqlc.CreateVillageParams) (sqlc.Village, error)
	List(ctx context.Context) ([]sqlc.Village, error)
	Get(ctx context.Context, id int64) (sqlc.Village, error)
	Delete(ctx context.Context, id int64) (sqlc.Village, error)
	Update(ctx context.Context, p sqlc.UpdateVillageParams) (sqlc.Village, error)
}
type CharacterStore interface {
	Create(ctx context.Context, v sqlc.CreateCharacterParams) (sqlc.Character, error)
	List(ctx context.Context) ([]sqlc.Character, error)
	Get(ctx context.Context, id int64) (sqlc.Character, error)
	Delete(ctx context.Context, id int64) (sqlc.Character, error)
}
type JutsuStore interface {
	Create(ctx context.Context, p sqlc.CreateJutsuParams) (sqlc.Jutsu, error)
	List(ctx context.Context) ([]sqlc.Jutsu, error)
	Get(ctx context.Context, id int64) (sqlc.Jutsu, error)
	Delete(ctx context.Context, id int64) (sqlc.Jutsu, error)
}
type CharacterJutsuStore interface {
	Create(ctx context.Context, p sqlc.CreateCharacterJutsuParams) (sqlc.CharactersJutsu, error)
	ListJutsusByCharacter(ctx context.Context, id int64) ([]sqlc.Jutsu, error)
	ListCharactersByJutsu(ctx context.Context, id int64) ([]sqlc.Character, error)
	Delete(ctx context.Context, characterId, jutsuId int64) (sqlc.CharactersJutsu, error)
}
type TeamStore interface {
	Create(ctx context.Context, p sqlc.CreateTeamParams) (sqlc.Team, error)
	Delete(ctx context.Context, id int64) (sqlc.Team, error)
	Get(ctx context.Context, id int64) (sqlc.Team, error)
	GetMembers(ctx context.Context, teamId int64) ([]sqlc.Character, error)
	List(ctx context.Context) ([]sqlc.Team, error)
}
type MissionStore interface {
	Create(ctx context.Context, p sqlc.CreateMissionParams) (sqlc.Mission, error)
	Get(ctx context.Context, id int64) (sqlc.Mission, error)
	List(ctx context.Context) ([]sqlc.Mission, error)
	GetByTeam(ctx context.Context, teamId int64) ([]sqlc.Mission, error)
	Delete(ctx context.Context, id int64) (sqlc.Mission, error)
}
