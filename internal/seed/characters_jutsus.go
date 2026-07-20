package seed

import (
	"time"

	"github.com/biplob-codes/mockly/internal/database/sqlc"
)

var CharacterJutsuData = []sqlc.CharactersJutsu{
	{
		ID:          1,
		CharacterID: 1,
		JutsuID:     1,
		CreatedAt:   time.Now(),
	},
	{
		ID:          2,
		CharacterID: 2,
		JutsuID:     2,
		CreatedAt:   time.Now(),
	},
	{
		ID:          3,
		CharacterID: 3,
		JutsuID:     3,
		CreatedAt:   time.Now(),
	},
	{
		ID:          4,
		CharacterID: 4,
		JutsuID:     4,
		CreatedAt:   time.Now(),
	},
	{
		ID:          5,
		CharacterID: 5,
		JutsuID:     5,
		CreatedAt:   time.Now(),
	},
}