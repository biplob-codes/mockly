package store

import "github.com/biplob-codes/mockly/internal/model"

var Products = []model.Product{
	{Id: 1, Name: "Kunai Knife", Description: "Standard-issue steel kunai for close combat.", Price: 19.99, Category: "Weapons", Sku: "WPN-001", Rating: 4.3, InStock: true, CreatedAt: "2023-04-01T09:00:00Z"},
	{Id: 2, Name: "Shuriken Set", Description: "Set of four-point throwing stars.", Price: 14.99, Category: "Weapons", Sku: "WPN-002", Rating: 4.6, InStock: true, CreatedAt: "2023-04-02T09:00:00Z"},
	{Id: 3, Name: "Explosive Tag", Description: "Paper tag sealed with explosive chakra.", Price: 9.99, Category: "Weapons", Sku: "WPN-003", Rating: 4.1, InStock: false, CreatedAt: "2023-04-03T09:00:00Z"},
	{Id: 4, Name: "Ninja Scroll", Description: "Blank scroll for sealing jutsu and summons.", Price: 24.5, Category: "Scrolls", Sku: "SCR-001", Rating: 4.4, InStock: true, CreatedAt: "2023-04-04T09:00:00Z"},
	{Id: 5, Name: "Chakra Paper", Description: "Test paper that reveals elemental chakra nature.", Price: 4.99, Category: "Training", Sku: "TRN-001", Rating: 4.2, InStock: true, CreatedAt: "2023-04-05T09:00:00Z"},
	{Id: 6, Name: "Ninja Tool Pouch", Description: "Leg-strap holster for kunai and shuriken.", Price: 39.99, Category: "Gear", Sku: "GER-001", Rating: 4.5, InStock: true, CreatedAt: "2023-04-06T09:00:00Z"},
	{Id: 7, Name: "Summoning Contract", Description: "Blood-sealed scroll to summon toads, snakes, or slugs.", Price: 149.99, Category: "Scrolls", Sku: "SCR-002", Rating: 4.9, InStock: false, CreatedAt: "2023-04-07T09:00:00Z"},
	{Id: 8, Name: "Chakra Weights", Description: "Leg weights for taijutsu training, adjustable resistance.", Price: 34.5, Category: "Training", Sku: "TRN-002", Rating: 4.7, InStock: true, CreatedAt: "2023-04-08T09:00:00Z"},
	{Id: 9, Name: "Soldier Pill", Description: "Emergency chakra restoration pill.", Price: 15.0, Category: "Consumables", Sku: "CON-001", Rating: 4.6, InStock: true, CreatedAt: "2023-04-09T09:00:00Z"},
	{Id: 10, Name: "Anbu Mask", Description: "Hand-carved porcelain mask for covert operations.", Price: 79.99, Category: "Apparel", Sku: "APP-001", Rating: 4.3, InStock: false, CreatedAt: "2023-04-10T09:00:00Z"},
	{Id: 11, Name: "Akatsuki Cloak", Description: "Black cloak with red clouds, one size fits all.", Price: 199.99, Category: "Apparel", Sku: "APP-002", Rating: 4.8, InStock: true, CreatedAt: "2023-04-11T09:00:00Z"},
	{Id: 12, Name: "Leaf Village Headband", Description: "Steel plate headband engraved with the Konoha leaf symbol.", Price: 29.99, Category: "Apparel", Sku: "APP-003", Rating: 4.9, InStock: true, CreatedAt: "2023-04-12T09:00:00Z"},
	{Id: 13, Name: "Katana Sword", Description: "Balanced steel katana for swordsmanship training.", Price: 189.99, Category: "Weapons", Sku: "WPN-004", Rating: 4.5, InStock: true, CreatedAt: "2023-04-13T09:00:00Z"},
	{Id: 14, Name: "Chakra Control Scroll", Description: "Instructional scroll on tree-climbing and water-walking exercises.", Price: 18.75, Category: "Training", Sku: "TRN-003", Rating: 4.0, InStock: false, CreatedAt: "2023-04-14T09:00:00Z"},
	{Id: 15, Name: "Sage Mode Mat", Description: "Meditation mat for gathering natural energy.", Price: 42.99, Category: "Training", Sku: "TRN-004", Rating: 4.4, InStock: true, CreatedAt: "2023-04-15T09:00:00Z"},
}