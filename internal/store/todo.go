package store

import "github.com/biplob-codes/mockly/internal/model"

var Todos = []model.Todo{
	{Id: 1, UserId: 1, Title: "Finish project", Completed: false, Priority: "high", DueDate: "2026-07-10"},
	{Id: 2, UserId: 2, Title: "Write tests", Completed: true, Priority: "medium", DueDate: "2026-07-05"},
	{Id: 3, UserId: 1, Title: "Read documentation", Completed: false, Priority: "low", DueDate: "2026-07-15"},
	{Id: 4, UserId: 3, Title: "Deploy app", Completed: false, Priority: "high", DueDate: "2026-07-08"},
	{Id: 5, UserId: 4, Title: "Book flights", Completed: true, Priority: "medium", DueDate: "2026-07-01"},
	{Id: 6, UserId: 5, Title: "Review script", Completed: false, Priority: "high", DueDate: "2026-07-12"},
	{Id: 7, UserId: 6, Title: "Plan workout", Completed: true, Priority: "low", DueDate: "2026-07-03"},
	{Id: 8, UserId: 7, Title: "Attend premiere", Completed: false, Priority: "medium", DueDate: "2026-07-20"},
	{Id: 9, UserId: 8, Title: "Finish reading", Completed: false, Priority: "low", DueDate: "2026-07-18"},
	{Id: 10, UserId: 9, Title: "Service motorcycle", Completed: true, Priority: "medium", DueDate: "2026-07-02"},
	{Id: 11, UserId: 10, Title: "Fashion shoot prep", Completed: false, Priority: "high", DueDate: "2026-07-09"},
	{Id: 12, UserId: 11, Title: "Record podcast", Completed: false, Priority: "medium", DueDate: "2026-07-14"},
	{Id: 13, UserId: 12, Title: "Rehearse choreography", Completed: true, Priority: "high", DueDate: "2026-07-06"},
	{Id: 14, UserId: 13, Title: "Stunt training", Completed: false, Priority: "high", DueDate: "2026-07-11"},
	{Id: 15, UserId: 14, Title: "Language lesson", Completed: false, Priority: "low", DueDate: "2026-07-16"},
}