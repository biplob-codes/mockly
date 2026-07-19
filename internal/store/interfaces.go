package store
import (
	"github.com/biplob-codes/mockly/internal/database/sqlc"
)
type VillageStore interface{
	Create (v sqlc.CreateVillageParams) (sqlc.Village,error)
	List() ([]sqlc.Village,error)
	Get(id int64) (sqlc.Village,error)
	Delete(id int64) (sqlc.Village,error)
}
