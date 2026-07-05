package model

type Comment struct {
	Id        int    `json:"id"`
	PostId    int    `json:"postId"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Body      string `json:"body"`
	CreatedAt string `json:"createdAt"`
}

type NewComment struct {
	PostId int    `json:"postId"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
		CreatedAt string `json:"createdAt"`

}