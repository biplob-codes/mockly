package model

type Post struct {
	Id        int      `json:"id"`
	UserId    int      `json:"userId"`
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	Tags      []string `json:"tags"`
	Views     int      `json:"views"`
	CreatedAt string   `json:"createdAt"`
}

type NewPost struct {
	UserId    int      `json:"userId"`
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"createdAt"`
}
