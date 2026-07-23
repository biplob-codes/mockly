package seed

import (
	"time"

	"github.com/biplob-codes/mockly/internal/database/sqlc"
)

var TeamData = []sqlc.Team{
	{
		ID:   1,
		Name: "Team Naruto",

		CreatedAt: time.Now(),
	},
	{
		ID:   2,
		Name: "Team Gaara",

		CreatedAt: time.Now(),
	},
}
