package seed

import (
	"database/sql"
	"time"

	"github.com/biplob-codes/mockly/internal/database/sqlc"
)

var VillageData = []sqlc.Village{
	{
		ID:          1,
		Name:        "Hidden Leaf Village",
		Description: sql.NullString{String: "The most powerful of the five great shinobi villages, built in a forest clearing and shaped by the Will of Fire philosophy passed down through generations of Hokage. It has produced legendary shinobi like the Sannin and remains the political and military heart of the Land of Fire.", Valid: true},
		Land:        "Fire",
		Population:  sql.NullInt64{Int64: 45000, Valid: true},
		KageID:      sql.NullInt64{Valid: false},
		FoundedAt:   time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          2,
		Name:        "Hidden Sand Village",
		Description: sql.NullString{String: "A village built into the harsh desert of the Land of Wind, where scarce resources forged a shinobi culture centered on efficiency and resourcefulness. Renowned for its mastery of puppet techniques and wind-based jutsu, it maintains a long-standing alliance with the Hidden Leaf Village.", Valid: true},
		Land:        "Wind",
		Population:  sql.NullInt64{Int64: 15000, Valid: true},
		KageID:      sql.NullInt64{Valid: false},
		FoundedAt:   time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          3,
		Name:        "Hidden Mist Village",
		Description: sql.NullString{String: "Situated on an island in the Land of Water, this village was once notorious for its brutal graduation trials during the era known as the Bloody Mist. It has since reformed under new leadership and is recognized today for its powerful swordsmen and mastery of water techniques.", Valid: true},
		Land:        "Water",
		Population:  sql.NullInt64{Int64: 20000, Valid: true},
		KageID:      sql.NullInt64{Valid: false},
		FoundedAt:   time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          4,
		Name:        "Hidden Rock Village",
		Description: sql.NullString{String: "Carved into the mountainous terrain of the Land of Earth, this village is known for its rugged, defensive fortifications and disciplined shinobi. Its forces played a decisive role in the Third Shinobi World War and it remains respected for its mastery of earth-style jutsu.", Valid: true},
		Land:        "Earth",
		Population:  sql.NullInt64{Int64: 18000, Valid: true},
		KageID:      sql.NullInt64{Valid: false},
		FoundedAt:   time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          5,
		Name:        "Hidden Cloud Village",
		Description: sql.NullString{String: "Perched high in the mountains of the Land of Lightning, this village is famed for producing some of the fastest and most physically powerful shinobi in the world. Its strict training regimens and emphasis on raw strength have made it a formidable military power among the great villages.", Valid: true},
		Land:        "Lightning",
		Population:  sql.NullInt64{Int64: 17000, Valid: true},
		KageID:      sql.NullInt64{Valid: false},
		FoundedAt:   time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          6,
		Name:        "Hidden Rain Village",
		Description: sql.NullString{String: "A perpetually rain-soaked village in a neutral region caught between larger nations, shaped by relentless war and poverty in its ninja's formative years. Though not one of the five great villages, it became a major power under Pain's rule and served as the birthplace of the Akatsuki organization.", Valid: true},
		Land:        "Rain",
		Population:  sql.NullInt64{Int64: 10000, Valid: true},
		KageID:      sql.NullInt64{Valid: false},
		FoundedAt:   time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
}
