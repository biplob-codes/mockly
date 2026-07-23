package seed

import (
	"time"

	"github.com/biplob-codes/mockly/internal/database/sqlc"
)

var CharacterJutsuData = []sqlc.CharactersJutsu{
	// Naruto Uzumaki (1)
	{ID: 1, CharacterID: 1, JutsuID: 3, CreatedAt: time.Now()},  // Substitution Jutsu
	{ID: 2, CharacterID: 1, JutsuID: 5, CreatedAt: time.Now()},  // Shadow Clone Jutsu
	{ID: 3, CharacterID: 1, JutsuID: 6, CreatedAt: time.Now()},  // Multi Shadow Clone Jutsu
	{ID: 4, CharacterID: 1, JutsuID: 7, CreatedAt: time.Now()},  // Rasengan
	{ID: 5, CharacterID: 1, JutsuID: 8, CreatedAt: time.Now()},  // Sage Art: Rasenshuriken
	{ID: 6, CharacterID: 1, JutsuID: 27, CreatedAt: time.Now()}, // Summoning Jutsu (toads)
	{ID: 7, CharacterID: 1, JutsuID: 30, CreatedAt: time.Now()}, // Body Flicker Jutsu
	{ID: 8, CharacterID: 1, JutsuID: 61, CreatedAt: time.Now()}, // Sage Mode: Toad Sage Training

	// Sasuke Uchiha (2)
	{ID: 9, CharacterID: 2, JutsuID: 9, CreatedAt: time.Now()},   // Fire Style: Fireball Jutsu
	{ID: 10, CharacterID: 2, JutsuID: 10, CreatedAt: time.Now()}, // Fire Style: Phoenix Sage Fire Technique
	{ID: 11, CharacterID: 2, JutsuID: 11, CreatedAt: time.Now()}, // Chidori
	{ID: 12, CharacterID: 2, JutsuID: 12, CreatedAt: time.Now()}, // Chidori Stream
	{ID: 13, CharacterID: 2, JutsuID: 13, CreatedAt: time.Now()}, // Kirin
	{ID: 14, CharacterID: 2, JutsuID: 54, CreatedAt: time.Now()}, // Sharingan
	{ID: 15, CharacterID: 2, JutsuID: 55, CreatedAt: time.Now()}, // Mangekyo Sharingan
	{ID: 16, CharacterID: 2, JutsuID: 58, CreatedAt: time.Now()}, // Susanoo
	{ID: 17, CharacterID: 2, JutsuID: 59, CreatedAt: time.Now()}, // Amaterasu

	// Sakura Haruno (3)
	{ID: 18, CharacterID: 3, JutsuID: 3, CreatedAt: time.Now()},  // Substitution Jutsu
	{ID: 19, CharacterID: 3, JutsuID: 28, CreatedAt: time.Now()}, // Summoning: Slug Divided Technique
	{ID: 20, CharacterID: 3, JutsuID: 40, CreatedAt: time.Now()}, // Creation Rebirth
	{ID: 21, CharacterID: 3, JutsuID: 44, CreatedAt: time.Now()}, // Strong Fist

	// Kakashi Hatake (4)
	{ID: 22, CharacterID: 4, JutsuID: 11, CreatedAt: time.Now()}, // Chidori
	{ID: 23, CharacterID: 4, JutsuID: 12, CreatedAt: time.Now()}, // Chidori Stream
	{ID: 24, CharacterID: 4, JutsuID: 14, CreatedAt: time.Now()}, // Water Style: Water Dragon Jutsu
	{ID: 25, CharacterID: 4, JutsuID: 17, CreatedAt: time.Now()}, // Water Style: Water Shark Bomb
	{ID: 26, CharacterID: 4, JutsuID: 30, CreatedAt: time.Now()}, // Body Flicker Jutsu
	{ID: 27, CharacterID: 4, JutsuID: 54, CreatedAt: time.Now()}, // Sharingan

	// Might Guy (5)
	{ID: 28, CharacterID: 5, JutsuID: 41, CreatedAt: time.Now()}, // Leaf Hurricane
	{ID: 29, CharacterID: 5, JutsuID: 42, CreatedAt: time.Now()}, // Leaf Whirlwind
	{ID: 30, CharacterID: 5, JutsuID: 43, CreatedAt: time.Now()}, // Eight Gates Release
	{ID: 31, CharacterID: 5, JutsuID: 44, CreatedAt: time.Now()}, // Strong Fist
	{ID: 32, CharacterID: 5, JutsuID: 47, CreatedAt: time.Now()}, // Front Lotus
	{ID: 33, CharacterID: 5, JutsuID: 48, CreatedAt: time.Now()}, // Afternoon Tiger

	// Shikamaru Nara (6)
	{ID: 34, CharacterID: 6, JutsuID: 20, CreatedAt: time.Now()}, // Shadow Possession Jutsu
	{ID: 35, CharacterID: 6, JutsuID: 21, CreatedAt: time.Now()}, // Shadow Neck Bind Jutsu

	// Hinata Hyuga (7)
	{ID: 36, CharacterID: 7, JutsuID: 45, CreatedAt: time.Now()}, // Gentle Fist
	{ID: 37, CharacterID: 7, JutsuID: 46, CreatedAt: time.Now()}, // Eight Trigrams Sixty-Four Palms
	{ID: 38, CharacterID: 7, JutsuID: 56, CreatedAt: time.Now()}, // Byakugan

	// Jiraiya (8)
	{ID: 39, CharacterID: 8, JutsuID: 7, CreatedAt: time.Now()},  // Rasengan
	{ID: 40, CharacterID: 8, JutsuID: 27, CreatedAt: time.Now()}, // Summoning Jutsu (toads)
	{ID: 41, CharacterID: 8, JutsuID: 61, CreatedAt: time.Now()}, // Sage Mode: Toad Sage Training
	{ID: 42, CharacterID: 8, JutsuID: 62, CreatedAt: time.Now()}, // Frog Kata

	// Tsunade (9)
	{ID: 43, CharacterID: 9, JutsuID: 27, CreatedAt: time.Now()}, // Summoning Jutsu (slugs)
	{ID: 44, CharacterID: 9, JutsuID: 28, CreatedAt: time.Now()}, // Summoning: Slug Divided Technique
	{ID: 45, CharacterID: 9, JutsuID: 40, CreatedAt: time.Now()}, // Creation Rebirth
	{ID: 46, CharacterID: 9, JutsuID: 44, CreatedAt: time.Now()}, // Strong Fist

	// Minato Namikaze (10)
	{ID: 47, CharacterID: 10, JutsuID: 7, CreatedAt: time.Now()},  // Rasengan
	{ID: 48, CharacterID: 10, JutsuID: 30, CreatedAt: time.Now()}, // Body Flicker Jutsu
	{ID: 49, CharacterID: 10, JutsuID: 31, CreatedAt: time.Now()}, // Flying Thunder God Jutsu

	// Hiruzen Sarutobi (11)
	{ID: 50, CharacterID: 11, JutsuID: 1, CreatedAt: time.Now()},  // Transformation Jutsu
	{ID: 51, CharacterID: 11, JutsuID: 3, CreatedAt: time.Now()},  // Substitution Jutsu
	{ID: 52, CharacterID: 11, JutsuID: 27, CreatedAt: time.Now()}, // Summoning Jutsu (monkeys)
	{ID: 53, CharacterID: 11, JutsuID: 30, CreatedAt: time.Now()}, // Body Flicker Jutsu

	// Itachi Uchiha (12)
	{ID: 54, CharacterID: 12, JutsuID: 9, CreatedAt: time.Now()},  // Fire Style: Fireball Jutsu
	{ID: 55, CharacterID: 12, JutsuID: 49, CreatedAt: time.Now()}, // Tsukuyomi
	{ID: 56, CharacterID: 12, JutsuID: 53, CreatedAt: time.Now()}, // Crow Genjutsu
	{ID: 57, CharacterID: 12, JutsuID: 54, CreatedAt: time.Now()}, // Sharingan
	{ID: 58, CharacterID: 12, JutsuID: 55, CreatedAt: time.Now()}, // Mangekyo Sharingan
	{ID: 59, CharacterID: 12, JutsuID: 58, CreatedAt: time.Now()}, // Susanoo
	{ID: 60, CharacterID: 12, JutsuID: 59, CreatedAt: time.Now()}, // Amaterasu

	// Yamato (13)
	{ID: 61, CharacterID: 13, JutsuID: 3, CreatedAt: time.Now()},  // Substitution Jutsu
	{ID: 62, CharacterID: 13, JutsuID: 18, CreatedAt: time.Now()}, // Wood Style: Wood Dragon Jutsu
	{ID: 63, CharacterID: 13, JutsuID: 19, CreatedAt: time.Now()}, // Wood Style: Deep Forest Emergence

	// Orochimaru (14)
	{ID: 64, CharacterID: 14, JutsuID: 3, CreatedAt: time.Now()},  // Substitution Jutsu
	{ID: 65, CharacterID: 14, JutsuID: 27, CreatedAt: time.Now()}, // Summoning Jutsu (snakes)
	{ID: 66, CharacterID: 14, JutsuID: 29, CreatedAt: time.Now()}, // Edo Tensei

	// Obito Uchiha (15)
	{ID: 67, CharacterID: 15, JutsuID: 19, CreatedAt: time.Now()}, // Wood Style: Deep Forest Emergence
	{ID: 68, CharacterID: 15, JutsuID: 51, CreatedAt: time.Now()}, // Demonic Illusion: Mirror Heaven and Earth Change
	{ID: 69, CharacterID: 15, JutsuID: 54, CreatedAt: time.Now()}, // Sharingan
	{ID: 70, CharacterID: 15, JutsuID: 55, CreatedAt: time.Now()}, // Mangekyo Sharingan

	// Madara Uchiha (16)
	{ID: 71, CharacterID: 16, JutsuID: 18, CreatedAt: time.Now()}, // Wood Style: Wood Dragon Jutsu
	{ID: 72, CharacterID: 16, JutsuID: 54, CreatedAt: time.Now()}, // Sharingan
	{ID: 73, CharacterID: 16, JutsuID: 55, CreatedAt: time.Now()}, // Mangekyo Sharingan
	{ID: 74, CharacterID: 16, JutsuID: 58, CreatedAt: time.Now()}, // Susanoo
	{ID: 75, CharacterID: 16, JutsuID: 59, CreatedAt: time.Now()}, // Amaterasu

	// Hashirama Senju (17)
	{ID: 76, CharacterID: 17, JutsuID: 18, CreatedAt: time.Now()}, // Wood Style: Wood Dragon Jutsu
	{ID: 77, CharacterID: 17, JutsuID: 19, CreatedAt: time.Now()}, // Wood Style: Deep Forest Emergence
	{ID: 78, CharacterID: 17, JutsuID: 57, CreatedAt: time.Now()}, // Wood Style Bloodline

	// Tobirama Senju (18)
	{ID: 79, CharacterID: 18, JutsuID: 14, CreatedAt: time.Now()}, // Water Style: Water Dragon Jutsu
	{ID: 80, CharacterID: 18, JutsuID: 30, CreatedAt: time.Now()}, // Body Flicker Jutsu
	{ID: 81, CharacterID: 18, JutsuID: 31, CreatedAt: time.Now()}, // Flying Thunder God Jutsu

	// Gaara (19)
	{ID: 82, CharacterID: 19, JutsuID: 22, CreatedAt: time.Now()}, // Sand Coffin
	{ID: 83, CharacterID: 19, JutsuID: 23, CreatedAt: time.Now()}, // Sand Tsunami
	{ID: 84, CharacterID: 19, JutsuID: 24, CreatedAt: time.Now()}, // Sand Shield

	// Sasori (20)
	{ID: 85, CharacterID: 20, JutsuID: 25, CreatedAt: time.Now()}, // Puppet Technique: Ten Puppets Performance
	{ID: 86, CharacterID: 20, JutsuID: 26, CreatedAt: time.Now()}, // Puppet Technique: Poison Mist

	// Zabuza Momochi (21)
	{ID: 87, CharacterID: 21, JutsuID: 3, CreatedAt: time.Now()},  // Substitution Jutsu
	{ID: 88, CharacterID: 21, JutsuID: 14, CreatedAt: time.Now()}, // Water Style: Water Dragon Jutsu
	{ID: 89, CharacterID: 21, JutsuID: 16, CreatedAt: time.Now()}, // Water Style: Water Prison Jutsu

	// Kisame Hoshigaki (22)
	{ID: 90, CharacterID: 22, JutsuID: 3, CreatedAt: time.Now()},  // Substitution Jutsu
	{ID: 91, CharacterID: 22, JutsuID: 15, CreatedAt: time.Now()}, // Water Style: Exploding Water Shockwave
	{ID: 92, CharacterID: 22, JutsuID: 17, CreatedAt: time.Now()}, // Water Style: Water Shark Bomb

	// Onoki (23)
	{ID: 93, CharacterID: 23, JutsuID: 30, CreatedAt: time.Now()}, // Body Flicker Jutsu
	{ID: 94, CharacterID: 23, JutsuID: 34, CreatedAt: time.Now()}, // Particle Style: Atomic Dismantling Jutsu

	// Killer Bee (24)
	{ID: 95, CharacterID: 24, JutsuID: 32, CreatedAt: time.Now()}, // Tailed Beast Ball

	// A / Fourth Raikage (25)
	{ID: 96, CharacterID: 25, JutsuID: 30, CreatedAt: time.Now()}, // Body Flicker Jutsu
	{ID: 97, CharacterID: 25, JutsuID: 35, CreatedAt: time.Now()}, // Lightning Style: Lariat

	// Nagato (26)
	{ID: 98, CharacterID: 26, JutsuID: 27, CreatedAt: time.Now()},  // Summoning Jutsu
	{ID: 99, CharacterID: 26, JutsuID: 36, CreatedAt: time.Now()},  // Shinra Tensei
	{ID: 100, CharacterID: 26, JutsuID: 37, CreatedAt: time.Now()}, // Chibaku Tensei
	{ID: 101, CharacterID: 26, JutsuID: 60, CreatedAt: time.Now()}, // Rinnegan

	// Konan (27)
	{ID: 102, CharacterID: 27, JutsuID: 38, CreatedAt: time.Now()}, // Paper Clone Jutsu
	{ID: 103, CharacterID: 27, JutsuID: 39, CreatedAt: time.Now()}, // Paper Shuriken Jutsu
}
