package seed

import (
	"database/sql"
	"time"

	"github.com/biplob-codes/mockly/internal/database/sqlc"
)

var JutsuData = []sqlc.Jutsu{
	{
		ID:          1,
		Name:        "Rasengan",
		Description: sql.NullString{String: "A spiraling sphere of concentrated chakra", Valid: true},
		Type:        "Ninjutsu",
		Rank:        "A",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          2,
		Name:        "Chidori",
		Description: sql.NullString{String: "A lightning-based technique that pierces through targets", Valid: true},
		Type:        "Ninjutsu",
		Rank:        "A",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          3,
		Name:        "Sharingan Genjutsu",
		Description: sql.NullString{String: "Illusion cast through direct eye contact with the Sharingan", Valid: true},
		Type:        "Genjutsu",
		Rank:        "B",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          4,
		Name:        "Shadow Clone Jutsu",
		Description: sql.NullString{String: "Creates solid clones capable of independent action", Valid: true},
		Type:        "Ninjutsu",
		Rank:        "B",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          5,
		Name:        "Eight Gates",
		Description: sql.NullString{String: "Releases the body's chakra limiters at great physical cost", Valid: true},
		Type:        "Taijutsu",
		Rank:        "S",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
}