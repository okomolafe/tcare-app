package data

type Post struct {
	UserId int `json:"-"`
	Id int `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
}