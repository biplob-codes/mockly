package model

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Sku         string  `json:"sku"`
	Rating      float64 `json:"rating"`
	InStock     bool    `json:"inStock"`
	CreatedAt   string  `json:"createdAt"`
}

type NewProduct struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Sku         string  `json:"sku"`
	Rating      float64 `json:"rating"`
	InStock     bool    `json:"inStock"`
		CreatedAt   string  `json:"createdAt"`

}