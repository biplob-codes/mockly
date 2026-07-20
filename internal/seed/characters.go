package seed

import (
	"database/sql"
	"time"

	"github.com/biplob-codes/mockly/internal/database/sqlc"
)

var CharacterData = []sqlc.Character{
	{
		ID:        1,
		Name:      "Naruto Uzumaki",
		Nickname:  sql.NullString{String: "Seventh Hokage", Valid: true},
		Clan:      sql.NullString{String: "Uzumaki", Valid: true},
		Age:       sql.NullInt64{Int64: 17, Valid: true},
		Rank:      "Kage",
		Birthdate: time.Date(1997, 10, 10, 0, 0, 0, 0, time.UTC),
		VillageID: 1,
		TeamID:    sql.NullInt64{Int64: 1, Valid: true}, // sensei of Team Naruto
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		ID:        2,
		Name:      "Gaara",
		Nickname:  sql.NullString{String: "Fifth Kazekage", Valid: true},
		Clan:      sql.NullString{Valid: false},
		Age:       sql.NullInt64{Int64: 16, Valid: true},
		Rank:      "Kage",
		Birthdate: time.Date(1988, 1, 19, 0, 0, 0, 0, time.UTC),
		VillageID: 2,
		TeamID:    sql.NullInt64{Int64: 2, Valid: true}, // sensei of Team Gaara
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		ID:        3,
		Name:      "Mei Terumi",
		Nickname:  sql.NullString{String: "Fifth Mizukage", Valid: true},
		Clan:      sql.NullString{Valid: false},
		Age:       sql.NullInt64{Int64: 33, Valid: true},
		Rank:      "Kage",
		Birthdate: time.Date(1970, 12, 17, 0, 0, 0, 0, time.UTC),
		VillageID: 3,
		TeamID:    sql.NullInt64{Int64: 2, Valid: true},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		ID:        4,
		Name:      "Onoki",
		Nickname:  sql.NullString{String: "Third Tsuchikage", Valid: true},
		Clan:      sql.NullString{Valid: false},
		Age:       sql.NullInt64{Int64: 82, Valid: true},
		Rank:      "Kage",
		Birthdate: time.Date(1918, 3, 1, 0, 0, 0, 0, time.UTC),
		VillageID: 4,
TeamID:    sql.NullInt64{Int64: 1, Valid: true},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		ID:        5,
		Name:      "Killer Bee",
		Nickname:  sql.NullString{String: "Eight-Tails Jinchuriki", Valid: true},
		Clan:      sql.NullString{Valid: false},
		Age:       sql.NullInt64{Int64: 30, Valid: true},
		Rank:      "Jonin",
		Birthdate: time.Date(1971, 5, 2, 0, 0, 0, 0, time.UTC),
		VillageID: 5,
	TeamID:    sql.NullInt64{Int64: 1, Valid: true},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}