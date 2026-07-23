package store

import (
	"context"
	"fmt"

	"github.com/biplob-codes/mockly/internal/database/sqlc"
)

type DBCharacterStore struct {
	db *sqlc.Queries
}

func CreateDBCharacterStore(db *sqlc.Queries) *DBCharacterStore {
	return &DBCharacterStore{db: db}
}

func (c *DBCharacterStore) Create(ctx context.Context, p sqlc.CreateCharacterParams) (sqlc.Character, error) {
	character, err := c.db.CreateCharacter(ctx, p)
	if err != nil {
		return sqlc.Character{}, fmt.Errorf("Something went wrong")
	}
	return character, nil
}
func (c *DBCharacterStore) List(ctx context.Context) ([]sqlc.Character, error) {
	characters, err := c.db.GetCharacters(ctx)
	if err != nil {
		return []sqlc.Character{}, fmt.Errorf("Something went wrong")
	}
	return characters, nil
}
func (c *DBCharacterStore) Get(ctx context.Context, id int64) (sqlc.Character, error) {
	character, err := c.db.GetCharacter(ctx, id)
	if err != nil {
		return sqlc.Character{}, fmt.Errorf("Character not found.")
	}
	return character, nil
}
func (c *DBCharacterStore) Delete(ctx context.Context, id int64) (sqlc.Character, error) {
	character, err := c.db.DeleteCharacter(ctx, id)
	if err != nil {
		return sqlc.Character{}, fmt.Errorf("Character delete failed")
	}
	return character, nil
}
func (c *DBCharacterStore) Update(ctx context.Context, p sqlc.UpdateCharacterParams) (sqlc.Character, error) {
	character, err := c.db.UpdateCharacter(ctx, p)
	if err != nil {
		return sqlc.Character{}, fmt.Errorf("Something went wrong")
	}
	return character, nil
}
