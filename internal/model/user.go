package model

type User struct {
	Id    int
	Name  string
	Email string
}
type NewUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
