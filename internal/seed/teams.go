package seed

import (
	"time"

	"github.com/biplob-codes/mockly/internal/database/sqlc"
)

var TeamData = []sqlc.Team{
	{
		ID:        1,
		Name:      "Team Naruto",
		SenseiID:  1, // Naruto
		CreatedAt: time.Now(),
	},
	{
		ID:        2,
		Name:      "Team Gaara",
		SenseiID:  2, // Gaara
		CreatedAt: time.Now(),
	},
}