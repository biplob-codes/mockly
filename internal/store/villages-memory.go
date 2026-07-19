package store

import (
	"fmt"
	"time"
	"github.com/biplob-codes/mockly/internal/database/sqlc"
	"github.com/biplob-codes/mockly/internal/seed"
)


type MemoryVillageStore struct{
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

func CreateMemoryVillageStore() *MemoryVillageStore{
	return &MemoryVillageStore{villages: seed.VillageData}
}

func (v *MemoryVillageStore) List() ([]sqlc.Village,error){
	return  v.villages,nil;
}

func (v *MemoryVillageStore) Get(id int64) (sqlc.Village,error){
	village,found:=findVillageById(id)
	if found{
		return village,nil;
	}
	return  sqlc.Village{},fmt.Errorf("Village not found");
}

func (v *MemoryVillageStore) Create(nv sqlc.CreateVillageParams) (sqlc.Village,error){
	var cv sqlc.Village
	cv.ID=int64(len(seed.VillageData)+1)
	cv.CreatedAt=time.Now()
	cv.UpdatedAt=time.Now()
	return  cv,nil;
}

func (v *MemoryVillageStore) Delete (id int64) (sqlc.Village,error){
	village,found:=findVillageById(id)
	if !found{
		return sqlc.Village{},fmt.Errorf("Village not found")
	}
	return village,nil;
}