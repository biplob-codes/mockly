package store

import "github.com/biplob-codes/mockly/internal/model"

var Products = []model.Product{
	{Id: 1, Name: "Wireless Mouse", Description: "Ergonomic 2.4GHz wireless mouse.", Price: 19.99, Category: "Electronics", Sku: "ELC-001", Rating: 4.3, InStock: true, CreatedAt: "2023-04-01T09:00:00Z"},
	{Id: 2, Name: "Mechanical Keyboard", Description: "RGB backlit mechanical keyboard.", Price: 59.99, Category: "Electronics", Sku: "ELC-002", Rating: 4.6, InStock: true, CreatedAt: "2023-04-02T09:00:00Z"},
	{Id: 3, Name: "Coffee Mug", Description: "Ceramic mug, 350ml.", Price: 9.99, Category: "Kitchen", Sku: "KIT-001", Rating: 4.1, InStock: false, CreatedAt: "2023-04-03T09:00:00Z"},
	{Id: 4, Name: "Notebook", Description: "A5 dotted notebook, 120 pages.", Price: 4.5, Category: "Stationery", Sku: "STA-001", Rating: 4.4, InStock: true, CreatedAt: "2023-04-04T09:00:00Z"},
	{Id: 5, Name: "Desk Lamp", Description: "LED desk lamp with adjustable brightness.", Price: 24.99, Category: "Home", Sku: "HOM-001", Rating: 4.2, InStock: true, CreatedAt: "2023-04-05T09:00:00Z"},
	{Id: 6, Name: "Backpack", Description: "Water-resistant laptop backpack, 20L.", Price: 39.99, Category: "Bags", Sku: "BAG-001", Rating: 4.5, InStock: true, CreatedAt: "2023-04-06T09:00:00Z"},
	{Id: 7, Name: "Bluetooth Speaker", Description: "Portable speaker with 12-hour battery.", Price: 34.99, Category: "Electronics", Sku: "ELC-003", Rating: 4.0, InStock: false, CreatedAt: "2023-04-07T09:00:00Z"},
	{Id: 8, Name: "Yoga Mat", Description: "Non-slip yoga mat, 6mm thick.", Price: 21.5, Category: "Fitness", Sku: "FIT-001", Rating: 4.7, InStock: true, CreatedAt: "2023-04-08T09:00:00Z"},
	{Id: 9, Name: "Water Bottle", Description: "Insulated stainless steel bottle, 750ml.", Price: 15.0, Category: "Fitness", Sku: "FIT-002", Rating: 4.6, InStock: true, CreatedAt: "2023-04-09T09:00:00Z"},
	{Id: 10, Name: "Office Chair", Description: "Ergonomic mesh office chair.", Price: 129.99, Category: "Furniture", Sku: "FUR-001", Rating: 4.3, InStock: false, CreatedAt: "2023-04-10T09:00:00Z"},
	{Id: 11, Name: "Standing Desk", Description: "Electric height-adjustable desk.", Price: 249.99, Category: "Furniture", Sku: "FUR-002", Rating: 4.8, InStock: true, CreatedAt: "2023-04-11T09:00:00Z"},
	{Id: 12, Name: "Sunglasses", Description: "Polarized UV400 sunglasses.", Price: 29.99, Category: "Accessories", Sku: "ACC-001", Rating: 4.2, InStock: true, CreatedAt: "2023-04-12T09:00:00Z"},
	{Id: 13, Name: "Wristwatch", Description: "Minimalist analog wristwatch.", Price: 89.99, Category: "Accessories", Sku: "ACC-002", Rating: 4.5, InStock: true, CreatedAt: "2023-04-13T09:00:00Z"},
	{Id: 14, Name: "Wall Clock", Description: "Silent-sweep wall clock, 12 inch.", Price: 18.75, Category: "Home", Sku: "HOM-002", Rating: 4.0, InStock: false, CreatedAt: "2023-04-14T09:00:00Z"},
	{Id: 15, Name: "Phone Stand", Description: "Adjustable aluminum phone stand.", Price: 12.99, Category: "Electronics", Sku: "ELC-004", Rating: 4.4, InStock: true, CreatedAt: "2023-04-15T09:00:00Z"},
}