package store

import (
	"context"
	"fmt"

	"github.com/biplob-codes/mockly/internal/database/sqlc"
)

type DBJutsuStore struct {
	db *sqlc.Queries
}

func CreateDBJutsuStore(db *sqlc.Queries) *DBJutsuStore {
	return &DBJutsuStore{db: db}
}

func (j *DBJutsuStore) Create(ctx context.Context, p sqlc.CreateJutsuParams) (sqlc.Jutsu, error) {
	jutsu, err := j.db.CreateJutsu(ctx, p)
	if err != nil {
		return sqlc.Jutsu{}, fmt.Errorf("Something went wrong")
	}
	return jutsu, nil
}
func (j *DBJutsuStore) List(ctx context.Context) ([]sqlc.Jutsu, error) {
	jutsus, err := j.db.GetJutsus(ctx)
	if err != nil {
		return []sqlc.Jutsu{}, fmt.Errorf("Something went wrong")
	}
	return jutsus, nil
}
func (j *DBJutsuStore) Get(ctx context.Context, id int64) (sqlc.Jutsu, error) {
	jutsu, err := j.db.GetJutsu(ctx, id)
	if err != nil {
		return sqlc.Jutsu{}, fmt.Errorf("Jutsu not found.")
	}
	return jutsu, nil
}
func (j *DBJutsuStore) Delete(ctx context.Context, id int64) (sqlc.Jutsu, error) {
	jutsu, err := j.db.DeleteJutsu(ctx, id)
	if err != nil {
		return sqlc.Jutsu{}, fmt.Errorf("Jutsu delete failed")
	}
	return jutsu, nil
}

func (j *DBJutsuStore) Update(ctx context.Context, p sqlc.UpdateJutsuParams) (sqlc.Jutsu, error) {
	jutsu, err := j.db.UpdateJutsu(ctx, p)
	if err != nil {
		return sqlc.Jutsu{}, fmt.Errorf("Something went wrong")
	}
	return jutsu, nil
}
