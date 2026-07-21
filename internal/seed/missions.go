package seed

import (
	"time"

	"github.com/biplob-codes/mockly/internal/database/sqlc"
)

var MissionData = []sqlc.Mission{
	{
		ID:          1,
		Name:        "Find Tora the Cat",
		Description: "Locate and retrieve the feudal lord's wife's missing pet cat.",
		AssignedTo:  1,
		Rank:        "D",
		Status:      "Success",
		Reward:      15000,
		StartsAt:    time.Date(2026, 1, 5, 9, 0, 0, 0, time.UTC),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          2,
		Name:        "Escort to the Land of Waves",
		Description: "Protect the bridge builder Tazuna during his journey home.",
		AssignedTo:  1,
		Rank:        "C",
		Status:      "Success",
		Reward:      80000,
		StartsAt:    time.Date(2026, 1, 20, 8, 0, 0, 0, time.UTC),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          3,
		Name:        "Retrieve the Fifth Hokage",
		Description: "Track down Tsunade and convince her to return to the village.",
		AssignedTo:  1,
		Rank:        "B",
		Status:      "Pending",
		Reward:      200000,
		StartsAt:    time.Date(2026, 2, 1, 7, 0, 0, 0, time.UTC),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          4,
		Name:        "Suna Village Reconstruction",
		Description: "Assist Suna in rebuilding defenses after the recent conflict.",
		AssignedTo:  2,
		Rank:        "A",
		Status:      "Pending",
		Reward:      500000,
		StartsAt:    time.Date(2026, 2, 10, 6, 0, 0, 0, time.UTC),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
}