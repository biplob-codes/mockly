package seed
 

import (
	"database/sql"
	"time"

	"github.com/biplob-codes/mockly/internal/database/sqlc"
)

var VillageData = []sqlc.Village{
	{
		ID:          1,
		Name:        "Konohagakure",
		Description: sql.NullString{String: "Village Hidden in the Leaves, home of the Will of Fire", Valid: true},
		Land:        "Fire",
		Population:  sql.NullInt64{Int64: 45000, Valid: true},
		KageID:      sql.NullInt64{Valid: false},
		FoundedAt:   time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          2,
		Name:        "Sunagakure",
		Description: sql.NullString{String: "Village Hidden in the Sand, known for puppet mastery", Valid: true},
		Land:        "Wind",
		Population:  sql.NullInt64{Int64: 15000, Valid: true},
		KageID:      sql.NullInt64{Valid: false},
		FoundedAt:   time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          3,
		Name:        "Kirigakure",
		Description: sql.NullString{String: "Village Hidden in the Mist, once feared for the Bloody Mist era", Valid: true},
		Land:        "Water",
		Population:  sql.NullInt64{Int64: 20000, Valid: true},
		KageID:      sql.NullInt64{Valid: false},
		FoundedAt:   time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          4,
		Name:        "Iwagakure",
		Description: sql.NullString{String: "Village Hidden in the Rocks, built into a mountain stronghold", Valid: true},
		Land:        "Earth",
		Population:  sql.NullInt64{Int64: 18000, Valid: true},
		KageID:      sql.NullInt64{Valid: false},
		FoundedAt:   time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          5,
		Name:        "Kumogakure",
		Description: sql.NullString{String: "Village Hidden in the Clouds, known for its lightning-fast shinobi", Valid: true},
		Land:        "Lightning",
		Population:  sql.NullInt64{Int64: 17000, Valid: true},
		KageID:      sql.NullInt64{Valid: false},
		FoundedAt:   time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
}