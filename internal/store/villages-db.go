package store

import (
	"context"
	"fmt"

	"github.com/biplob-codes/mockly/internal/database/sqlc"
)

type DBVillageStore struct {
	db *sqlc.Queries
}

func CreateDBVillageStore(db *sqlc.Queries) *DBVillageStore {
	return &DBVillageStore{db: db}
}

func (v *DBVillageStore) Create(ctx context.Context, p sqlc.CreateVillageParams) (sqlc.Village, error) {
	village, err := v.db.CreateVillage(ctx, p)
	if err != nil {
		return sqlc.Village{}, fmt.Errorf("Something went wrong")
	}
	return village, nil
}
func (v *DBVillageStore) List(ctx context.Context) ([]sqlc.Village, error) {
	villages, err := v.db.GetVillages(ctx)
	if err != nil {
		return []sqlc.Village{}, fmt.Errorf("Something went wrong")
	}
	return villages, nil
}
func (v *DBVillageStore) Get(ctx context.Context, id int64) (sqlc.Village, error) {
	village, err := v.db.GetVillage(ctx, id)
	if err != nil {
		return sqlc.Village{}, fmt.Errorf("Village not found.")
	}
	return village, nil
}

func (v *DBVillageStore) Update(ctx context.Context, p sqlc.UpdateVillageParams) (sqlc.Village, error) {
	return v.db.UpdateVillage(ctx, p)
}

func (v *DBVillageStore) Delete(ctx context.Context, id int64) (sqlc.Village, error) {
	village, err := v.db.DeleteVillage(ctx, id)
	if err != nil {
		return sqlc.Village{}, fmt.Errorf("Village delete failed")
	}
	return village, nil
}
