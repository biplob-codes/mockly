package store

import "github.com/biplob-codes/mockly/internal/model"

var Comments = []model.Comment{
	{Id: 1, PostId: 1, Name: "Jane Doe", Email: "jane@example.com", Body: "Great post!", CreatedAt: "2023-03-01T10:00:00Z"},
	{Id: 2, PostId: 1, Name: "John Smith", Email: "john@example.com", Body: "Very helpful, thanks.", CreatedAt: "2023-03-01T11:00:00Z"},
	{Id: 3, PostId: 2, Name: "Alice", Email: "alice@example.com", Body: "I learned a lot from this.", CreatedAt: "2023-03-02T12:00:00Z"},
	{Id: 4, PostId: 3, Name: "Bob", Email: "bob@example.com", Body: "Nice explanation of REST.", CreatedAt: "2023-03-03T13:00:00Z"},
	{Id: 5, PostId: 4, Name: "Carol", Email: "carol@example.com", Body: "Mock servers saved my sprint.", CreatedAt: "2023-03-04T14:00:00Z"},
	{Id: 6, PostId: 5, Name: "Dave", Email: "dave@example.com", Body: "Packing light is underrated.", CreatedAt: "2023-03-05T15:00:00Z"},
	{Id: 7, PostId: 6, Name: "Eve", Email: "eve@example.com", Body: "Loved the behind the scenes look.", CreatedAt: "2023-03-06T16:00:00Z"},
	{Id: 8, PostId: 7, Name: "Frank", Email: "frank@example.com", Body: "Following this routine now.", CreatedAt: "2023-03-07T17:00:00Z"},
	{Id: 9, PostId: 8, Name: "Grace", Email: "grace@example.com", Body: "Thought-provoking read.", CreatedAt: "2023-03-08T18:00:00Z"},
	{Id: 10, PostId: 9, Name: "Heidi", Email: "heidi@example.com", Body: "Adding these to my list.", CreatedAt: "2023-03-09T19:00:00Z"},
	{Id: 11, PostId: 10, Name: "Ivan", Email: "ivan@example.com", Body: "Good intro for beginners.", CreatedAt: "2023-03-10T20:00:00Z"},
	{Id: 12, PostId: 11, Name: "Judy", Email: "judy@example.com", Body: "Trends look great this season.", CreatedAt: "2023-03-11T21:00:00Z"},
	{Id: 13, PostId: 12, Name: "Mallory", Email: "mallory@example.com", Body: "Made me laugh out loud.", CreatedAt: "2023-03-12T22:00:00Z"},
	{Id: 14, PostId: 13, Name: "Niaj", Email: "niaj@example.com", Body: "The choreography looks intense.", CreatedAt: "2023-03-13T23:00:00Z"},
	{Id: 15, PostId: 14, Name: "Olivia", Email: "olivia@example.com", Body: "Respect for the prep work.", CreatedAt: "2023-03-14T23:30:00Z"},
}
