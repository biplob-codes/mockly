package seed

import (
	"time"

	"github.com/biplob-codes/mockly/internal/database/sqlc"
)

var TeamData = []sqlc.Team{
	{ID: 1, Name: "Squad Seven", CreatedAt: time.Now()},
	{ID: 2, Name: "The Yellow Flash Corps", CreatedAt: time.Now()},
	{ID: 3, Name: "The Toad Sage's Disciples", CreatedAt: time.Now()},
	{ID: 4, Name: "The Legendary Sannin", CreatedAt: time.Now()},
}
