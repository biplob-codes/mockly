package store

import "github.com/biplob-codes/mockly/internal/model"

var Posts = []model.Post{
	{Id: 1, UserId: 1, Title: "Path of an Avenger", Body: "Reflecting on the choices that led me away from Konoha, and back.", Tags: []string{"uchiha", "redemption"}, Views: 120, CreatedAt: "2023-03-01T09:00:00Z"},
	{Id: 2, UserId: 2, Title: "A Thousand Jutsu, One Eye", Body: "Notes on copying techniques with the Sharingan without losing yourself.", Tags: []string{"sharingan", "training"}, Views: 340, CreatedAt: "2023-03-02T10:00:00Z"},
	{Id: 3, UserId: 3, Title: "Sacrifice for the Village", Body: "Some missions are never meant to be understood, only completed.", Tags: []string{"akatsuki", "sacrifice"}, Views: 210, CreatedAt: "2023-03-03T11:00:00Z"},
	{Id: 4, UserId: 4, Title: "Sealing Techniques 101", Body: "A breakdown of Edo Tensei and the risks of forbidden sealing jutsu.", Tags: []string{"sealing", "fuinjutsu"}, Views: 95, CreatedAt: "2023-03-04T12:00:00Z"},
	{Id: 5, UserId: 5, Title: "Speed Is Everything", Body: "How the Flying Thunder God Technique changed the shape of war.", Tags: []string{"hokage", "speed"}, Views: 410, CreatedAt: "2023-03-05T13:00:00Z"},
	{Id: 6, UserId: 6, Title: "Research Never Sleeps", Body: "Traveling the elemental nations for intel, inspiration, and a little mischief.", Tags: []string{"sannin", "research"}, Views: 560, CreatedAt: "2023-03-06T14:00:00Z"},
	{Id: 7, UserId: 7, Title: "Healing Hands, Iron Will", Body: "Medical ninjutsu isn't just chakra control, it's knowing when to gamble.", Tags: []string{"medical-ninjutsu", "hokage"}, Views: 275, CreatedAt: "2023-03-07T15:00:00Z"},
	{Id: 8, UserId: 8, Title: "Strategy Over Strength", Body: "Why the smartest move is often the laziest looking one.", Tags: []string{"strategy", "nara"}, Views: 180, CreatedAt: "2023-03-08T16:00:00Z"},
	{Id: 9, UserId: 9, Title: "From Monster to Kazekage", Body: "Learning that a bond can rewrite what you thought you were.", Tags: []string{"kazekage", "growth"}, Views: 300, CreatedAt: "2023-03-09T17:00:00Z"},
	{Id: 10, UserId: 10, Title: "The Weight of Legacy", Body: "On clan pride, old wounds, and rewriting the shinobi world.", Tags: []string{"uchiha", "legacy"}, Views: 150, CreatedAt: "2023-03-10T18:00:00Z"},
	{Id: 11, UserId: 11, Title: "Gentle Fist, Strong Heart", Body: "How quiet confidence became my greatest technique.", Tags: []string{"hyuga", "growth"}, Views: 430, CreatedAt: "2023-03-11T19:00:00Z"},
	{Id: 12, UserId: 12, Title: "Fate Is Not Absolute", Body: "Breaking free of the caged bird cage, one match at a time.", Tags: []string{"hyuga", "fate"}, Views: 220, CreatedAt: "2023-03-12T20:00:00Z"},
	{Id: 13, UserId: 13, Title: "Founding a Village From Nothing", Body: "What it took to turn a battlefield into Konohagakure.", Tags: []string{"senju", "hokage"}, Views: 310, CreatedAt: "2023-03-13T21:00:00Z"},
	{Id: 14, UserId: 14, Title: "Two Sides of the Same Mask", Body: "On wearing a mask so long you forget which face is real.", Tags: []string{"akatsuki", "identity"}, Views: 500, CreatedAt: "2023-03-14T22:00:00Z"},
	{Id: 15, UserId: 1, Title: "Rebuilding Team 7", Body: "It takes more than strength to repair broken bonds.", Tags: []string{"team7", "redemption"}, Views: 140, CreatedAt: "2023-03-15T23:00:00Z"},
}