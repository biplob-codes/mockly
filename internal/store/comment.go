package store

import "github.com/biplob-codes/mockly/internal/model"

var Comments = []model.Comment{
	{Id: 1, PostId: 1, Name: "Naruto Uzumaki", Email: "naruto@uzumaki.konoha", Body: "You're not alone anymore, believe it!", CreatedAt: "2023-03-01T10:00:00Z"},
	{Id: 2, PostId: 1, Name: "Sakura Haruno", Email: "sakura@haruno.konoha", Body: "We're bringing you home, no matter what.", CreatedAt: "2023-03-01T11:00:00Z"},
	{Id: 3, PostId: 2, Name: "Rock Lee", Email: "rocklee@konoha.konoha", Body: "The power of youth burns bright in your eyes!", CreatedAt: "2023-03-02T12:00:00Z"},
	{Id: 4, PostId: 3, Name: "Konohamaru Sarutobi", Email: "konohamaru@sarutobi.konoha", Body: "Someday I'll surpass every sacrifice made for this village.", CreatedAt: "2023-03-03T13:00:00Z"},
	{Id: 5, PostId: 4, Name: "Sai", Email: "sai@root.konoha", Body: "Forbidden techniques carry a cost worth studying.", CreatedAt: "2023-03-04T14:00:00Z"},
	{Id: 6, PostId: 5, Name: "Kushina Uzumaki", Email: "kushina@uzumaki.konoha", Body: "That technique runs in the family now.", CreatedAt: "2023-03-05T15:00:00Z"},
	{Id: 7, PostId: 6, Name: "Ino Yamanaka", Email: "ino@yamanaka.konoha", Body: "Sneaking off to research again, Sensei?", CreatedAt: "2023-03-06T16:00:00Z"},
	{Id: 8, PostId: 7, Name: "Shizune", Email: "shizune@konoha.konoha", Body: "Please don't gamble away the medical budget again.", CreatedAt: "2023-03-07T17:00:00Z"},
	{Id: 9, PostId: 8, Name: "Choji Akimichi", Email: "choji@akimichi.konoha", Body: "Being lazy is a strategy too, I respect it.", CreatedAt: "2023-03-08T18:00:00Z"},
	{Id: 10, PostId: 9, Name: "Temari", Email: "temari@subaku.suna", Body: "Proud of how far you've come, little brother.", CreatedAt: "2023-03-09T19:00:00Z"},
	{Id: 11, PostId: 10, Name: "Izuna Uchiha", Email: "izuna@uchiha.konoha", Body: "The clan's legacy is heavier than any sword.", CreatedAt: "2023-03-10T20:00:00Z"},
	{Id: 12, PostId: 11, Name: "Kiba Inuzuka", Email: "kiba@inuzuka.konoha", Body: "Never doubted the Byakugan for a second.", CreatedAt: "2023-03-11T21:00:00Z"},
	{Id: 13, PostId: 12, Name: "Tenten", Email: "tenten@konoha.konoha", Body: "Breaking free of fate looks good on you.", CreatedAt: "2023-03-12T22:00:00Z"},
	{Id: 14, PostId: 13, Name: "Tobirama Senju", Email: "tobirama@senju.konoha", Body: "Brother, your vision built more than you know.", CreatedAt: "2023-03-13T23:00:00Z"},
	{Id: 15, PostId: 14, Name: "Kakashi Hatake", Email: "kakashi@hatake.konoha", Body: "Some masks protect more than they hide.", CreatedAt: "2023-03-14T23:30:00Z"},
}