package store

import (
	"context"

	"github.com/biplob-codes/mockly/internal/database/sqlc"
)

type DBCharacterJutsuStore struct {
	q *sqlc.Queries
}

func CreateDBCharacterJutsuStore(q *sqlc.Queries) *DBCharacterJutsuStore {
	return &DBCharacterJutsuStore{q: q}
}

func (cjs *DBCharacterJutsuStore) Create(ctx context.Context, p sqlc.CreateCharacterJutsuParams) (sqlc.CharactersJutsu, error) {
	return cjs.q.CreateCharacterJutsu(ctx, p)
}

func (cjs *DBCharacterJutsuStore) ListJutsusByCharacter(ctx context.Context, id int64) ([]sqlc.Jutsu, error) {
	return cjs.q.GetJutsusByCharacterID(ctx, id)
}

func (cjs *DBCharacterJutsuStore) ListCharactersByJutsu(ctx context.Context, id int64) ([]sqlc.Character, error) {
	return cjs.q.GetCharactersByJutsuID(ctx, id)
}

func (cjs *DBCharacterJutsuStore) Delete(ctx context.Context, characterId, jutsuId int64) (sqlc.CharactersJutsu, error) {
	return cjs.q.DeleteCharacterJutsu(ctx, sqlc.DeleteCharacterJutsuParams{
		CharacterID: characterId,
		JutsuID:     jutsuId,
	})
}