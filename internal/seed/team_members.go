package seed

import (
	"time"

	"github.com/biplob-codes/mockly/internal/database/sqlc"
)

var TeamMemberData = []sqlc.TeamMember{
	// Team Kakashi (1) — Team 7
	{TeamID: 1, CharacterID: 4, Role: "Sensei", JoinedAt: time.Now()},  // Kakashi Hatake
	{TeamID: 1, CharacterID: 1, Role: "Student", JoinedAt: time.Now()}, // Naruto Uzumaki
	{TeamID: 1, CharacterID: 2, Role: "Student", JoinedAt: time.Now()}, // Sasuke Uchiha
	{TeamID: 1, CharacterID: 3, Role: "Student", JoinedAt: time.Now()}, // Sakura Haruno

	// Team Minato (2)
	{TeamID: 2, CharacterID: 10, Role: "Sensei", JoinedAt: time.Now()},  // Minato Namikaze
	{TeamID: 2, CharacterID: 4, Role: "Student", JoinedAt: time.Now()},  // Kakashi Hatake
	{TeamID: 2, CharacterID: 15, Role: "Student", JoinedAt: time.Now()}, // Obito Uchiha

	// Team Jiraiya (3)
	{TeamID: 3, CharacterID: 8, Role: "Sensei", JoinedAt: time.Now()},   // Jiraiya
	{TeamID: 3, CharacterID: 10, Role: "Student", JoinedAt: time.Now()}, // Minato Namikaze
	{TeamID: 3, CharacterID: 26, Role: "Student", JoinedAt: time.Now()}, // Nagato
	{TeamID: 3, CharacterID: 27, Role: "Student", JoinedAt: time.Now()}, // Konan

	// Team Hiruzen (4) — the Legendary Sannin
	{TeamID: 4, CharacterID: 11, Role: "Sensei", JoinedAt: time.Now()},  // Hiruzen Sarutobi
	{TeamID: 4, CharacterID: 8, Role: "Student", JoinedAt: time.Now()},  // Jiraiya
	{TeamID: 4, CharacterID: 9, Role: "Student", JoinedAt: time.Now()},  // Tsunade
	{TeamID: 4, CharacterID: 14, Role: "Student", JoinedAt: time.Now()}, // Orochimaru
}
