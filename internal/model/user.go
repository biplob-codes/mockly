package model

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zipCode"`
	Country string `json:"country"`
}

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Industry    string `json:"industry"`
}

type User struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	Phone     string  `json:"phone"`
	Website   string  `json:"website"`
	Avatar    string  `json:"avatar"`
	Address   Address `json:"address"`
	Company   Company `json:"company"`
	CreatedAt string  `json:"createdAt"`
}

type NewUser struct {
	Name      string  `json:"name"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	Phone     string  `json:"phone"`
	Website   string  `json:"website"`
	Avatar    string  `json:"avatar"`
	Address   Address `json:"address"`
	Company   Company `json:"company"`
	CreatedAt string  `json:"createdAt"`
}
