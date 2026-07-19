package store

import (
	"context"

	"github.com/biplob-codes/mockly/internal/database/sqlc"
)
type VillageStore interface{
	Create(ctx context.Context, v sqlc.CreateVillageParams) (sqlc.Village,error)
	List(ctx context.Context, ) ([]sqlc.Village,error)
	Get(ctx context.Context, id int64) (sqlc.Village,error)
	Delete(ctx context.Context, id int64) (sqlc.Village,error)
}
type CharacterStore interface{
	Create(ctx context.Context, v sqlc.CreateCharacterParams) (sqlc.Character,error)
	List(ctx context.Context, ) ([]sqlc.Character,error)
	Get(ctx context.Context, id int64) (sqlc.Character,error)
	Delete(ctx context.Context, id int64) (sqlc.Character,error)
}
