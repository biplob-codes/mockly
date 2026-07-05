package store

import "github.com/biplob-codes/mockly/internal/model"

var Posts = []model.Post{
	{Id: 1, UserId: 1, Title: "First Post", Body: "This is the body of the first post.", Tags: []string{"intro"}, Views: 120, CreatedAt: "2023-03-01T09:00:00Z"},
	{Id: 2, UserId: 2, Title: "Learning Go", Body: "Go concurrency is fun to learn.", Tags: []string{"go", "concurrency"}, Views: 340, CreatedAt: "2023-03-02T10:00:00Z"},
	{Id: 3, UserId: 1, Title: "REST APIs", Body: "Building REST APIs with net/http.", Tags: []string{"go", "api"}, Views: 210, CreatedAt: "2023-03-03T11:00:00Z"},
	{Id: 4, UserId: 3, Title: "Mock Servers", Body: "Why fake APIs help frontend devs.", Tags: []string{"mocking", "frontend"}, Views: 95, CreatedAt: "2023-03-04T12:00:00Z"},
	{Id: 5, UserId: 4, Title: "Traveling Light", Body: "Tips for minimalist travel.", Tags: []string{"travel"}, Views: 410, CreatedAt: "2023-03-05T13:00:00Z"},
	{Id: 6, UserId: 5, Title: "On Set Stories", Body: "Behind the scenes of my last shoot.", Tags: []string{"behind-the-scenes"}, Views: 560, CreatedAt: "2023-03-06T14:00:00Z"},
	{Id: 7, UserId: 6, Title: "Fitness Routine", Body: "My weekly workout split.", Tags: []string{"fitness", "health"}, Views: 275, CreatedAt: "2023-03-07T15:00:00Z"},
	{Id: 8, UserId: 7, Title: "Cinema and Culture", Body: "How films shape society.", Tags: []string{"cinema", "culture"}, Views: 180, CreatedAt: "2023-03-08T16:00:00Z"},
	{Id: 9, UserId: 8, Title: "Reading List", Body: "Books that changed my perspective.", Tags: []string{"books"}, Views: 300, CreatedAt: "2023-03-09T17:00:00Z"},
	{Id: 10, UserId: 9, Title: "Motorcycles 101", Body: "Getting started with riding.", Tags: []string{"motorcycles"}, Views: 150, CreatedAt: "2023-03-10T18:00:00Z"},
	{Id: 11, UserId: 10, Title: "Fashion Forward", Body: "My take on this season's trends.", Tags: []string{"fashion"}, Views: 430, CreatedAt: "2023-03-11T19:00:00Z"},
	{Id: 12, UserId: 11, Title: "Comedy Writing Tips", Body: "How I write punchlines.", Tags: []string{"comedy", "writing"}, Views: 220, CreatedAt: "2023-03-12T20:00:00Z"},
	{Id: 13, UserId: 12, Title: "Dance Practice", Body: "Choreography behind the scenes.", Tags: []string{"dance"}, Views: 310, CreatedAt: "2023-03-13T21:00:00Z"},
	{Id: 14, UserId: 13, Title: "Superhero Training", Body: "How I prep for action roles.", Tags: []string{"fitness", "acting"}, Views: 500, CreatedAt: "2023-03-14T22:00:00Z"},
	{Id: 15, UserId: 14, Title: "Language Learning", Body: "Learning a new language for a role.", Tags: []string{"language"}, Views: 140, CreatedAt: "2023-03-15T23:00:00Z"},
}