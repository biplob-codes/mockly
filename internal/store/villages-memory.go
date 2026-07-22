package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/biplob-codes/mockly/internal/database/sqlc"
	"github.com/biplob-codes/mockly/internal/seed"
)

type MemoryVillageStore struct {
	villages []sqlc.Village
}

func findVillageById(id int64) (sqlc.Village, bool) {
	for _, t := range seed.VillageData {
		if t.ID == id {
			return t, true
		}
	}
	return sqlc.Village{}, false
}

func CreateMemoryVillageStore() *MemoryVillageStore {
	return &MemoryVillageStore{villages: seed.VillageData}
}

func (v *MemoryVillageStore) List(ctx context.Context) ([]sqlc.Village, error) {
	return v.villages, nil
}

func (v *MemoryVillageStore) Get(ctx context.Context, id int64) (sqlc.Village, error) {
	village, found := findVillageById(id)
	if found {
		return village, nil
	}
	return sqlc.Village{}, fmt.Errorf("Village not found")
}

func (v *MemoryVillageStore) Create(ctx context.Context, cv sqlc.CreateVillageParams) (sqlc.Village, error) {
	var vlg sqlc.Village

	vlg.ID = int64(len(seed.VillageData) + 1)
	vlg.Name = cv.Name
	vlg.Description = cv.Description
	vlg.Land = cv.Land
	vlg.Population = cv.Population
	vlg.KageID = cv.KageID
	vlg.FoundedAt = cv.FoundedAt
	vlg.CreatedAt = time.Now()
	vlg.UpdatedAt = time.Now()

	return vlg, nil
}
func (v *MemoryVillageStore) Update(ctx context.Context, uv sqlc.UpdateVillageParams) (sqlc.Village, error) {
	var existing sqlc.Village
	found := false

	for _, vlg := range seed.VillageData {
		if vlg.ID == uv.ID {
			existing = vlg // struct copy, not a pointer — original slice untouched
			found = true
			break
		}
	}

	if !found {
		return sqlc.Village{}, sql.ErrNoRows
	}

	if uv.Name.Valid {
		existing.Name = uv.Name.String
	}
	if uv.Description.Valid {
		existing.Description = uv.Description
	}
	if uv.Land.Valid {
		existing.Land = uv.Land.String
	}
	if uv.Population.Valid {
		existing.Population = uv.Population
	}
	if uv.KageID.Valid {
		existing.KageID = uv.KageID
	}
	if uv.FoundedAt.Valid {
		existing.FoundedAt = uv.FoundedAt.Time
	}

	existing.UpdatedAt = time.Now()

	return existing, nil
}

func (v *MemoryVillageStore) Delete(ctx context.Context, id int64) (sqlc.Village, error) {
	village, found := findVillageById(id)
	if !found {
		return sqlc.Village{}, fmt.Errorf("Village not found")
	}
	return village, nil
}
